package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	welcomeMessage := `
--Welcome to the number guessing game--
Enter 4 digit number, and find the secret number!
- 4 digit can not start with 0
- Can not have repetative numbers
- You can use hint keyword to get an hint
- You can surrender by typing ff
- + means the digit is in correct position
- - means the digit is in wrong position
- 0 means nothing is matched
`
	fmt.Println(welcomeMessage)
	secretNumber := createSecretN()
	//fmt.Println("secret ---------> ", secretNumber)
	compareInput(secretNumber)
}

func getNumber(secretNumber string) string {
	// get input from user
	fmt.Print("Give a 4 digit number: ")
	var number string
	fmt.Scanln(&number)
	if number == "ff" {
		return number
	} else if number == "hint" {
		randomHint := rand.Intn(3)
		fmt.Printf("The %vth digit is: %c\n", randomHint+1, secretNumber[randomHint])
		return getNumber(secretNumber)
	} else if len(number) != 4 {
		fmt.Println("Please give a 4 digit number.")
		return getNumber(secretNumber)
	}
	return number
}

func convertToString(secretNumber []int8) string {
	valuesString := []string{}
	for i := range secretNumber {
		number := secretNumber[i]
		text := strconv.Itoa(int(number))
		valuesString = append(valuesString, text)
	}
	convertToString := strings.Join(valuesString, "")
	return convertToString
}

func createSecretN() string {
	number := []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Seed(time.Now().Unix())
	var randomV int8
	var flag bool
	var secretNumber []int8

	for {
		if len(secretNumber) == 4 {
			break
		}
		if !flag {
			for {
				randomV = number[rand.Intn(len(number))]
				if randomV != 0 {
					flag = true
					secretNumber = append(secretNumber, randomV)
					break
				} else {
					randomV = number[rand.Intn(len(number))]
				}
			}
		}
		randomV2 := number[rand.Intn(len(number))]
		for len(secretNumber) < 4 {
			if randomV != randomV2 {
				secretNumber = append(secretNumber, randomV2)
				randomV = randomV2
				randomV2 = number[rand.Intn(len(number))]
			} else {
				break
			}
		}
	}
	// fmt.Println(secretNumber)
	return convertToString(secretNumber)
}

func compareInput(secretNumber string) {
	var positionCounter int
	var valueCounter int
	var guessedNumbers []string

	for {
		positionCounter = 0
		valueCounter = 0
		input := getNumber(secretNumber)
		if input == "ff" {
			fmt.Printf("The secret number was: %v", secretNumber)
			os.Exit(0)
		} else if secretNumber == input {
			fmt.Printf("+4-0\nCongratulations! %v was the secret number.", secretNumber)
		}
		for i := 0; i < 4; i++ {
			if secretNumber[i] == input[i] {
				positionCounter++
			} else if strings.Contains(secretNumber, string(input[i])) {
				valueCounter++
			}
		}
		if positionCounter == 4 {
			break
		}
		guessedNumbers = append(guessedNumbers, input)
		fmt.Print("\033[H\033[2J")
		fmt.Println("Your guesses so far: ", guessedNumbers)
		fmt.Printf("+%v-%v\n", positionCounter, valueCounter)
	}
}
