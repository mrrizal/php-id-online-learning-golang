package main

import (
	"fmt"
	"time"
)

type User struct {
	ID         int
	Username   string
	ProfilePic string
}

type Item struct {
	ID    int
	Nama  string
	Harga float64
}

type Order struct {
	ID   int
	User User
	Item Item
}

func getUserData(order *Order) {
	// hanya contoh, pada kenyataan nya
	// function ini akan nge call service lain
	// baik menggunakan rest atau grpc
	time.Sleep(500 * time.Millisecond)
	userData := User{
		ID:         1,
		Username:   "Rizal",
		ProfilePic: "https://google.com/image.jpg",
	}
	order.User = userData
}

func getItemData(order *Order) {
	time.Sleep(500 * time.Millisecond)
	itemData := Item{
		ID:    1,
		Nama:  "alat penyedot ingus",
		Harga: 100.000,
	}
	order.Item = itemData
}

func getOrderData() Order {
	var orderData Order
	orderData.ID = 1
	// todo: use wait group and call as goroutine
	getUserData(&orderData)
	getItemData(&orderData)
	return orderData
}

func main() {
	orderData := getOrderData()
	fmt.Printf("%+v\n", orderData)
}
