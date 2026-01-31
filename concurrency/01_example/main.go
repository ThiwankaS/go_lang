package main

import (
	"fmt"
)

/* #### ruotines ####
go someFunc("100")
go someFunc("101")
go someFunc("102")

// wait 02 seconds in main process
time.Sleep(time.Second * 2)
*/

/* #### channels ####
myChannel := make(chan string)

go func() {
	myChannel <- "data"
}()

msg := <-myChannel
*/

func someFunc(num string) {
	fmt.Println("this is not a number", num)
}

func main() {

	// creating a channel with buffer
	charChannel := make(chan string, 3)
	// creating a character array
	chars := []string{"a", "b", "c"}

	// iterating trough the char array and write character by charater to the channel
	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)

	// iterating trough the channel to read the data
	for result := range charChannel {
		fmt.Println(result)
	}
}
