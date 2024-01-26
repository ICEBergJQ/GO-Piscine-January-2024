package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	str := food{}
	if order == "nuggets" {
		str.preptime = 12
	} else if order == "burger" {
		str.preptime = 15
	} else if order == "chips" {
		str.preptime = 10
	} else {
		str.preptime = 404
	}
	return str.preptime
}
