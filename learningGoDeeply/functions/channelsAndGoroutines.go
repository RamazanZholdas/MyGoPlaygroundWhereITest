package functions

import "fmt"

func GoRoutinesAndChannels1() {
	ch := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("sending", i, "to the channel")
			ch <- i
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Receiving", <-ch)
	}
}

func GoRoutinesAndChannels2() {
	chCom := make(chan int, 2)
	exit := make(chan struct{})

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Sending", i)
			chCom <- i
		}

		close(chCom)
	}()

	go func() {
		for v := range chCom {
			fmt.Println("Receiving", v)
		}

		close(exit)
	}()

	<-exit
}

func FanInPattern() {

}

