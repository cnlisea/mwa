package mwa

// 带癞子是否3N(刻子/顺子)牌型
// 返回需要的癞子数
func MahJong3NCardsRazz(cards *[10]int) int {
	var (
		needRazzNum int
		cardLen     = len(cards)
		nextOne     int
		nextTwo     int
		i           int
	)

	for i = 1; i < cardLen; i++ {
		switch cards[i] - nextOne {
		case 1, 4:
			if i+1 >= cardLen || cards[i+1]-nextTwo == 0 {
				needRazzNum++
				nextOne = nextTwo
			} else {
				nextOne = 1 + nextTwo
			}
			if i+2 >= cardLen || cards[i+2] == 0 {
				needRazzNum++
				nextTwo = 0
			} else {
				nextTwo = 1
			}
		case 2:
			if i+1 >= cardLen || i+2 >= cardLen || cards[i+1]-nextTwo <= 0 || ((cards[i+1]-nextTwo)%3 <= 1 && cards[i+2] <= 1) {
				needRazzNum++
				nextOne, nextTwo = nextTwo, 0
				break
			}

			nextOne, nextTwo = nextTwo+2, 2
		case 0, 3:
			nextOne, nextTwo = nextTwo, 0
		default:
			needRazzNum = needRazzNum + nextOne - cards[i]
			nextOne, nextTwo = nextTwo, 0
		}
	}

	var reverseNeedRazzNum int
	nextOne, nextTwo = 0, 0
	for i = cardLen - 1; i > 0; i-- {
		switch cards[i] - nextOne {
		case 1, 4:
			if i-1 <= 0 || cards[i-1]-nextTwo == 0 {
				reverseNeedRazzNum++
				nextOne = nextTwo
			} else {
				nextOne = 1 + nextTwo
			}
			if i-2 <= 0 || cards[i-2] == 0 {
				reverseNeedRazzNum++
				nextTwo = 0
			} else {
				nextTwo = 1
			}
		case 2:
			if i-1 <= 0 || i-2 <= 0 || cards[i-1]-nextTwo <= 0 || ((cards[i-1]-nextTwo)%3 <= 1 && cards[i-2] <= 1) {
				reverseNeedRazzNum++
				nextOne, nextTwo = nextTwo, 0
				break
			}

			nextOne, nextTwo = nextTwo+2, 2
		case 0, 3:
			nextOne, nextTwo = nextTwo, 0
		default:
			reverseNeedRazzNum = reverseNeedRazzNum + nextOne - cards[i]
			nextOne, nextTwo = nextTwo, 0
		}
	}

	if reverseNeedRazzNum < needRazzNum {
		needRazzNum = reverseNeedRazzNum
	}

	return needRazzNum
}
