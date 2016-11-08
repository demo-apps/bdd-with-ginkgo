package cart_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shopping Cart Suite")
}

var _ = Describe("Shopping cart", func() {
	Context("initially", func() {
		It("has 0 items", func() {})
		It("has 0 units", func() {})
		Specify("the total amount is 0.00", func() {})
	})

	Context("when a new item is added", func() {
		Context("the shopping cart", func() {
			It("has 1 more unique item than it had earlier", func() {})
			It("has 1 more unit than it had earlier", func() {})
			Specify("total amount increases by item price", func() {})
		})
	})

	Context("when an existing item is added", func() {
		Context("the shopping cart", func() {
			It("has the same number of unique items as earlier", func() {})
			It("has 1 more unit than it had earlier", func() {})
			Specify("total amount increases by item price", func() {})
		})
	})

	Context("that has 0 unit of item A", func() {
		Context("removing item A", func() {
			It("should not change the number of items", func() {})
			It("should not change the number of units", func() {})
			It("should not change the amount", func() {})
		})
	})

	Context("that has 1 unit of item A", func() {
		Context("removing 1 unit item A", func() {
			It("should reduce the number of items by 1", func() {})
			It("should reduce the number of units by 1", func() {})
			It("should reduce the amount by item price", func() {})
		})
	})

	Context("that has 2 units of item A", func() {

		Context("removing 1 unit of item A", func() {
			It("should not reduce the number of items", func() {})
			It("should reduce the number of units by 1", func() {})
			It("should reduce the amount by the item price", func() {})
		})

		Context("removing 2 units of item A", func() {
			It("should reduce the number of items by 1", func() {})
			It("should reduce the number of units by 2", func() {})
			It("should reduce the amount by twice the item price", func() {})
		})
	})
})
