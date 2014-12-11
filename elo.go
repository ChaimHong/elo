package elo

import (
	"lms/lib/assert"
)

type EloI interface {
	GetK(float64) float64
}

const INT16MAX = 0x7fff

func Elo(data EloI, w, l int16) (winnerN, loserN int16) {
	winner, loser := float64(w), float64(l)

	var EWinner, Eloser float64

	if winner >= loser {
		EWinner = getE(winner - loser)
		Eloser = 1 - EWinner
	} else {
		Eloser = getE(loser - winner)
		EWinner = 1 - Eloser
	}

	// Rn = Ro + K * (Sa - Ea)
	// 新的积分 ＝ 旧的积分 ＋ k * (得分 － 期望)

	// 四舍五入 (+0.5)
	wn := winner + data.GetK(winner)*(1-EWinner) + 0.5
	assert.True2(wn <= INT16MAX, "the value is too big")

	ln := loser + data.GetK(loser)*(0-Eloser) + 0.5
	assert.True2(ln <= INT16MAX, "the value is too big")

	return int16(wn), int16(ln)
}

// 根据积分差获得elo胜率期望
// 对应的公式：1 / (1 + math.Pow(10, float64(loser-winner)/400))
func getE(d float64) float64 {
	if d <= 3 {
		return 0.50
	}
	if d <= 10 {
		return 0.51
	}
	if d <= 17 {
		return 0.52
	}
	if d <= 25 {
		return 0.53
	}
	if d <= 32 {
		return 0.54
	}
	if d <= 39 {
		return 0.55
	}
	if d <= 46 {
		return 0.56
	}
	if d <= 53 {
		return 0.57
	}
	if d <= 61 {
		return 0.58
	}
	if d <= 68 {
		return 0.59
	}
	if d <= 76 {
		return 0.60
	}
	if d <= 83 {
		return 0.61
	}
	if d <= 91 {
		return 0.62
	}
	if d <= 98 {
		return 0.63
	}
	if d <= 106 {
		return 0.64
	}
	if d <= 113 {
		return 0.65
	}
	if d <= 121 {
		return 0.66
	}
	if d <= 129 {
		return 0.67
	}
	if d <= 137 {
		return 0.68
	}
	if d <= 145 {
		return 0.69
	}
	if d <= 153 {
		return 0.70
	}
	if d <= 162 {
		return 0.71
	}
	if d <= 170 {
		return 0.72
	}
	if d <= 179 {
		return 0.73
	}
	if d <= 188 {
		return 0.74
	}
	if d <= 197 {
		return 0.75
	}
	if d <= 206 {
		return 0.76
	}
	if d <= 215 {
		return 0.77
	}
	if d <= 225 {
		return 0.78
	}
	if d <= 235 {
		return 0.79
	}
	if d <= 245 {
		return 0.80
	}
	if d <= 256 {
		return 0.81
	}
	if d <= 267 {
		return 0.82
	}
	if d <= 278 {
		return 0.83
	}
	if d <= 290 {
		return 0.84
	}
	if d <= 302 {
		return 0.85
	}
	if d <= 315 {
		return 0.86
	}
	if d <= 328 {
		return 0.87
	}
	if d <= 344 {
		return 0.88
	}
	if d <= 357 {
		return 0.89
	}
	if d <= 374 {
		return 0.90
	}
	if d <= 391 {
		return 0.91
	}
	if d <= 411 {
		return 0.92
	}
	if d <= 432 {
		return 0.93
	}
	if d <= 456 {
		return 0.94
	}
	if d <= 484 {
		return 0.95
	}
	if d <= 517 {
		return 0.96
	}
	if d <= 559 {
		return 0.97
	}
	if d <= 619 {
		return 0.98
	}
	if d <= 735 {
		return 0.99
	}
	return 1
}
