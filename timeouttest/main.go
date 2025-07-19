package main

import (
	"fmt"
	"time"
)

// timeout testing

func main() {

	timout:= time.NewTimer(time.Second*3)
	fmt.Println("empezamos timer")


	channelcito:= make(chan string)

	select{
	case <- channelcito:
		fmt.Println("llegmaos al channel")
	case <- timout.C: 
	fmt.Println("llegamos al teimouttt!!!")	
	timout.Stop()
	}
}



func coso(ch chan string){
	fmt.Println("vamos a esperar 10 chan")
	time.Sleep(time.Second*10)


}