package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Dictionary struct {
	filePath string
	addCh    chan entryOperation
	removeCh chan string
	mu       sync.Mutex
}

type entryOperation struct {
	word       string
	definition string
}

func NewDictionary(filePath string) *Dictionary {
	dict := &Dictionary{
		filePath: filePath,
		addCh:    make(chan entryOperation),
		removeCh: make(chan string),
	}

	go dict.handleOperations()

	return dict
}

func (d *Dictionary) handleOperations() {
	for {
		select {
		case operation := <-d.addCh:
			d.handleAdd(operation.word, operation.definition)
		case word := <-d.removeCh:
			d.handleRemove(word)
		}
	}
}

func (d *Dictionary) handleAdd(word, definition string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	existingDef, err := d.Get(word)
	if err == nil {
		fmt.Printf("Le mot '%s' existe déjà avec la définition : %s\n", word, existingDef)
		return
	}

	entry := fmt.Sprintf("%s:%s\n", word, definition)
	file, err := os.OpenFile(d.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Erreur lors de l'ouverture du fichier : %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(entry)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture dans le fichier : %v\n", err)
	}
}

func (d *Dictionary) handleRemove(word string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	lines, err := readLines(d.filePath)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		return
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
		fmt.Printf("Mot non trouvé : %s\n", word)
		return
	}

	err = writeLines(d.filePath, newLines)
	if err != nil {
		fmt.Printf("Erreur lors de l'écriture dans le fichier : %v\n", err)
	}
}

func (d *Dictionary) Get(word string) (string, error) {
	lines, err := readLines(d.filePath)
	if err != nil {
		return "", err
	}

	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 && parts[0] == word {
			return parts[1], nil
		}
	}

	return "", fmt.Errorf("Mot non trouvé : %s", word)
}

// readLines et writeLines sont des fonctions auxiliaires pour la lecture et l'écriture de lignes dans le fichier.

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

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
func (d *Dictionary) Add(word, definition string) {
	d.addCh <- entryOperation{word, definition}
}

// Remove supprime un mot du dictionnaire.
func (d *Dictionary) Remove(word string) {
	d.removeCh <- word
}

// List retourne une liste triée des mots et de leurs définitions.
func (d *Dictionary) List() ([]string, error) {
	return readLines(d.filePath)
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
