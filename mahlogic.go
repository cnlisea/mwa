package mwa

// 不带癞子是否3N(刻子/顺子)牌型
func Mahjong3NCards(cards *[10]int) bool {
	if cards[0] == 0 {
		return true
	}

	var (
		nextOne int
		nextTwo int
	)
	for i := 1; i < 10; i++ {
		switch cards[i] - nextOne {
		case 1, 4:
			if i > 7 || cards[i+1]-nextTwo == 0 || cards[i+2] == 0 {
				return false
			}
			nextOne, nextTwo = 1+nextTwo, 1
		case 2:
			if i > 7 || cards[i+1]-nextTwo < 2 || cards[i+2] < 2 {
				return false
			}
			nextOne, nextTwo = 2+nextTwo, 2
		default:
			nextOne, nextTwo = nextTwo, 0
		}
	}

	return true
}
