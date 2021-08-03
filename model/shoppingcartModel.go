package model

type ShoppingCartModel interface {
	Get() []Shopping_cart
	Insert(Shopping_cart) error
}
