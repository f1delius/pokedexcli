package main

import (
	"errors"
	"fmt"

	"github.com/f1delius/pokedexcli/internal/api"
)

func commandMap(con *config) error {
	fmt.Println(con)
	data, err := api.GetLocationArea(con.next)
	if err != nil {
		return err
	}
	if data.Previous != nil {
		con.prev = data.Previous.(string)
	}
	con.next = data.Next
	for _, country := range data.Results {
		fmt.Println(country.Name)
	}
	return nil
}

func commandMapb(con *config) error {
	if con.prev == "" {
		return errors.New("no prev locations found")
	}
	data, err := api.GetLocationArea(con.prev)
	if err != nil {
		return err
	}
	if data.Previous != nil {
		con.prev = data.Previous.(string)
	}
	for _, country := range data.Results {
		fmt.Println(country.Name)
	}
	return nil
}
