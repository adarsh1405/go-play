package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUserData() string {
	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// fmt.Printf("Response Type: %T\n", content). //content is of type []uint8 -
	// fmt.Println(content)

	return string(content)
}

func Decodedata(data string) {

	valid := json.Valid([]byte(data))

	if valid {

		err := json.Unmarshal([]byte(data), &accounts)
		if err != nil {
			panic(err)
		}

		// fmt.Printf("Decoded Data: %+v\n", accounts)
		// for _, account := range accounts {
		// 	fmt.Printf("ID: %d\nName: %s\nUsername: %s\nEmail: %s\nCompany Name: %s\nCatch Phrase: %s\n\n",
		// 		account.ID, account.Name, account.Username, account.Email, account.Company.Name, account.Company.CatchPhrase)
		// }
	} else {
		fmt.Println("The data is NOT valid JSON")
	}

}