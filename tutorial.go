/**
* This is a quiz game made in Go.
* Author: Abhishek Biswas Deep
**/

package main

import "fmt"

//Welcoming the user and asking for name and age.
func main() {
	fmt.Println("Welcome to the computer quiz game!")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	//After verification, the user is allowed to play the game.
	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You cannot play!")
		return
	}

	score := 0
	num_questions := 5

	//Different types of questions are asked during the game.
	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer + " " + answer2 == "RTX 3090" || answer + " " + answer2 == "rtx 3090" || answer + " " + answer2 == "rtX 3090" {
		fmt.Println("Correct!")
		score ++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("Who is the father of Computer science? ")
	var fatherName string
	var fatherName2 string
	fmt.Scan(&fatherName, &fatherName2)

	if fatherName + " " + fatherName2 == "Charles Babbage" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("In a computer, most processing takes place in______? ")
	var part string
	fmt.Scan(&part)

	if part == "CPU" || part == "cpu"{
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("Scientific Name of Computer?  ")
	var person string
	var person2 string
	fmt.Scan(&person, &person2)

	if person + " " + person2 == "Sillico sapiens" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	//After the game has ended, the user is able to see the score.
	fmt.Printf("You scored %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%%.", percent)

}