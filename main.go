package main

import (
	"github.com/kqnd/dsa/dsa"
)


func main() {
	queue := dsa.CreateQueue(1)
	queue.Add(2)
	queue.Add(3)
	queue.Remove()
	queue.Show()
}
