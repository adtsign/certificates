package authority

import (
	"crypto/rand"
	"encoding/binary"
	"net/http"
	"strings"

	"github.com/smallstep/certificates/templates"

	"github.com/pkg/errors"
	"github.com/smallstep/certificates/authority/provisioner"
	"github.com/smallstep/cli/crypto/randutil"
	"golang.org/x/crypto/ssh"
)

const (
	// SSHAddUserPrincipal is the principal that will run the add user command.
	// Defaults to "provisioner" but it can be changed in the configuration.
	SSHAddUserPrincipal = "provisioner"

	// SSHAddUserCommand is the default command to run to add a new user.
	// Defaults to "sudo useradd -m <principal>; nc -q0 localhost 22" but it can be changed in the
	// configuration. The string "<principal>" will be replace by the new
	// principal to add.
	SSHAddUserCommand = "sudo useradd -m <principal>; nc -q0 localhost 22"
)

// SSHKeys represents the SSH User and Host public keys.
type SSHKeys struct {
	UserKey ssh.PublicKey
	HostKey ssh.PublicKey
}

// GetSSHKeys returns the SSH User and Host public keys.
func (a *Authority) GetSSHKeys() (*SSHKeys, error) {
	var keys SSHKeys
	if a.sshCAUserCertSignKey != nil {
		keys.UserKey = a.sshCAUserCertSignKey.PublicKey()
	}
	if a.sshCAHostCertSignKey != nil {
		keys.HostKey = a.sshCAHostCertSignKey.PublicKey()
	}
	if keys.UserKey == nil && keys.HostKey == nil {
		return nil, &apiError{
			err:  errors.New("getSSHKeys: ssh is not configured"),
			code: http.StatusNotFound,
		}
	}
	return &keys, nil
}

// GetSSHConfig returns rendered templates for clients (user) or servers (host).
func (a *Authority) GetSSHConfig(typ string) ([]templates.Output, error) {
	if a.sshCAUserCertSignKey == nil && a.sshCAHostCertSignKey == nil {
		return nil, &apiError{
			err:  errors.New("getSSHConfig: ssh is not configured"),
			code: http.StatusNotFound,
		}
	}

	var ts []templates.Template
	switch typ {
	case provisioner.SSHUserCert:
		if a.config.Templates != nil && a.config.Templates.SSH != nil {
			ts = a.config.Templates.SSH.User
		}
	case provisioner.SSHHostCert:
		if a.config.Templates != nil && a.config.Templates.SSH != nil {
			ts = a.config.Templates.SSH.Host
		}
	default:
		return nil, &apiError{
			err:  errors.Errorf("getSSHConfig: type %s is not valid", typ),
			code: http.StatusBadRequest,
		}
	}

	// Render templates.
	output := []templates.Output{}
	for _, t := range ts {
		o, err := t.Output(a.config.Templates.Variables)
		if err != nil {
			return nil, err
		}
		output = append(output, o)
	}
	return output, nil
}

// SignSSH creates a signed SSH certificate with the given public key and options.
func (a *Authority) SignSSH(key ssh.PublicKey, opts provisioner.SSHOptions, signOpts ...provisioner.SignOption) (*ssh.Certificate, error) {
	var mods []provisioner.SSHCertificateModifier
	var validators []provisioner.SSHCertificateValidator

	for _, op := range signOpts {
		switch o := op.(type) {
		// modify the ssh.Certificate
		case provisioner.SSHCertificateModifier:
			mods = append(mods, o)
		// modify the ssh.Certificate given the SSHOptions
		case provisioner.SSHCertificateOptionModifier:
			mods = append(mods, o.Option(opts))
		// validate the ssh.Certificate
		case provisioner.SSHCertificateValidator:
			validators = append(validators, o)
		// validate the given SSHOptions
		case provisioner.SSHCertificateOptionsValidator:
			if err := o.Valid(opts); err != nil {
				return nil, &apiError{err: err, code: http.StatusForbidden}
			}
		default:
			return nil, &apiError{
				err:  errors.Errorf("signSSH: invalid extra option type %T", o),
				code: http.StatusInternalServerError,
			}
		}
	}

	nonce, err := randutil.ASCII(32)
	if err != nil {
		return nil, &apiError{err: err, code: http.StatusInternalServerError}
	}

	var serial uint64
	if err := binary.Read(rand.Reader, binary.BigEndian, &serial); err != nil {
		return nil, &apiError{
			err:  errors.Wrap(err, "signSSH: error reading random number"),
			code: http.StatusInternalServerError,
		}
	}

	// Build base certificate with the key and some random values
	cert := &ssh.Certificate{
		Nonce:  []byte(nonce),
		Key:    key,
		Serial: serial,
	}

	// Use opts to modify the certificate
	if err := opts.Modify(cert); err != nil {
		return nil, &apiError{err: err, code: http.StatusForbidden}
	}

	// Use provisioner modifiers
	for _, m := range mods {
		if err := m.Modify(cert); err != nil {
			return nil, &apiError{err: err, code: http.StatusForbidden}
		}
	}

	// Get signer from authority keys
	var signer ssh.Signer
	switch cert.CertType {
	case ssh.UserCert:
		if a.sshCAUserCertSignKey == nil {
			return nil, &apiError{
				err:  errors.New("signSSH: user certificate signing is not enabled"),
				code: http.StatusNotImplemented,
			}
		}
		signer = a.sshCAUserCertSignKey
	case ssh.HostCert:
		if a.sshCAHostCertSignKey == nil {
			return nil, &apiError{
				err:  errors.New("signSSH: host certificate signing is not enabled"),
				code: http.StatusNotImplemented,
			}
		}
		signer = a.sshCAHostCertSignKey
	default:
		return nil, &apiError{
			err:  errors.Errorf("signSSH: unexpected ssh certificate type: %d", cert.CertType),
			code: http.StatusInternalServerError,
		}
	}
	cert.SignatureKey = signer.PublicKey()

	// Get bytes for signing trailing the signature length.
	data := cert.Marshal()
	data = data[:len(data)-4]

	// Sign the certificate
	sig, err := signer.Sign(rand.Reader, data)
	if err != nil {
		return nil, &apiError{
			err:  errors.Wrap(err, "signSSH: error signing certificate"),
			code: http.StatusInternalServerError,
		}
	}
	cert.Signature = sig

	// User provisioners validators
	for _, v := range validators {
		if err := v.Valid(cert); err != nil {
			return nil, &apiError{err: err, code: http.StatusForbidden}
		}
	}

	return cert, nil
}

// SignSSHAddUser signs a certificate that provisions a new user in a server.
func (a *Authority) SignSSHAddUser(key ssh.PublicKey, subject *ssh.Certificate) (*ssh.Certificate, error) {
	if a.sshCAUserCertSignKey == nil {
		return nil, &apiError{
			err:  errors.New("signSSHAddUser: user certificate signing is not enabled"),
			code: http.StatusNotImplemented,
		}
	}
	if subject.CertType != ssh.UserCert {
		return nil, &apiError{
			err:  errors.New("signSSHProxy: certificate is not a user certificate"),
			code: http.StatusForbidden,
		}
	}
	if len(subject.ValidPrincipals) != 1 {
		return nil, &apiError{
			err:  errors.New("signSSHProxy: certificate does not have only one principal"),
			code: http.StatusForbidden,
		}
	}

	nonce, err := randutil.ASCII(32)
	if err != nil {
		return nil, &apiError{err: err, code: http.StatusInternalServerError}
	}

	var serial uint64
	if err := binary.Read(rand.Reader, binary.BigEndian, &serial); err != nil {
		return nil, &apiError{
			err:  errors.Wrap(err, "signSSHProxy: error reading random number"),
			code: http.StatusInternalServerError,
		}
	}

	signer := a.sshCAUserCertSignKey
	principal := subject.ValidPrincipals[0]
	addUserPrincipal := a.getAddUserPrincipal()

	cert := &ssh.Certificate{
		Nonce:           []byte(nonce),
		Key:             key,
		Serial:          serial,
		CertType:        ssh.UserCert,
		KeyId:           principal + "-" + addUserPrincipal,
		ValidPrincipals: []string{addUserPrincipal},
		ValidAfter:      subject.ValidAfter,
		ValidBefore:     subject.ValidBefore,
		Permissions: ssh.Permissions{
			CriticalOptions: map[string]string{
				"force-command": a.getAddUserCommand(principal),
			},
		},
		SignatureKey: signer.PublicKey(),
	}

	// Get bytes for signing trailing the signature length.
	data := cert.Marshal()
	data = data[:len(data)-4]

	// Sign the certificate
	sig, err := signer.Sign(rand.Reader, data)
	if err != nil {
		return nil, err
	}
	cert.Signature = sig
	return cert, nil
}

func (a *Authority) getAddUserPrincipal() (cmd string) {
	if a.config.SSH.AddUserPrincipal == "" {
		return SSHAddUserPrincipal
	}
	return a.config.SSH.AddUserPrincipal
}

func (a *Authority) getAddUserCommand(principal string) string {
	var cmd string
	if a.config.SSH.AddUserCommand == "" {
		cmd = SSHAddUserCommand
	} else {
		cmd = a.config.SSH.AddUserCommand
	}
	return strings.Replace(cmd, "<principal>", principal, -1)
}