package demo

import "fmt"

func SayHello() {
	fmt.Println("----> SayHello: ");
    sayHello();
}

func sayHello() {
	fmt.Println("----> sayHello: ");
    fmt.Println("Hello, This Is Demo!");
}