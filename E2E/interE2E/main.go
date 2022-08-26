package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("1. account balance test:")
	accountBalance()
	fmt.Println("Ok")

	fmt.Println("2. contract list test:")
	contractList()
	fmt.Println("Ok")

	fmt.Println("3. deploy sample contract test:")
	deploySample()
	fmt.Println("Ok")

	fmt.Println("4. get sample contract version test:")
	getSampleVersion()
	fmt.Println("Ok")

	fmt.Println("5. deploy erc20 contract test:")
	deployErc20()
	fmt.Println("Ok")

	fmt.Println("6. erc20 mint function test:")
	erc20Mint()
	fmt.Println("Ok")

	fmt.Println("7. erc20 burn function test:")
	erc20Burn()
	fmt.Println("Ok")

	fmt.Println("8. erc20 transfer function test:")
	erc20Tx()
	fmt.Println("Ok")

	fmt.Println("clean data...")
	err := os.Remove("./sample.address")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove("./erc20.address")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done.")
}
