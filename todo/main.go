package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func SaveList(filename string, list []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(list)
}

func LoadList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return []string{}, nil
	} else if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var list []string
	err = decoder.Decode(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func main() {
	// Define flags
	addPtr := flag.String("add", "", "adds a task to the list")
	removePtr := flag.Int("remove", -1, "removes a task")
	helpPtr := flag.Bool("help", false, "prints help message")

	// Parse
	flag.Parse()

	// Get values
	list, err := LoadList("list.json")
	if err != nil {
		fmt.Println("Error loading list:", err)
		return
	}

	// Handle flags
	if *addPtr != "" {
		list = append(list, *addPtr)
	} else if *removePtr != -1 && *removePtr < len(list) && *removePtr >= 0 {
		// Remove task
		copy(list[*removePtr:], list[*removePtr+1:])
		list = list[:len(list)-1]
	} else if *helpPtr {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		return
	} else {
		fmt.Println("Invalid flag use -help for help")
	}
	
	for i, task := range list {
		fmt.Println(i, task)
	}

	err = SaveList("list.json", list)
	if err != nil {
		fmt.Println("Error saving list:", err)
		return
	}
	
}