package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/AmulyaParitosh/GoCrypt/filecrypt"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandler()
	case "decrypt":
		decryptHandler()
	default:
		fmt.Println("Run encrypt to encrypt a file, nad decrypt to decrypt a file.")
	}
}

func printHelp() {
	fmt.Println("GoCrypt")
	fmt.Println("Simple file encrypter for your day-to-day needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tCryptoGo encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file given a password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("")
}

func encryptHandler() {
	if len(os.Args) < 3 {
		println("missing the path to the file. Run help for more info.")
		os.Exit(0)
	}

	filePath := os.Args[2]
	if !validateFile(filePath) {
		panic("File not found")
	}

	password := getPassword()

	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(filePath, password)
	fmt.Println("File successfully encrypted")
}

func decryptHandler() {
	if len(os.Args) < 3 {
		println("missing the path to the file. Run help for more info.")
		os.Exit(0)
	}

	filePath := os.Args[2]
	if !validateFile(filePath) {
		panic("File not found")
	}

	fmt.Print("Enter password: ")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDecrypting...")

	filecrypt.Decrypt(filePath, password)
	fmt.Println("File successfully decrypted.")

}

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm Password: ")
	password2, _ := term.ReadPassword(0)

	if !validatePassword(password, password2) {
		fmt.Println("\nPasswords do not match. Please try again")
		return getPassword()
	}

	return password
}

func validatePassword(password1 []byte, password2 []byte) bool {
	return bytes.Equal(password1, password2)
}

func validateFile(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
