package secretservice

import (
	"log"
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

func GetSecrets(secretType string) (SecretList, error) {
	session, collections, err := GetAllCollections()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := session.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	search := map[string]string{
		"type": secretType,
	}
	result, err := collections[0].SearchItems(search)
	if err != nil {
		return nil, err
	}
	var secrets SecretList

	for _, i := range result {
		attributes, _ := i.GetAttributes()
		sec, err := i.GetSecret(session.Path())
		if err != nil {
			return nil, err
		}

		secrets = append(secrets, RenderSecret(attributes, string(sec.Value)))
	}
	return secrets, nil
}

func RenderSecret(attributes map[string]string, secret string) Secret {
	return Secret{
		UUID:     attributes["Uuid"],
		Notes:    attributes["Notes"],
		Path:     attributes["Path"],
		Title:    attributes["Title"],
		URL:      attributes["URL"],
		UserName: attributes["UserName"],
		Type:     attributes["Type"],
		Secret:   secret,
	}
}
