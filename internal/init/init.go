package init

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
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
		filepath.Join("tasks", "next"),
		filepath.Join("tasks", "defer"),
		filepath.Join("tasks", "waiting_for"),
		"tickler",
	}

	for _, dir := range dirs {
		err := os.Mkdir(filepath.Join(name, dir), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func Init(directory string) error {
	isEmpty, err := IsDirEmpty(directory)
	if err != nil {
		return err
	}

	if !isEmpty {
		return fmt.Errorf("init %s: directory is not empty", directory)
	}

	fmt.Println("Creating structure...")
	return CreateDirs(directory)
}
