package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f,err := os.Create("proverb.txt")
	if err !=nil{
		log.Fatal(err.Error())
	}
	n,err :=f.Write([]byte("don't communicate by sharing memory share memory by communicating"))
	if err !=nil{
		log.Fatal(err.Error())
	}


	fmt.Println(n)
}
