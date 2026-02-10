package media

import (
	"math/rand"
	"os"
)


func RandomImage(dir string)(string, error){
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	var files []string
	for _, entry := range entries{
		files = append(files, entry.Name())
	}

	randomFile := files[rand.Intn(len(files))]

	return dir + "/" + randomFile, nil
}