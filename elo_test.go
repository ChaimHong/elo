package elo

import (
	"fmt"
	"testing"
)

type KStruct struct{}

var KI = KStruct{}

func (k KStruct) GetK(points float64) float64 {
	// 普通级别
	if points < 2300 {
		return 32
	}

	// 大师级别
	return 16
}

func TestElo(t *testing) {
	var winner, loser int16 = 1200, 1000
	wp, lp := Elo(KI, 1200, 1000)
	fmt.Println(wp, lp)
}
