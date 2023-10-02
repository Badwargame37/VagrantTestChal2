package main

import (
	"crypto/sha256"
	"fmt"
)

// Fonction pour générer un mot de passe avec Fibonacci
func generatePassword() string {
	password := ""
	a, b := 0, 1
	for i := 0; i < 16; i++ {
		a, b = b, a+b                        // Suite de Fibonacci
		password += string('A' + byte(a%26)) // Convertir en lettre majuscule
	}
	// Inverser la chaîne de caractères
	reversedPassword := []rune(password)
	for i, j := 0, len(reversedPassword)-1; i < j; i, j = i+1, j-1 {
		reversedPassword[i], reversedPassword[j] = reversedPassword[j], reversedPassword[i]
	}
	return string(reversedPassword)
}

// Fonction pour vérifier le mot de passe
func checkPassword(userInput string) (bool, string, string) {
	generatedPassword := generatePassword()

	// Hasher le mot de passe généré (pour des raisons de sécurité)
	hashedGeneratedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(generatedPassword)))

	// Comparer le mot de passe de l'utilisateur avec le mot de passe généré (hashé)
	return userInput == hashedGeneratedPassword, generatedPassword, userInput
}

func main() {
	var userInput string
	fmt.Print("Entrez le mot de passe : ")
	fmt.Scanln(&userInput)

	match, generatedPassword, enteredPassword := checkPassword(userInput)

	if match {
		fmt.Println("Mot de passe correct. Accès autorisé.")
	} else {
		fmt.Println("Mot de passe incorrect. Accès refusé.")
		fmt.Printf("Mot de passe généré : %s\n", generatedPassword)
		fmt.Printf("Mot de passe entré par l'utilisateur : %s\n", enteredPassword)
	}
}
