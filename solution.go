package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	hwMessage := emoji.Sprint("Hello :world_map:!")
	return hwMessage
}

func main() {

	fmt.Println(GetMessage())
}