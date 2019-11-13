package transactions

import (
	"sync"

	"github.com/pl0q1n/No_More_Flex/models"
)

type MemoryStorage struct {
	mutex *sync.Mutex
	mem   map[UserID][]models.Transaction
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		mutex: &sync.Mutex{},
		mem:   make(map[UserID][]models.Transaction),
	}
}

func CopyTransaction(transaction models.Transaction) models.Transaction {
	return models.Transaction{
		Category: transaction.Category,
		Receiver: &*transaction.Receiver,
		Sender:   &*transaction.Sender,
		Time:     &*transaction.Time,
		Value:    &*transaction.Value,
	}
}

func (storage *MemoryStorage) Add(id UserID, transaction models.Transaction) error {
	storage.mutex.Lock()
	storage.mem[id] = append(storage.mem[id], CopyTransaction(transaction))
	storage.mutex.Unlock()
	return nil
}

func (storage *MemoryStorage) Filter(id UserID, filter Filter) ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)
	storage.mutex.Lock()
	for _, transaction := range storage.mem[id] {
		if filter.Apply(&transaction) {
			t := CopyTransaction(transaction)
			transactions = append(transactions, &t)
		}
	}
	storage.mutex.Unlock()
	return transactions, nil
}
