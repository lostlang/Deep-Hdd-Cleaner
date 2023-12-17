package main

import (
	"fmt"

	"github.com/lostlang/deep-hdd-cleaner/common"
)

func main() {
	fmt.Print("How many RAM do you want to allocate? RAM: ")
	ram_allocate, err := common.ReadValue()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("How many HDD do you want to clean? HDD: ")
	hdd_size, err := common.ReadValue()
	if err != nil {
		fmt.Println(err)
		return
	}

	if ram_allocate > hdd_size {
		ram_allocate = hdd_size
	}

	fmt.Print("Allocating RAM...")
	freezeMem := common.LazyGenerateBytes(ram_allocate)
	fmt.Println("\rAllocating RAM... Done")

	common.CleanHDD(hdd_size, freezeMem)
	fmt.Println("\nCleaning... Done")
}
