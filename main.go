package main

import (
	"fmt"

	"github.com/kqnd/dsa/dsa"
)

func main() {
	
	book1 := dsa.CreateBook("The Great Gatsby")
	book2 := dsa.CreateBook("1984")

	lt := dsa.CreateLeafTree()
	lt.Add(nil)
	lt.Add(nil)
	lt.Add(nil)
	lt.Add(nil)
	lt.Add(nil)
	bookOneDB := lt.Add(book1)
	bookTwoDB := lt.Add(book2)

	fmt.Println("========= LEAF TREE =========")
	lt.Show()

	fmt.Println("========= END LEAF TREE =========")

	bookByKey := lt.Find(bookOneDB.Key)

	bookByID := lt.FindByBookID(bookTwoDB.Leaf.ID)

	if bookByKey != nil {
		fmt.Printf("FIND_KEY: [%d] Book: %s\n", bookOneDB.Key, bookByKey.Leaf.Title)
	}

	if bookByID != nil {
		fmt.Printf("FIND_BY_ID: [%d] Book: %s\n", bookTwoDB.Key, bookByID.Leaf.Title)
	}

}