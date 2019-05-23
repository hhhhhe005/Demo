package Demo

import (
	"fmt"
	"Demo/methods"
)

func Hello(name string){
	n := methods.GetHello(name)
	fmt.Printf("hello %s !\n",n)
}
