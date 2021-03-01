package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tkircsi/proto/person"
	"google.golang.org/protobuf/proto"
)

func main() {
	fname := "../boldi_v3.bin"
	fmt.Println("Server is running...")
	p := person.Person{
		Age:  11,
		Name: "Boldi",
		Phones: []string{
			"555-444",
			"222-3333",
		},
		Email:        "boldi@mail.hu",
		FavoriteGame: person.FavoriteGames_TWO_POINT_HOSPITAL,
		Subject: map[string]int32{
			"Matematika": 4,
			"Angol":      5,
			"Magyar":     5,
		},
	}

	err := writeToFile(fname, &p)
	if err != nil {
		log.Fatal("error save to file: ", err)
	}
}

func writeToFile(fn string, pb proto.Message) error {
	bs, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	err = os.WriteFile(fn, bs, 0644)
	if err != nil {
		return err
	}
	return nil
}
