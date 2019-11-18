package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"

	"booking-service/commons"
)

type Client struct { // Our example struct, you can use "-" to ignore a field
	Id      string `csv:"client_id"`
	Name    string `csv:"client_name"`
	Age     string `csv:"client_age"`
	NotUsed string `csv:"-"`
}

func main() {
	clientsFile, err := os.OpenFile("../../../data/room_detail.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	// csvContent, err := gocsv.CSVToMap
	// if err != nil {
	// panic(err)
	// }
	clients := []*commons.Room{}

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}
	for _, client := range clients {
		fmt.Println("RoomID = ", client.RoomID, " ,RoomType", client.RoomType, " ,RoomName", client.RoomName)
	}
}
