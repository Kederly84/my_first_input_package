package input_package

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func StringReader() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	return input
}

func IntReader() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	res, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	return int(res)
}

func FloatReader() float64 {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	res, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func BoolReader() bool {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	res, err := strconv.ParseBool(input)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
