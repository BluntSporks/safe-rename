// safe_rename transliterates file names to ASCII-equivalent with no spaces for easier processing.
package main

import (
	"flag"
	"fmt"
	"github.com/fiam/gounidecode/unidecode"
	"log"
	"os"
	"path"
	"regexp"
)

func main() {
	// Parse flags.
	file := flag.String("file", "", "Name of file to rename")
	commit := flag.Bool("commit", false, "Commit the changes by actually renaming the file")
	flag.Parse()

	// Check arguments.
	if *file == "" {
		log.Fatal("Missing -file argument")
	}

	// Check that source file exists.
	if _, err := os.Stat(*file); err != nil {
		log.Fatal("Source file " + *file + " does not exist")
	}

	// Do processing.
	dir := path.Dir(*file)
	base := path.Base(*file)

	// Transliterate Unicode to ASCII.
	safeBase := unidecode.Unidecode(base)

	// Replace initial hyphen with underscore for argument safety.
	if safeBase[0] == '-' {
		safeBase = "_" + safeBase[1:]
	}

	// Remove any characters that are not word, dash, or dot.
	re := regexp.MustCompile(`[^\w.-]`)
	safeBase = re.ReplaceAllLiteralString(safeBase, "_")

	// Only output information if the file is being renamed.
	if safeBase != base {
		// Create new joined path in same dir location.
		joined := path.Join(dir, safeBase)

		// Abort with error message if new joined path already exists.
		if _, err := os.Stat(joined); err == nil {
			log.Fatal("Target file " + safeBase + " already exists")
		}
		if *commit {
			err := os.Rename(*file, joined)
			if err == nil {
				fmt.Printf("Renamed %s to %s\n", base, safeBase)
			} else {
				log.Fatal("Unable to rename file")
			}
		} else {
			fmt.Printf("Renaming %s to %s\n", base, safeBase)
		}
	}
}
