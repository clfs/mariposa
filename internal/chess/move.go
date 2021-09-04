package chess

// 6: from
// 6: to
// 3: piece captured
// 2: promotion
type Move uint32

const (
	moveFromMask        = 0b0000_0000_0000_0000_0000_0000_0011_1111
	moveToMask          = 0b0000_0000_0000_0000_0000_1111_1100_0000
	movePromotionMask   = 0b0000_0000_0000_0000_0111_0000_0000_0000
	moveIsPromotionMask = 0b0000_0000_0000_0000_1000_0000_0000_0000
	moveCaptureMask     = 0b0000_0000_0000_0111_0000_0000_0000_0000
	moveIsCaptureMask   = 0b0000_0000_0000_1000_0000_0000_0000_0000

	moveFromShift      = 0
	moveToShift        = 6
	movePromotionShift = 12
	moveCaptureShift   = 16
)

func (m Move) From() Square {
	return Square(m & moveFromMask >> moveFromShift)
}

func (m Move) To() Square {
	return Square(m & moveToMask >> moveToShift)
}

func (m Move) IsPromotion() bool {
	return m&moveIsPromotionMask != 0
}

func (m Move) Promotion() (Role, bool) {
	return Role(m & movePromotionMask >> movePromotionShift), m.IsPromotion()
}

func (m Move) IsCapture() bool {
	return m&moveIsCaptureMask != 0
}

func (m Move) Capture() (Role, bool) {
	return Role(m & moveCaptureMask >> moveCaptureShift), m.IsCapture()
}
