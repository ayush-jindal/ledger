package equity

import "fmt"

func GetPrice(id, exchange string) float32 {
	var price float32
	fmt.Printf("\nEnter the price for %s: ", id)
	fmt.Scan(&price)
	return price
}
