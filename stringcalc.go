package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	exp, _ := reader.ReadString('\n')
	exp = strings.TrimSpace(exp)
	var action rune
	var data []string

	if strings.Contains(exp, " + ") {
		data = strings.Split(exp, " + ")
		action = '+'
	} else if strings.Contains(exp, " - ") {
		data = strings.Split(exp, " - ")
		action = '-'
	} else if strings.Contains(exp, " * ") {
		data = strings.Split(exp, " * ")
		action = '*'
	} else if strings.Contains(exp, " / ") {
		data = strings.Split(exp, " / ")
		action = '/'
	} else {
		panic("Некорректный знак действия")
	}

	if action == '*' || action == '/' {
		if strings.Contains(data[1], "\"") {
			panic("Строчку можно делить или умножать только на число")
		}
	}

	for i := range data {
		data[i] = strings.ReplaceAll(data[i], "\"", "")
	}

	switch action {
	case '+':
		printInQuotes(data[0] + data[1])
	case '*':
		multiplier, _ := strconv.Atoi(data[1])
		result := ""
		for i := 0; i < multiplier; i++ {
			result += data[0]
		}
		printInQuotes(result)
	case '-':
		index := strings.Index(data[0], data[1])
		if index == -1 {
			printInQuotes(data[0])
		} else {
			result := data[0][:index] + data[0][index+len(data[1]):]
			printInQuotes(result)
		}
	case '/':
		newLen := len(data[0]) / atoi(data[1])
		result := data[0][:newLen]
		printInQuotes(result)
	}
}

func printInQuotes(text string) {
	fmt.Printf("\"%s\"\n", text)
}

func atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}
