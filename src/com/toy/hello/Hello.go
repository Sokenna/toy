package main

import (
	"fmt"
	"github.com/Sokenna/toy/src/com/toy/greeting"
	"log"
	"math/rand"
)

func main() {
	log.SetPrefix("greetins: ")
	log.SetFlags(2)
	hello, err := greeting.Hello("Jimi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hello)
	fmt.Println(rand.Int())
	names := []string{"snopy", "jack", "tom"}
	hellos, err := greeting.Hellos(names)
	if err != nil {
		return
	}
	fmt.Println(hellos)
	for s, s2 := range hellos {
		fmt.Println(s, s2)
	}

}
