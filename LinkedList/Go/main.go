package main

import (
	"fmt"
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

func (b *BackPack) pickupItem(itemName string) error {
	newItem := &InventorySlot{
		itemName: itemName,
	}

	if b.head == nil {
		b.head = newItem
	} else {
		currentNode := b.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newItem
	}
	return nil
}

func (b *BackPack) showAllItems() error {
	currentNode := b.head
	if currentNode == nil {
		fmt.Println("Playlist is empty.")
		return nil
	}

	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}

	return nil
}

func main() {
	fmt.Println("hello world")
}
