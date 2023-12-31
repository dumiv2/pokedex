package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"

	"io/ioutil"
	"net/http"
	"os"
)

func main(){
	makeMap()
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex >")
		reader.Scan()
		executeCommand(reader.Text())
	}
}

type comman struct {
	name string
	description string
	callback func() error
}

func executeCommand(input string) error {
	err := errors.New("Command not found")
	count := 0 
	switch {
	case input == "exit" :
		return commandExit()
	case input == "help" :
		return commandHelp()
	case input == "map" :
		count ++
		return Map()
	case input == "mapb" :
		count --
		return Mapb()
	default :
		fmt.Println(err.Error())
		return nil
	}
}

func makeMap(){
	command := map[string]comman{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp ,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map" : {
			name : "map",
			description: "print 20 locations",
			callback: Map,
		},
		"mapb" : {
			name : "map",
			description: "print 20 locations",
			callback: Mapb,
		},
	}
	fmt.Sprintf("%v",command)
}

func commandHelp() error {
	err := errors.New("helppppp")
	fmt.Printf("%v","This is our available command : \n")
	fmt.Println(".help : show available command")
	fmt.Println(".exit : close your connection")
	return err
}

func commandExit() error {
	err := errors.New("1")
	os.Exit(1)
	return err
}

func Map() error {
	maplocation := AutoGenerated{}
	res , _ := http.Get("https://pokeapi.co/api/v2/location-area/")
	body, _ := ioutil.ReadAll(res.Body)
	finalresult := string(body)
	err := json.Unmarshal([]byte(finalresult),&maplocation)
	for i := 0 ; i < 20 ; i ++ {
		fmt.Println(maplocation.Results[i].Name)
	}
	fmt.Println(maplocation.Next)
	return err
}

func Mapb() error {
	maplocation := AutoGenerated{}
	res , _ := http.Get("https://pokeapi.co/api/v2/location-area//")
	body, _ := ioutil.ReadAll(res.Body)
	finalresult := string(body)
	err := json.Unmarshal([]byte(finalresult),&maplocation)
	for i := 0 ; i < 20 ; i ++ {
		fmt.Println(maplocation.Results[i].Name)
	}
	fmt.Println(maplocation.Next)
	return err
}


type AutoGenerated struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}