package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customerName := strings.ToUpper(customer)
	message := "Welcome to the Tech Palace, " + customerName
	return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	message := border + "\n" + welcomeMsg + "\n" + border
	return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	message := strings.ReplaceAll(oldMsg, "*", "")
	formattedMsg := strings.TrimSpace(message)
	return formattedMsg
}
