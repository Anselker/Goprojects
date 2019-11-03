package main

import (
	"fmt"
	"os"
)

func main() {
	//f,_ :=os.Create("proverb.txt")
	//_,_ =f.Write([]byte("don't communicate by sharing memory share memory by communicating"))
	//b :=make([]byte,1024)

	f, _ := os.OpenFile("proverb.txt", os.O_RDONLY, 0)
	b:=make([]byte,1024)
	n,err :=f.Read(b)
	fmt.Println(n,err)
	fmt.Println(string (b))
}
