package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Setup the configuration from env file
func LoadConfigs() {

	wd, err := os.Getwd()

	if err != nil {
		log.Fatal("Failed to find current working directory")
	}

	projectRootPath := findRootDir(wd)

	if projectRootPath == "" {
		log.Fatalf("Failed to find the root directory")
	}

	err = godotenv.Load(projectRootPath + "/.env")
	if err != nil {
		log.Fatalf("Couldn't fetch the env %s", err.Error())
	}

	loadAppConfig()
	loadAuthConfig()
	loadDBConfig()
	loadSearchEngineConfig()
	loadSMTPConfig()
	loadWhatsappConfig()
}

// Get the path to root directory by looking for go.mod file
func findRootDir(dirPath string) string {

	//Check if the file has .root file
	if _, err := os.Stat(filepath.Join(dirPath, ".root")); err != nil {
		return dirPath
	}

	//Move up one level and search
	parentDir := filepath.Dir(dirPath)
	if parentDir == dirPath {
		//Reached file system root without finding go.mod
		return ""
	}

	return findRootDir(parentDir)
}
