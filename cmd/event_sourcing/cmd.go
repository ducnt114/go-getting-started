package event_sourcing

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "event_sourcing",
	Short: "event_sourcing",
	Long:  `event_sourcing`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
	// Initialize event store
	store := NewEventStore()

	// Create events
	accountCreated := AccountCreated{AccountID: "12345", Owner: "ducnt"}
	moneyDeposited := MoneyDeposited{AccountID: "12345", Amount: 500}
	moneyWithdrawn := MoneyWithdrawn{AccountID: "12345", Amount: 200}

	// Append events to the store
	store.Append(accountCreated)
	store.Append(moneyDeposited)
	store.Append(moneyWithdrawn)

	// Reconstruct account state from events
	account := &BankAccount{}
	for _, event := range store.GetAllEvents() {
		account.Apply(event)
	}

	// Display the final state
	fmt.Printf("AccountID: %s, Owner: %s, Balance: %.2f\n",
		account.AccountID, account.Owner, account.Balance)
}
