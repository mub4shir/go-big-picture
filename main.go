package main

import "fmt"

func main() {

	i := 1
	i++ // increment and decrements are statements in Go
	fmt.Println(i)
	i++
	fmt.Println(i)

	for i := 0; i < 5; i++ { // loop with increment
		fmt.Print(i, " ")
	}
	for i < 10 { // loop till condition
		fmt.Print(i, " ")
		i++
	}
	for { // infinite loop
		fmt.Print("hi")
		break
	}
	users := [2]string{"hi", "hey"}
	fmt.Println(users)
	for user := range users {
		fmt.Println(user, " ")
	}
}
