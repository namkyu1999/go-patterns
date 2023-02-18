package pipeline

import (
	"fmt"
	"time"
)

func DoPipeline() {
	startTime := time.Now()

	initChannel := make(chan string)
	process1Channel := make(chan string)
	process2Channel := make(chan string)
	process3Channel := make(chan string)
	resultChannel := make(chan string)

	go doProcessOne(initChannel, process1Channel)
	go doProcessTwo(process1Channel, process2Channel)
	go doProcessThree(process2Channel, process3Channel)
	go doProcessFour(process3Channel, resultChannel)

	requestValue(initChannel)
	close(initChannel)

	result := <-resultChannel
	fmt.Println("the result is: ", result)
	fmt.Println("total time: ", time.Since(startTime))
}

func doProcessOne(initChannel chan string, process1Channel chan string) {
	time.Sleep(150 * time.Millisecond)
	process1Channel <- concatStage(<-initChannel, "first step")
	close(process1Channel)
}

func doProcessTwo(process1Channel chan string, process2Channel chan string) {
	time.Sleep(200 * time.Millisecond)
	process2Channel <- concatStage(<-process1Channel, "second step")
	close(process2Channel)
}

func doProcessThree(process2Channel chan string, process3Channel chan string) {
	time.Sleep(20 * time.Millisecond)
	process3Channel <- concatStage(<-process2Channel, "third step")
	close(process3Channel)
}

func doProcessFour(process3Channel chan string, resultChannel chan string) {
	time.Sleep(90 * time.Millisecond)
	resultChannel <- concatStage(<-process3Channel, "result step")
	close(resultChannel)
}

func concatStage(current, next string) string {
	return current + " -> " + next
}

func requestValue(initChannel chan string) {
	time.Sleep(300 * time.Millisecond)
	initChannel <- "start point"
}
