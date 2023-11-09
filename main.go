package main

import (
	"bufio"
	cRand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"os/user"
	"strconv"
	"strings"
)

const (
	// all lowercase letters
	lower = "abcdefghijklmnopqrstuvwxyz"
	// all uppercase letters
	upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// all numbers
	numbers = "0123456789"
	// all symbols
	symbols  = "!@#$%^&*()_+{}[];':,./<>?`~"
	allChars = lower + upper + numbers + symbols
)

type PasswordInfo struct {
	InicialLength   int
	RemainingLength int
	Lower           int
	Upper           int
	Numbers         int
	Symbols         int
}

func getUserInput(prompt string) (string, error) {

	fmt.Println(prompt)

	// Scanning for user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), nil

}

func convertToInt(s string) (int, error) {
	if s == "" {
		return 0, nil
	}

	return strconv.Atoi(s)
}

func getRandomCryptoNumber(max int) (int, error) {
	// Get random number
	idx, err := cRand.Int(cRand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}

	return int(idx.Int64()), nil
}

func getRandomNumber(max int) int {
	return rand.Int() % max
}

func buildPasswordInfoRandom(passwordSize int) (PasswordInfo, error) {

	passwordInfo := PasswordInfo{
		InicialLength:   0,
		RemainingLength: 0,
		Lower:           0,
		Upper:           0,
		Numbers:         0,
		Symbols:         0,
	}

	if passwordSize <= 0 {
		return PasswordInfo{}, fmt.Errorf("password length must be greater than 0")
	}

	// Get password length
	passwordInfo.InicialLength = passwordSize
	passwordInfo.RemainingLength = passwordSize

	// Get number of lowercase letters
	passwordInfo.Lower = getRandomNumber(passwordSize)

	if passwordInfo.Lower >= passwordInfo.RemainingLength {
		return passwordInfo, nil
	} else {
		passwordInfo.RemainingLength -= passwordInfo.Lower
	}

	// Get number of uppercase letters
	passwordInfo.Upper = getRandomNumber(passwordSize)

	if passwordInfo.Upper >= passwordInfo.RemainingLength {
		return passwordInfo, nil
	} else {
		passwordInfo.RemainingLength -= passwordInfo.Upper
	}

	// Get number of numbers
	passwordInfo.Numbers = getRandomNumber(passwordSize)

	if passwordInfo.Numbers >= passwordInfo.RemainingLength {
		return passwordInfo, nil
	} else {
		passwordInfo.RemainingLength -= passwordInfo.Numbers
	}

	// Get number of symbols
	passwordInfo.Symbols = getRandomNumber(passwordSize)

	if passwordInfo.Symbols >= passwordInfo.RemainingLength {
		passwordInfo.Symbols = passwordInfo.RemainingLength
		passwordInfo.RemainingLength = 0
		return passwordInfo, nil
	} else {
		passwordInfo.RemainingLength -= passwordInfo.Symbols
	}

	return passwordInfo, nil

}

func buildPasswordInfo() (PasswordInfo, error) {

	passwordInfo := PasswordInfo{
		InicialLength:   0,
		RemainingLength: 0,
		Lower:           0,
		Upper:           0,
		Numbers:         0,
		Symbols:         0,
	}

	// Get password length
	passwordLength, err := getUserInput("Enter password length: ")
	if err != nil && passwordLength == "" {
		return passwordInfo, err
	}

	// Convert password length to int
	passwordInfo.InicialLength, err = convertToInt(passwordLength)
	if err != nil {
		return passwordInfo, err
	}

	if passwordInfo.InicialLength <= 0 || passwordInfo.InicialLength > 100 {
		fmt.Println("password length must be greater than 0 and less than 100")
	}

	passwordInfo.RemainingLength = passwordInfo.InicialLength

	// Get number of lowercase letters
	lowercase, err := getUserInput("Enter number of lowercase letters: ")
	if err != nil {
		return passwordInfo, err
	}

	// Convert lowercase to int
	passwordInfo.Lower, err = convertToInt(lowercase)
	if err != nil {
		return passwordInfo, err
	}

	if passwordInfo.Lower >= passwordInfo.RemainingLength {
		return passwordInfo, nil
	} else {
		passwordInfo.RemainingLength -= passwordInfo.Lower
	}

	// Get number of uppercase letters
	uppercase, err := getUserInput("Enter number of uppercase letters: ")
	if err != nil {
		return passwordInfo, err
	}

	// Convert uppercase to int
	passwordInfo.Upper, err = convertToInt(uppercase)
	if err != nil {
		return passwordInfo, err
	}

	if passwordInfo.Upper >= passwordInfo.RemainingLength {
		return passwordInfo, nil
	} else {
		passwordInfo.RemainingLength -= passwordInfo.Upper
	}

	// Get number of numbers
	numbers, err := getUserInput("Enter number of numbers: ")
	if err != nil {
		return passwordInfo, err
	}

	// Convert numbers to int
	passwordInfo.Numbers, err = convertToInt(numbers)
	if err != nil {
		return passwordInfo, err
	}

	if passwordInfo.Numbers >= passwordInfo.RemainingLength {
		return passwordInfo, nil
	} else {
		passwordInfo.RemainingLength -= passwordInfo.Numbers
	}

	// Get number of symbols
	symbols, err := getUserInput("Enter number of symbols: ")
	if err != nil {
		return passwordInfo, err
	}

	// Convert symbols to int
	passwordInfo.Symbols, err = convertToInt(symbols)
	if err != nil {
		return passwordInfo, err
	}

	if passwordInfo.Symbols >= passwordInfo.RemainingLength {
		passwordInfo.Symbols = passwordInfo.RemainingLength
		passwordInfo.RemainingLength = 0
		return passwordInfo, nil
	} else {
		passwordInfo.RemainingLength -= passwordInfo.Symbols
	}

	return passwordInfo, nil
}

func appendRandomChars(lenght int, s *string) error {

	for i := 0; i < lenght; i++ {
		idx, err := getRandomCryptoNumber(len(allChars))
		if err != nil {
			return err
		}

		*s += string(allChars[idx])
	}

	return nil
}

func createPassword(passInfo PasswordInfo) (string, error) {
	var password string

	// Append lowercase letters
	err := appendRandomChars(passInfo.Lower, &password)

	if err != nil {
		return "", err
	}

	// Append uppercase letters
	err = appendRandomChars(passInfo.Upper, &password)

	if err != nil {
		return "", err
	}

	// Append numbers

	err = appendRandomChars(passInfo.Numbers, &password)

	if err != nil {
		return "", err
	}

	// Append symbols

	err = appendRandomChars(passInfo.Symbols, &password)

	if err != nil {
		return "", err
	}

	// Append remaining chars
	if passInfo.RemainingLength > 0 {
		err = appendRandomChars(passInfo.RemainingLength, &password)

		if err != nil {
			return "", err
		}
	}

	return password, nil
}

func scramblePassword(password string) string {
	r := []rune(password)

	rand.Shuffle(len(r), func(i, j int) {
		r[i], r[j] = r[j], r[i]
	})

	return password
}

func main() {

	var password string
	currentUser, err := user.Current()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if currentUser.Username == "root" {
		fmt.Println("You should not run this program as root")
		os.Exit(1)
	}

	fmt.Println("Welcome to password generator " + strings.ToUpper(currentUser.Username) + "!")

	// decide if password is random or not
	random, err := getUserInput("Random password? (y/n): ")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if random == "y" {
		// Get password length
		passwordLength, err := getUserInput("Enter password length: ")
		if err != nil && passwordLength == "" {
			fmt.Println(err)
			os.Exit(1)
		}

		// Convert password length to int
		passwordSize, err := convertToInt(passwordLength)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if passwordSize <= 0 || passwordSize > 100 {
			fmt.Println("password length must be greater than 0 and less than 100")
			os.Exit(1)
		}

		// Build password info
		passwordInfo, err := buildPasswordInfoRandom(passwordSize)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create password
		password, err = createPassword(passwordInfo)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Scramble password
		password = scramblePassword(password)

		fmt.Println("Your password is: " + password)

		os.Exit(0)

	}

	// Get password info
	passwordInfo, err := buildPasswordInfo()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create password
	password, err = createPassword(passwordInfo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Scramble password
	password = scramblePassword(password)

	fmt.Println("Your password is: " + password)

	os.Exit(0)
}
