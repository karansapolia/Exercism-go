package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitsStore := map[string]int{}
	unitsStore["quarter_of_a_dozen"] = 3
	unitsStore["half_of_a_dozen"] = 6
	unitsStore["dozen"] = 12
	unitsStore["small_gross"] = 120
	unitsStore["gross"] = 144
	unitsStore["great_gross"] = 1728

	return unitsStore
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, compatibleUnit := units[unit]
	if compatibleUnit {
		bill[item] += Units()[unit]
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, validUnit := units[unit]
	_, itemExists := bill[item]
	if !validUnit || !itemExists {
		return false
	} else {
		if Units()[unit] > bill[item] {
			return false
		} else if Units()[unit] == bill[item] {
			delete(bill, item)
			return true
		} else {
			bill[item] -= Units()[unit]
			return true
		}
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, present := bill[item]
	return value, present
}
