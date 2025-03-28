package utils

import (
	"os"
	"path/filepath"
)

// GetPDFFiles recursively walks a directory and collects all PDF file paths
func GetPDFFiles(root string) ([]string, error) {
	var pdfFiles []string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err // Return error if any issue occurs while walking
		}
		if !d.IsDir() && filepath.Ext(path) == ".pdf" {
			pdfFiles = append(pdfFiles, path)
		}
		return nil
	})

	return pdfFiles, err
}
