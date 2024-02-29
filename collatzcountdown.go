package piscine

func CollatzCountdown(start int) int {
	count := 0
	nb := start
	for nb > 0 {
		if nb%2 == 0 {
			nb /= 2
			count++
		} else {
			nb = (nb * 3) + 1
			count++
		}
		if nb == 1 {
			break
		}
	}
	if count == 0 {
		return -1
	} else {
		return count
	}
}
