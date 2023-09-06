package init

import (
	"fmt"
	"io"
	"log"
	"os"
)

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // read 1 child of the directory
	if err == io.EOF {         // if there are no children, directory is empty
		return true, nil
	}
	return false, err // either not empty or error, suits both cases
}

func CreateDirs(name string) error {
	dirs := []string{
		"inbox",
		"projects",
		"tasks",
		"tasks/next",
		"tasks/defer",
		"tasks/waiting_for",
		"tickler",
	}

	for _, dir := range dirs {
		err := os.Mkdir(name+"/"+dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func Init(directory string) error {
	if directory == "" {
		directory = "."
	}

	isEmpty, err := IsDirEmpty(directory)
	if err != nil {
		return err
	}

	if !isEmpty {
		log.Printf("[ERROR] Directory %s is not empty!", directory)
		return nil
	}

	fmt.Println("Creating structure...")
	return CreateDirs(directory)
}
