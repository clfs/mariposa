package common

type Move struct {
	from, to    Square
	isPromotion bool
	promotion   Role
}

func (m Move) From() Square {
	return m.from
}

func (m Move) To() Square {
	return m.to
}

func (m Move) IsPromotion() bool {
	return m.isPromotion
}

func (m Move) PromotionTo() (Role, bool) {
	return m.promotion, m.isPromotion
}
