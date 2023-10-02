package main

import (
	"crypto/sha256"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"
)

// Fonction pour générer un mot de passe à partir de chiffres ASCII et d'une suite de Cauchy
func generatePassword() string {
	password := ""
	cauchy := 0.0
	for i := 0; i < 32; i++ {
		// Utilisez la suite de Cauchy pour générer un nombre entre 32 et 126 (ASCII)
		cauchy = math.Mod(cauchy+1.0, 95.0) + 32.0
		password += string(int(cauchy))
	}
	fmt.Println(password)
	return password
}

// Fonction pour vérifier le mot de passe
func checkPassword(userInput string) bool {
	generatedPassword := generatePassword()
	hashedGeneratedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(generatedPassword)))
	hashedUserInput := fmt.Sprintf("%x", sha256.Sum256([]byte(userInput)))

	return hashedUserInput == hashedGeneratedPassword
}

// Fonction pour détecter ltrace
func detectLtrace() bool {
	cmd := exec.Command("ltrace", "ls")
	output, _ := cmd.CombinedOutput()
	return strings.Contains(string(output), "ltrace")
}

func main() {
	if detectLtrace() {
		fmt.Println("Détection de ltrace : Accès refusé.")
		os.Exit(1)
	}

	var userInput string
	fmt.Print("Entrez le mot de passe : ")
	fmt.Scanln(&userInput)

	if len(userInput) != 32 {
		fmt.Println("Mot de passe incorrect. Accès refusé.")
		os.Exit(1)
	}

	if checkPassword(userInput) {
		fmt.Println("Mot de passe correct. Accès autorisé.")
	} else {
		fmt.Println("Mot de passe incorrect. Accès refusé.")
	}
}
