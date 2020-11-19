package files

import (
	"log"
	"os"
)

//ToCreate project files
func ToCreate(infra bool) {

	var files = [4]string{"main", "variables", "outputs", "datasources"}

	if infra {
		_, err := os.OpenFile("providers.tf", os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, file := range files {
		_, err := os.OpenFile(file+".tf", os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

}
