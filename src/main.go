package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/brunobmello25/http-client/src/components"
	"github.com/brunobmello25/http-client/src/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	collectionPath := flag.String("collection", "collection.json", "Path to the collection file")
	flag.Parse()

	collection, err := models.LoadCollection(*collectionPath)
	if err != nil {
		fmt.Printf("Error loading collection: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(components.NewUI(collection))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
