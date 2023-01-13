package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	arg := os.Args[1:]

	for _, arg := range arg {

		fmt.Printf("\nProg Arg: %s", arg)
	}
	// reader := bufio.NewReader(os.Stdin)
	scann := bufio.NewScanner(os.Stdin)

	fmt.Printf("\nyour name ?")

	// name, _ := reader.ReadString('\n')
	// fmt.Printf("\nHello %s", strings.TrimSpace(name))
	if scann.Scan() {
		name := scann.Text()
		fmt.Printf("\nHello %s", name)
	}
	// var n string
	// fmt.Print("enter name")
	// fmt.Scanf("%s", &n)
	// fmt.Println("my name", n)

	const milkm = 200
	var miles float64

	fmt.Print("\nEnter miles:")
	fmt.Scanf("%f", &miles)
	kilometers := miles * milkm
	fmt.Println(kilometers)
}
