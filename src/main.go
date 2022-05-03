package main

import (
	"fmt"
	proto2 "github.com/golang/protobuf/proto"
	"os"
	"protobuff/proto"
	"time"
)

func main() {

	message := proto.Game{
		Team: &proto.Team{
			Home: "Barcelona",
			Away: "Madrid",
		},
		Venue: "Barcelona Stadium",
		Date:  time.Now().String(),
	}

	marshal, err := proto2.Marshal(&message)

	if err != nil {
		return
	}

	file, _ := os.OpenFile("game", os.O_WRONLY|os.O_CREATE, 0644)

	file.Write(marshal)

	file.Close()

	bytes, _ := os.ReadFile("game")

	data := proto.Game{}

	err = proto2.Unmarshal(bytes, &data)

	if err != nil {
		return
	}

	fmt.Println(data.Date)

	fmt.Println(data.Team.Away)

	fmt.Println(data.Team.Home)

	fmt.Println(data.Venue)

}
