package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int {
        "dozen": 12,
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen":6,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
        }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int {     
	}
}
// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    value, ok := units[unit]
    if !ok {
    return false
    } else {
    bill[item] += value
	return true
	}
}
// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, okunit := units[unit]
    _, okitem := bill[item]
    	if !okunit || !okitem{
    		return false
    } else {
    Amount := bill[item] - units[unit]
    if Amount < 0 {
        return false
    }else if Amount == 0 {
        delete(bill,item)
        return true
    }else {
		bill[item] = Amount
        return true
    		}
		}
	}    
    

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    itemvalue, okitem := bill[item]
    	if !okitem{
    		return 0,false
    } else {
            return itemvalue,true
    }  
}

