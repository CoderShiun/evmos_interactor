package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ercBasic() returns the function name that the user wants to use.
func ercBasic() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("function list: ")
	fmt.Println("total_supply, decimals, get_balance, mint, burn, transfer")
	fmt.Println("which function are you going to use?")
	fmt.Print("-> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	text = strings.Join(strings.Fields(text), "")
	return strings.ToLower(text)
}

func ercGetBalance() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter the address")
	fmt.Print("-> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.Join(strings.Fields(input), "")
}

func ercMint() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter the amount")
	fmt.Print("-> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.Join(strings.Fields(input), "")
}

func ercBurn() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter the amount")
	fmt.Print("-> ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.Join(strings.Fields(input), "")
}

func ercTransfer() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("transfer to")
	fmt.Print("address -> ")
	addr, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	addr = strings.Join(strings.Fields(addr), "")

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("amount")
	fmt.Print("-> ")
	value, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	value = strings.Join(strings.Fields(value), "")

	return addr, value
}

func sampBasic() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("function list:")
	fmt.Println("version, set_item, get_item")
	fmt.Println("which function are you going to use?")
	fmt.Print("-> ")

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	text = strings.Join(strings.Fields(text), "")
	return strings.ToLower(text)
}

func sampSetItem() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please give key")
	fmt.Print("-> ")
	key, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	key = strings.Join(strings.Fields(key), "")

	reader = bufio.NewReader(os.Stdin)
	fmt.Println("please give value")
	fmt.Print("-> ")
	value, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	value = strings.Join(strings.Fields(value), "")

	return key, value
}

func sampGetItem() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("which item are you searching for?")
	fmt.Print("-> ")
	key, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Join(strings.Fields(key), "")
}
