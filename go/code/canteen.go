// +build OMIT

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	t "time"
)

type nom struct {
	Day  string `json:"day"`
	Meal []struct {
		Info []struct {
			Menu string `json:"menu"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"info"`
		Type string `json:"type"`
	} `json:"meal"`
}

// START1 OMIT
// getJSON handles JSON retrieval and decoding to struct
func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// STOP1 OMIT

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printMenu(social nom) {
	for _, meal := range social.Meal {
		fmt.Println(meal.Type)
		for _, info := range meal.Info {
			fmt.Printf("%s - %s\n", info.Menu, info.Name)
		}
	}
}

func main() {
	social := make(chan []nom, 1)
	today := t.Now()

	go func() {
		var socialget []nom
		err := getJSON("https://fenix.tecnico.ulisboa.pt/api/fenix/v1/canteen", &socialget)
		handleError(err)
		social <- socialget
	}()

	select {
	case socialget := <-social:
		//For each social meal
		for _, k := range socialget {
			parsedTime, err := t.Parse("2/1/2006", k.Day)
			handleError(err)
			switch {
			case parsedTime.Day() == today.Day():
				fmt.Println("Hoje,", k.Day, "foi:")
				printMenu(k)
				os.Exit(0)
			}
		}
	case <-t.After(t.Second * 3):
		fmt.Println("Connection timed out.")
	}
}
