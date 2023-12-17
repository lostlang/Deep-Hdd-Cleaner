package common

import (
	"fmt"
	"os"
)

func CleanHDD(size int, data *[]byte) error {
	iter := size / len(*data)
	fmt.Printf("Cleaning HDD need %d iterations\n", iter)
	fmt.Print("Cleaning... Iteration: 0")
	for i := 0; i < iter; i++ {
		writeFile(fmt.Sprintf("%d", i), data)
		fmt.Printf("\rCleaning... Iteration: %d", i+1)
	}
	return nil
}

func writeFile(name string, data *[]byte) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(*data); err != nil {
		return err
	}

	return nil
}
