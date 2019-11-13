package transactions

import"github.com/pl0q1n/No_More_Flex/models"


const (
	MinInt64 int64 = -9223372036854775808
	MaxInt64 int64 = 9223372036854775807
)


type Filter struct {
	Category string
	From     int64
	To       int64
	Receiver string
	Sender   string
}

func NewFilter() Filter {
	return Filter{
		Category: "",
		From:     MinInt64,
		To:       MaxInt64,
		Receiver: "",
		Sender:   "",
	}
}

// Returns true if transaction passes the filter
func (filter *Filter) Apply(transaction *models.Transaction) bool {
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

