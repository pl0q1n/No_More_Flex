package wallet

// Wallet is structure that holds most of money-related info for each user
type Wallet struct {
	mounthlyCredits map[string]int
	income          int
	balance         int
	bankAccounts    []string
}
