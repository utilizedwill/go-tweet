package utils

import (
	"encoding/json"
	"go-tweet/constants"
	"io/ioutil"
	"log"
)

func AuthDeveloper() *constants.AuthCreds {

	var authStruct constants.AuthCreds

	authJson, err := ioutil.ReadFile("static/auth.json")
	if err != nil {
		log.Fatal("Error happened while reading auth.json file")
	}

	err = json.Unmarshal(authJson, &authStruct)
	if err != nil {
		log.Fatal(err)
	}

	return &authStruct
}
