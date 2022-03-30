package main

import (
	"fmt"
	"log"

	"github.com/godbus/dbus/v5"
	"github.com/ktr0731/go-fuzzyfinder"
	keyring "github.com/ppacher/go-dbus-keyring"
)

type Secret struct {
	UUID     string
	Notes    string
	Path     string
	Title    string
	URL      string
	UserName string
	Type     string
	Secret   string
}

type SecretList []Secret

func main() {
	conn, err := dbus.SessionBus()
	checkErr(err)

	svc, err := keyring.GetSecretService(conn)
	checkErr(err)

	// session is required to create/retrieve secrets
	session, err := svc.OpenSession()
	checkErr(err)

	defer func() { checkErr(session.Close()) }()
	// Get all collections available
	collection, err := svc.GetAllCollections()
	checkErr(err)

	search := map[string]string{
		"type": "kubeconfig",
	}
	result, err := collection[0].SearchItems(search)
	checkErr(err)

	var secrets SecretList

	for _, i := range result {
		lala, _ := i.GetAttributes()
		sec, err := i.GetSecret(session.Path())
		if err != nil {
			fmt.Println(err)
		}

		secrets = append(secrets, Secret{
			UUID:     lala["Uuid"],
			Notes:    lala["Notes"],
			Path:     lala["Path"],
			Title:    lala["Title"],
			URL:      lala["URL"],
			UserName: lala["UserName"],
			Type:     lala["Type"],
			Secret:   string(sec.Value),
		})
	}

	idx, err := fuzzyfinder.Find(
		secrets,
		func(i int) string {
			return secrets[i].Title
		})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("selected: %v\n", secrets[idx].Secret)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
