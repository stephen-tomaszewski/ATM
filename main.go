package main

import (
	"fmt"
	"io"
	"os"
)

type atmState int

const (
	invalidATMState = atmState(iota)
	needCardATMState
	mainScreenATMState
	// TODO: Fill in other states
)

// Bank will represent a bank that contains a list of accounts associated with it
type Bank struct {
	accounts map[string]Account
}

// GetAccount will return a given account based on the card number. If the card numbers format is
// invalid or the number was not in the bank's records, it will return an error
func (b *Bank) GetAccount(cardNumber string) (Account, error) {
	if err := validateCardNumber(cardNumber); err != nil {
		return Account{}, err
	}

	acct, ok := b.accounts[cardNumber]
	if !ok {
		return Account{}, fmt.Errorf("No account is associated with that card number")
	}

	return acct, nil
}

// validateCardNumber will ensure that the card number is of format XXXX-XXXX-XXXX-XXXX and non '-' are
// numbers.
func validateCardNumber(cardNumber string) error {
	return fmt.Errorf("TODO: not implemented")
}

// TODO: Implement Account structure

// ATM is a simulated ATM machine that can read a card and do various actions that the card is related to
type ATM struct {
	out   io.Writer
	in    io.Reader
	state atmState
	bank  *Bank

	cardNumber *string
}

// NewATM constructs a new ATM machine with a given display and input.
func NewATM(out io.Writer, in io.Reader, bank *Bank) *ATM {
	return &ATM{
		out:   out,
		in:    in,
		bank:  bank,
		state: needCardATMState,
	}
}

// Run will output the given display based on the given ATM state
func (atm *ATM) Run() {
	switch atm.state {
	case needCardATMState:
		fmt.Fprintf(atm.out, "Welcome to Awesome Bank\nPlease enter a card number:\n")

		cardNumber := ""
		fmt.Fscanf(atm.in, "%s", &cardNumber)
		atm.cardNumber = &cardNumber
		atm.state = mainScreenATMState
	case mainScreenATMState:
		// TODO: Implement validate card.
		// Should be a valid card number of XXXX-XXXX-XXXX-XXXX and should only contain numbers
		// TODO: Call the proper method to get the Account
		// TODO: Write to the display what the user may want to do
		// TODO: Read the input and transition state.
	}
}

func main() {
	bank := &Bank{
		accounts: map[string]Account{
			// TODO: Add more accounts here and populate the accounts with a balance
			"0000-0000-0000-0000": Account{},
		},
	}

	atm := NewATM(os.Stdout, os.Stdin, bank)
	for {
		atm.Run()
	}
}
