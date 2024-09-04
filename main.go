package main

import (
	"fmt"
	"go-hash-uniq/model"
)

func main() {
	var itemNumberOne = 10
	var itemNumberTwo = 20

	items := []model.PreOrderItems{
		{
			Material:   "7JH2K9LP",
			Quantity:   1,
			SalesUnit:  "ZPTF",
			ItemNumber: &itemNumberOne,
		},
		{
			Material:   "A3N8B1XD",
			Quantity:   1,
			SalesUnit:  "ZPTF",
			ItemNumber: &itemNumberTwo,
		},
	}

	order := model.PreOrder{
		DocType:      "JMLQ",
		ClientID:     "0006163986",
		PurchaseDate: "2024-09-03",
		Items:        items,
	}

	fmt.Println("Hash:", order.GenerateHash())
	fmt.Println("--------------------------------------------------------------")

	itemsZPRE := []model.PreOrderItems{
		{
			Material:   "A3N8B1XD",
			Quantity:   1,
			SalesUnit:  "ZPTF",
			ItemNumber: &itemNumberOne,
		},
		{
			Material:   "7JH2K9LP",
			Quantity:   1,
			SalesUnit:  "ZPTF",
			ItemNumber: &itemNumberTwo,
		},
	}

	orderZPRE := model.PreOrder{
		DocType:      "JMLQ",
		ClientID:     "0006163986",
		PurchaseDate: "2024-09-03",
		Items:        itemsZPRE,
	}

	fmt.Println("Hash:", orderZPRE.GenerateHash())
}
