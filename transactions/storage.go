package transactions

import "github.com/pl0q1n/No_More_Flex/models"

type UserID int

type Storage interface {
	Add(uid UserID, transaction models.Transaction) error
	Filter(uid UserID, filter Filter) ([]*models.Transaction, error)
}
