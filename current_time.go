package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		// Récupérer l'heure actuelle
		heureActuelle := time.Now()

		// Afficher l'heure au format HH:MM:SS
		fmt.Printf("\rHeure actuelle : %s", heureActuelle.Format("15:04:05"))

		// Attendre une seconde avant de mettre à jour
		time.Sleep(1 * time.Second)
	}
}
