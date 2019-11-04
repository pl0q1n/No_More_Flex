package restapi

import (
	"sync"

	"github.com/pl0q1n/No_More_Flex/models"
)

const (
	MinInt64 int64 = -9223372036854775808
	MaxInt64 int64 = 9223372036854775807
)

type UserID int

type Storage struct {
	mutex *sync.Mutex
	mem   map[UserID][]models.Transaction
}

func NewStorage() *Storage {
	return &Storage{
		mutex: &sync.Mutex{},
		mem:   make(map[UserID][]models.Transaction),
	}
}

func (storage *Storage) AddTransaction(id UserID, transaction models.Transaction) error {
	storage.mutex.Lock()
	storage.mem[id] = append(storage.mem[id], transaction)
	storage.mutex.Unlock()
	return nil
}

type TransactionsFilter struct {
	Category string
	From     int64
	To       int64
	Receiver string
	Sender   string
}

func NewTransactionsFilter() TransactionsFilter {
	return TransactionsFilter{
		Category: "",
		From:     MinInt64,
		To:       MaxInt64,
		Receiver: "",
		Sender:   "",
	}
}

func (filter *TransactionsFilter) Check(transaction *models.Transaction) bool {
	if filter.Category != "" && transaction.Category != filter.Category {
		return false
	}

	if !(filter.From <= *transaction.Time && *transaction.Time <= filter.To) {
		return false
	}

	if filter.Receiver != "" && *transaction.Receiver != filter.Receiver {
		return false
	}

	if filter.Sender != "" && *transaction.Sender != filter.Sender {
		return false
	}

	return true
}

func (storage *Storage) FilterTransactions(id UserID, filter TransactionsFilter) ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)
	storage.mutex.Lock()
	for _, transaction := range storage.mem[id] {
		if filter.Check(&transaction) {
			transactions = append(transactions, &transaction)
		}
	}
	storage.mutex.Unlock()
	return transactions, nil
}
