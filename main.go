package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Read sentences from the file
	file, err := os.Open("sentences.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sentences []string
	for scanner.Scan() {
		sentences = append(sentences, scanner.Text())
	}

	// Start the touch typing practice
	fmt.Println("Welcome to Touch Typing Practice!")
		
	fmt.Println("Type the sentences below as accurately and quickly as you can.\n")

	// Give the user a moment to prepare
	time.Sleep(2 * time.Second)

	// Loop through sentences
	for _, sentence := range sentences {
		fmt.Println("Sentence to type:")
		fmt.Println(sentence)

		// Start the timer
		startTime := time.Now()

		// Get user's input
		fmt.Print("Your input: ")
		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		// Stop the timer
		endTime := time.Now()

		// Calculate the time taken to type the sentence
		duration := endTime.Sub(startTime)

		// Calculate typing speed in words per minute (WPM)
		words := strings.Fields(sentence)
		wordCount := len(words)
		seconds := duration.Seconds()
		typingSpeed := float64(wordCount) / (seconds / 60.0)

		// Compare user's input with the correct sentence
		if userInput == sentence {
		
			fmt.Printf("Correct! Well done! Typing Speed: %.2f WPM\n", typingSpeed)
			time.Sleep(2 * time.Second)
			fmt.Println("\nCongratulations! You have completed the touch typing practice.")
						
		} else {
			fmt.Printf("Incorrect. Try again. Typing Speed: %.2f WPM\n", typingSpeed)
			
		}

		// Give the user a moment before showing the next sentence
		time.Sleep(2 * time.Second)
	}
	



}

