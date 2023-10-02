package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"
)

// Function to generate a password from ASCII digits and a Cauchy sequence
func generatePassword() string {
	password := ""
	cauchy := 0.0
	q := 0
	for i := 0; i < 32; i++ {
		// Use the Cauchy sequence to generate a number between 41 ('A') and 5A ('Z') in ASCII
		cauchy = (math.Mod(cauchy+.2, 96.0) + 12.0)
		q = ((int(cauchy)) % 24) + 65
		fmt.Println(q)
		password += string(q)
	}
	return password
}

// Function to check the password
func checkPassword(userInput string) bool {
	generatedPassword := generatePassword()
	//hashedGeneratedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(generatedPassword)))
	//hashedUserInput := fmt.Sprintf("%x", sha256.Sum256([]byte(userInput)))

	//return hashedUserInput == hashedGeneratedPassword
	return generatedPassword == userInput
}

// Function to detect ltrace
func detectLtrace() bool {
	cmd := exec.Command("ltrace", "ls")
	output, _ := cmd.CombinedOutput()
	return strings.Contains(string(output), "ltrace")
}

func main() {
	generatedPassword := generatePassword() // Generate the password at the beginning
	fmt.Println("Mot de passe généré au début : ", generatedPassword)

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
