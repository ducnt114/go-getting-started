package event_sourcing

type Event interface {
	EventType() string
	Payload() interface{}
}

type AccountCreated struct {
	AccountID string
	Owner     string
}

func (e AccountCreated) EventType() string {
	return "AccountCreated"
}

func (e AccountCreated) Payload() interface{} {
	return e
}

type MoneyDeposited struct {
	AccountID string
	Amount    float64
}

func (e MoneyDeposited) EventType() string {
	return "MoneyDeposited"
}

func (e MoneyDeposited) Payload() interface{} {
	return e
}

type MoneyWithdrawn struct {
	AccountID string
	Amount    float64
}

func (e MoneyWithdrawn) EventType() string {
	return "MoneyWithdrawn"
}

func (e MoneyWithdrawn) Payload() interface{} {
	return e
}

type BankAccount struct {
	AccountID string
	Owner     string
	Balance   float64
}

func (a *BankAccount) Apply(event Event) {
	switch e := event.Payload().(type) {
	case AccountCreated:
		a.AccountID = e.AccountID
		a.Owner = e.Owner
	case MoneyDeposited:
		a.Balance += e.Amount
	case MoneyWithdrawn:
		a.Balance -= e.Amount
	}
}
