package main

import "fmt"

func main() {
	sendsSoFar := 340
	const sendsToAdd = 25

	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("You've sent", sendsSoFar, "messages so far!")

}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar += sendsToAdd
	return sendsSoFar
}
