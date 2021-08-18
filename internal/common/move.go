package common

type Move struct {
	From, To    Square
	IsPromotion bool
	Promotion   Role
}
