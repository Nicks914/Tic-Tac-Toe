package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read integer safely
func ReadInt(prompt string) int {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print(prompt)

		input, err := reader.ReadString('\n')

		if err != nil {

			fmt.Println("Error reading input")
			continue
		}

		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)

		if err != nil {

			fmt.Println("Please enter a valid number")
			continue
		}

		return num
	}
}
