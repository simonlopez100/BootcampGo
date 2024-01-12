package main

import (
	"fmt"
)

type Item struct {
	Name  string
	Price float64
}

type StorageItemSlice struct {
	Items  map[int]Item
	lasdId int
}

func (s *StorageItemSlice) Getitem(id int) Item {
	return s.Items[id]
}

func (s *StorageItemSlice) AddItem(item *Item) {
	// Validate item
	if item.Name == "" {
		return
	}
	// Add item to storage slice
	// s.Items = append(s.Items, *item)

	// Add item to storage map
	s.Items[s.lasdId] = *item
}

func main() {
	fmt.Println("Theory Panic")
	fmt.Println("Starting...")

	// Storage as slice of items
	// store := StorageItemSlice{
	// 	Items: []Item{
	// 		{Name: "Item 1", Price: 10.0},
	// 		{Name: "Item 2", Price: 20.0},
	// 	},
	// }

	// method to add item to storage
	// store.AddItem(&Item{Name: "Item 3", Price: 30.0})

	// // method to get item from storage
	// item3 := store.Getitem(2)
	// fmt.Println(item3)

	// Example to show ow to create a panic
	// fmt.Println("Starting...")
	// _, err := os.Open("no-file.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// create the map
	store := &StorageItemSlice{
		// Items: make(map[int]Item),
	}

	store.AddItem(&Item{Name: "Item 1", Price: 10.0})
	// fmt.Println(store.Items)
	fmt.Println("End")

}
