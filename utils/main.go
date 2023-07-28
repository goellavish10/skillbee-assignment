package utils

import (
	"fmt"
	"os"
)

func CreateDir(dirname string) {
	directoryName := dirname

	// Check if the directory already exists
	_, err := os.Stat(directoryName)
	if os.IsNotExist(err) {
		// Directory doesn't exist, create it
		err = os.Mkdir(directoryName, 0755) // 0755 sets the directory permissions
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		fmt.Println("Directory created:", directoryName)
	} else if err == nil {
		// Directory already exists
		fmt.Println("Directory already exists:", directoryName)
	} else {
		// Some other error occurred
		fmt.Println("Error checking directory:", err)
	}
}
