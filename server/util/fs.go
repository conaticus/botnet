package util

import (
	"io"
	"log"
	"os"
)

func CreateIfNotExists(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		CreateFile(path, true)
	}
}

// Returns true if successful.
func CreateFile(path string, fatal bool) bool {
	file, err := os.Create(path)
	if err != nil {
		if fatal {
			log.Fatalf("Failed to create file: %s", err.Error())
		}

		Error("Failed to create file, %s", err.Error())	
		return false
	}

	file.Close()
	return true
}

func ReadFile(path string, fatal bool) *string {
	CreateIfNotExists(path)
	
	file, err := os.Open(path)
	if err != nil {
		if fatal {
			log.Fatalf("Failed to open file, %s", err.Error())
		}

		Error("Failed to open file, %s", err.Error())	
		return nil
	}

	buffer := make([]byte, 0)
	tempBuffer := make([]byte, 1024)

	for {
		n, err := file.Read(tempBuffer)
		buffer = append(buffer, tempBuffer[:n]...)
		if n == 0 || err == io.EOF {
			break
		}
		
		if err != nil {
			if fatal {
				log.Fatalf("Failed to read file, %s", err.Error())
			}

			Error("Failed to read file, %s", err.Error())	
			return nil
		}
	}

	contents := string(buffer)
	return &contents
}

func AppendFile(path string, data string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		Error("Failed to open file, %s", err.Error())
		return
	}

	defer file.Close()

	_, err = file.Write([]byte(data))
	if err != nil {
		Error("Failed to append file, %s", err.Error())
		return
	}
}