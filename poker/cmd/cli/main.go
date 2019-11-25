package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game.PlayPoker()
	fileSystemPlayerStoreClose()
}
