package file

import (
	"bufio"
	"errors"
	"os"
)

// Create a file with the given path.
func CreateFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

// Write to a file, if it does not exist create it.
func WriteFile(w string, path string) error {
	e, err := CheckFile(path)
	if err != nil {
		return err
	}
	if !e {
		err = CreateFile(path)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	b := []byte(w + "\n")
	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}

// Delete a file with the given path.
func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

// Read the contents of a given file into a []string.
func ReadFile(path string) ([]string, error) {
	e, err := CheckFile(path)
	if err != nil {
		return nil, err
	}
	if !e {
		return nil, os.ErrNotExist
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var l []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}
	return l, scanner.Err()
}

// Check if a file exists at the given path.
func CheckFile(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// Check if the given file contains the given string.
func FileContains(c string, path string) (bool, error) {
	lines, err := ReadFile(path)
	if err != nil {
		return false, err
	}

	for _, l := range lines {
		if l == c {
			return true, nil
		}
	}

	return false, nil
}
