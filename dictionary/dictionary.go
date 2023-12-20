package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Dictionary struct {
	filePath string
}

func NewDictionary(filePath string) *Dictionary {
	return &Dictionary{filePath: filePath}
}

func (d *Dictionary) Add(word, definition string) error {
	existingDef, err := d.Get(word)
	if err == nil {
		return fmt.Errorf("Le mot '%s' existe déjà avec la définition : %s", word, existingDef)
	}

	entry := fmt.Sprintf("%s:%s\n", word, definition)
	file, err := os.OpenFile(d.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(entry)
	if err != nil {
		return err
	}

	return nil
}



func (d *Dictionary) Get(word string) (string, error) {
	file, err := os.Open(d.filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 && parts[0] == word {
			return parts[1], nil
		}
	}

	return "", fmt.Errorf("Mot non trouvé : %s", word)
}


func (d *Dictionary) Remove(word string) error {
	lines, err := readLines(d.filePath)
	if err != nil {
		return err
	}

	var newLines []string
	found := false
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 && parts[0] == word {
			found = true
			continue
		}
		newLines = append(newLines, line)
	}

	if !found {
		return fmt.Errorf("Mot non trouvé : %s", word)
	}

	return writeLines(d.filePath, newLines)
}


func (d *Dictionary) List() ([]string, error) {
	lines, err := readLines(d.filePath)
	if err != nil {
		return nil, err
	}

	sort.Strings(lines)
	return lines, nil
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func writeLines(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
