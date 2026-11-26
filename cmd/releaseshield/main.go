package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/altmanAI/releaseshield-by-altmanai/internal/scanner"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "scan" {
		fmt.Fprintln(os.Stderr, "usage: releaseshield scan DIRECTORY")
		os.Exit(2)
	}
	r, err := scanner.Scan(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(r); err != nil {
		fmt.Fprintln(os.Stderr, "report write failed")
		os.Exit(1)
	}
	// A successful report is not an approval. Never return a release-pass status.
	os.Exit(3)
}
