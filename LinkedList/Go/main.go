package main

import (
	"fmt"
	"strconv"
)

type BackPack struct {
	player int32
	head   *InventorySlot
}

type InventorySlot struct {
	itemName string
	next     *InventorySlot
}

func createBackpack(name string) *BackPack {
	return &BackPack{
		player: 1,
	}
}

func (backpack *BackPack) pickupItem(itemName string) error {
	newItem := &InventorySlot{itemName: itemName}

	if backpack.head == nil {
		backpack.head = newItem
	} else {
		currentNode := backpack.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newItem
	}
	return nil
}

func (backpack *BackPack) showAllItems() error {
	currentNode := backpack.head
	fmt.Println(currentNode.itemName)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Println(currentNode.itemName)
	}

	return nil
}

func main() {
	backPack := createBackpack("Player")
	fmt.Printf("Player %s \n", strconv.FormatInt(int64(backPack.player), 10))

	backPack.pickupItem("Item 1")
	backPack.pickupItem("Item 2")
	backPack.pickupItem("Item 3")
	backPack.pickupItem("Item 4")
	backPack.pickupItem("Item 5")

	backPack.showAllItems()

}
