package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/matheuspolitano/thetvdb/client"
	"github.com/matheuspolitano/thetvdb/utils"
)

func main() {

	conf, err := utils.NewConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	a, b := json.Marshal(nil)
	fmt.Println(a, b)

	apiClient, err := client.NewClient(conf.BaseURL, conf.APIKey, conf.DurationToken)
	if err != nil {
		log.Fatal(err)
	}

	companies, err := apiClient.ListCompanies()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(companies)
}
