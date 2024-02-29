package piscine

type food struct {
	name     string
	preptime int
}

func FoodDeliveryTime(order string) int {
	prepTime := 0
	if order == "burger" || order == "chips" || order == "nuggets" {
		if order == "burger" {
			prepTime = 15
		} else if order == "chips" {
			prepTime = 10
		} else if order == "nuggets" {
			prepTime = 12
		}
		foodOrder := food{
			name:     order,
			preptime: prepTime,
		}
		return foodOrder.preptime
	} else {
		return 404
	}
}
