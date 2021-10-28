package main

import (
	"fmt"
	"github.com/Torebekov/lesson3/atoi"
	"github.com/Torebekov/lesson3/fibo"
	"github.com/Torebekov/lesson3/itoa"
	"github.com/Torebekov/lesson3/reverse"
	"github.com/Torebekov/lesson3/runeindex"
	"github.com/Torebekov/lesson3/sortimports"
	"log"
)

func main() {
	fmt.Println("itoa:", itoa.Itoa(-125))


	s := "-0125"
	x, err := atoi.Atoi(s)
	if err != nil {
		log.Println("atoi:", err)
	} else {
		fmt.Println("atoi:", x)
	}


	s = "!длров уоллеХ"
	fmt.Println("reverse:", reverse.Reverse(s))


	err = sortimports.SortImports("itoa/sort3.go")
	if err != nil {
		log.Println("sort:", err)
	} else {
		fmt.Println("sort: success")
	}

	err = sortimports.SortInDir("fibo")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("sort: success")
	}


	generator := fibo.Fibonacci()
	fmt.Print("fibonacci: ")
	for i := 0; i < 10; i++ {
		fmt.Print(generator(), " ")
	}
	fmt.Print("\n")


	s = "OneLab"
	i := 3
	var r rune
	r, err = runeindex.RuneByIndex(&s, &i)
	if err != nil {
		log.Println("runeByIndex:", err)
	} else {
		fmt.Println("runeByIndex:", string(r))
	}
}
