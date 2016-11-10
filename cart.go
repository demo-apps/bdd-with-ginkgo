package cart

// Cart represents the state of a buyer's shopping cart
type Cart struct {
	items map[string]Item
}

// Item represents any item available for sale
type Item struct {
	ID    string
	Name  string
	Price float64
	Qty   int
}

// AddItem adds an item to the cart
func (c *Cart) AddItem(i Item) {}

// RemoveItem removes n number of items with give id from the cart
func (c *Cart) RemoveItem(id string, n int) {}

// TotalAmount returns the total amount of the cart
func (c *Cart) TotalAmount() float64 {
	return 324 // return a random number as a placeholder
}

// TotalUnits returns the total number of units across all items in the cart
func (c *Cart) TotalUnits() int {
	return 9932 // return a random number as a placeholder
}

// TotalUniqueItems returns the number of unique items in the cart
func (c *Cart) TotalUniqueItems() int {
	return 24 // return a random number as a placeholder
}
