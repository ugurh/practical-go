package main

import (
	"fmt"
	"math/rand"
	"time"
)

const version string = "1.3.2"

func main() {

	now := time.Now()
	fmt.Println("Now :", now)
	n := 2548
	fmt.Printf("%x\n", n)
	n2 := 0x9F4
	fmt.Printf("Decimal : %d\n", n2)

	var roomNumber, floorNumber = 4, 3

	fmt.Println("Room:", roomNumber)
	fmt.Println("Floor", floorNumber)

	language := "Go"
	fmt.Println("Language :", language)

	fmt.Println("Version:", version)

	const defaultType = 12
	var type1 uint8
	var type2 int64
	var type3 float32

	type1 = defaultType
	type2 = defaultType
	type3 = defaultType

	fmt.Println(type1, type2, type3)

	// default type is bool
	const isOpen = true
	// default type is rune (alias for int32)
	const MyRune = 'r'
	// default type is int
	const occupancyLimit = 12
	// default type is float64
	const vatRate = 29.87
	// default type is complex128
	const complexNumber = 1 + 2i
	// default type is string
	const hotelName0 = "Gopher Hotel"

	fmt.Println("Is Open:", isOpen)
	fmt.Println("MyRune:", MyRune)
	fmt.Println("Occupancy Limit:", occupancyLimit)
	fmt.Println("Vat Rate:", vatRate)
	fmt.Println("Complex Number:", complexNumber)

	const profit = 9223372036854775807
	var profit2 int64 = profit
	fmt.Println(profit2)

	UKTimezoneName := "Europe/London"

	result, err := time.LoadLocation(UKTimezoneName)
	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Println(result)

	// Random values seeded
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Hotel " + hotelName0)
	var roomsAvailable int
	var rooms, roomsOccupied = 100, rand.Intn(100)
	roomsAvailable = rooms - roomsOccupied

	if rooms != roomsOccupied {
		fmt.Println(roomsAvailable, "rooms available")
	} else {
		fmt.Println("No rooms available")
	}

	var agePaul, ageJohn = rand.Intn(110), rand.Intn(110)
	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else if agePaul == ageJohn {
		fmt.Println("Paul and John have the same age")
	} else {
		fmt.Println("Paul is younger than John")
	}

	// switch expression
	switch ageJohn {
	case 10:
		fmt.Println("John is 10 years old")
	case 20:
		fmt.Println("John is 20 years old")
	case 100:
		fmt.Println("John is 100 years old")
	default:
		fmt.Println("John has neither 10,20 nor 100 years old")
	}

	// switch statement; expression
	switch ageSum := ageJohn + agePaul; ageSum {
	case 10:
		fmt.Println("ageJohn + agePaul = 10")
	case 20, 30, 40:
		fmt.Println("ageJohn + agePaul = 20 or 30 or 40")
	case 2 * agePaul:
		fmt.Println("ageJohn + agePaul = 2 times agePaul")
	}

	// switch (without statement, without expression)
	switch {
	case agePaul > ageJohn:
		fmt.Println("agePaul > ageJohn")
	case agePaul == ageJohn:
		fmt.Println("agePaul == ageJohn")
	case agePaul < ageJohn:
		fmt.Println("agePaul < ageJohn")
	}

	const emailToSend = 3
	emailSent := 0

	for emailSent < emailToSend {
		fmt.Println("sending email", emailSent+1, "times")
		emailSent++
	}
	fmt.Println("end of program")

	const hotelName = "Gopher Paris Inn"
	const totalRooms = 134
	const firstRoomNumber = 110

	rand.Seed(time.Now().UTC().UnixNano())
	occupied := rand.Intn(totalRooms)
	available := totalRooms - roomsOccupied

	occupancyRate := calOccupancyRate(occupied, totalRooms)
	occupancyLevel := getOccupancyLevel(occupancyRate)

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms", totalRooms)
	fmt.Println("Rooms available", available)
	fmt.Println("Occupancy Level:", occupancyLevel)
	fmt.Printf("Occupancy Rate: %0.2f %%\n", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; i < roomsAvailable; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			printRoomDetails(roomNumber, size, nights)
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}
}

func getOccupancyLevel(occupancyRate float64) string {
	var occupancyLevel string
	if occupancyRate > 70 {
		occupancyLevel = "High"
	} else if occupancyRate > 20 {
		occupancyLevel = "Medium"
	} else {
		occupancyLevel = "Low"
	}
	return occupancyLevel
}

func printRoomDetails(roomNumber int, size int, nights int) {
	nightText := "nights"
	if nights == 1 {
		nightText = "night"
	}
	fmt.Println(roomNumber, ":", size, "people /", nights, nightText)
}

func calOccupancyRate(roomsOccupied int, totalRooms int) float64 {
	return (float64(roomsOccupied) / float64(totalRooms)) * 100
}

/**
1- Run
	go run main.go
2- Compilation
	go build main.go
3- Launch Application
	./main

Untyped constants default types are :
	- bool (for any boolean value)
	- rune (for any rune value)
	- int (for any integer value)
	- float64 (for any floating-point value)
	- complex128 (for any complex value)
	- string (for any string value)
*/
