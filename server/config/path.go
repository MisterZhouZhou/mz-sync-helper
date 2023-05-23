package config

import (
	"log"
	"os"
	"path/filepath"
)

func DefaultPath() (uploads string) {
	ep, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	// ext/../Resources
	dir := filepath.Join(filepath.Dir(ep), "..", "Resources")
	return dir
}

// ext/../Resources/uploads
var UploadsDir = filepath.Join(DefaultPath(), "uploads")

// ext/../Resources/fronted/dist
var FrontedDistDir = filepath.Join(DefaultPath(), "fronted/dist")
