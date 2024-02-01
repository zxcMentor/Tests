package main

import (
	"context"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/model"
	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
	"log"
)

type GG struct {
	Api *suggest.Api
}

func main() {

	clC := client.Credentials{
		ApiKeyValue:    "5086f9aa3d01c20cab4d1477df59cb0f1ab75497",
		SecretKeyValue: "01c3fde0996a6e08e1d51d5203c57cdb211739b2",
	}

	api := &GG{Api: dadata.NewSuggestApi(client.WithCredentialProvider(&clC))}

	params := suggest.RequestParams{Query: "ул Свободы"}
	suggestions, err := api.Api.Address(context.Background(), &params)
	if err != nil {
		log.Fatal("err:", err)
	}

}

type AddressSuggestion struct {
	Value             string         `json:"value"`
	UnrestrictedValue string         `json:"unrestricted_value"`
	Data              *model.Address `json:"data"`
}
