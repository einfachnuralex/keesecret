package secretservice

import (
	"github.com/godbus/dbus/v5"
	keyring "github.com/ppacher/go-dbus-keyring"
)

func GetAllCollections() (keyring.Session, []keyring.Collection, error) {
	conn, err := dbus.SessionBus()
	if err != nil {
		return nil, nil, err
	}

	svc, err := keyring.GetSecretService(conn)
	if err != nil {
		return nil, nil, err
	}

	// session is required to create/retrieve secrets
	session, err := svc.OpenSession()
	if err != nil {
		return nil, nil, err
	}

	// Get all collections available
	collection, err := svc.GetAllCollections()
	if err != nil {
		return nil, nil, err
	}

	return session, collection, nil
}
