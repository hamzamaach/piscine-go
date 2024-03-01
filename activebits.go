package piscine

func ActiveBits(number int) int {
	activeBits := 0
	for number > 0 {
		if number%2 == 1 {
			activeBits++
		}
		number /= 2
	}
	return activeBits
}
