package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	length := 200
	var e error

	if len(os.Args) == 2 {
		length, e = strconv.Atoi(os.Args[1])
		if e != nil {
			fmt.Println("pass only integer for speed in milliseconds")
			return
		}
		fmt.Println(length)
	}

	var currentTime = time.Now()
	var hour, minute, second = currentTime.Hour(), currentTime.Minute(), currentTime.Second()
	var timeArray [8]placeholder
	var displayArray [8]placeholder
	var counter int

	for {
		screen.Clear()

		currentTime = time.Now()
		hour = currentTime.Hour()
		minute = currentTime.Minute()
		second = currentTime.Second()

		timeArray[0] = nums[hour/10]
		timeArray[1] = nums[hour%10]
		switch second % 2 {
		case 0:
			timeArray[2] = separator
			timeArray[5] = separator
		case 1:
			timeArray[2] = blank
			timeArray[5] = blank
		}
		timeArray[3] = nums[minute/10]
		timeArray[4] = nums[minute%10]
		timeArray[6] = nums[second/10]
		timeArray[7] = nums[second%10]

		offset := counter % 16

		if offset < len(timeArray) {
			for i := range timeArray {
				diff := i + offset
				if diff < len(timeArray) {
					displayArray[i] = timeArray[diff]
				}
			}
		} else {
			for i := range timeArray {
				diff := i + offset - 16
				if diff >= 0 {
					displayArray[i] = timeArray[diff]
				}
			}
		}

		for i := range one {
			for j := range displayArray {
				fmt.Printf("%s", displayArray[j][i])
				fmt.Printf(" ")
			}
			fmt.Printf("\n")
		}

		time.Sleep(time.Millisecond * time.Duration(length))
		for i := range displayArray {
			displayArray[i] = blank
		}

		counter++

	}
}
