package main

import (
	"fmt"
	"log"
	"os/exec"
)

func accountBalance() {
	balance, err := exec.Command("evmosInteractor", "account", "balance").Output()
	if err != nil {
		log.Fatal(err)
	}

	if len(balance) == 0 {
		log.Fatal("get contract balance error")
	}

	fmt.Println(string(balance))
}
