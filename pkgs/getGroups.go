package pkgs

import (
	"encoding/json"
	"fmt"
	"os"
	"whatsappWiz/models"
)

func GetGroups() []models.Group {
	file, err := os.Open("data/groups.json")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return nil
	}
	defer file.Close()

	var groups []models.Group
	err = json.NewDecoder(file).Decode(&groups)
	if err != nil {
		fmt.Println("Error al decodificar el archivo JSON:", err)
		return nil
	}

	return groups
}
