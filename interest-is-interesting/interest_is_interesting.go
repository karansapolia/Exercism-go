package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return 3.213
	} else if balance >= 0 && balance < 1000 {
		return 0.5
	} else if balance >= 1000 && balance < 5000 {
		return 1.621
	} else {
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) * 0.01
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
// I = PRT / 100
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + balance*float64(InterestRate(balance))*0.01
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	totalBalance := balance
	years := 0
	for ; totalBalance < targetBalance; years++ {
		totalBalance += totalBalance * float64(InterestRate(totalBalance)) * 0.01
	}
	return years
}
