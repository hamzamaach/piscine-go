package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium); i++ {
		for j := i + 1; j < len(podium); j++ {
			if podium[i][0] > podium[j][0] {
				temp := podium[i][0]
				podium[i][0] = podium[j][0]
				podium[j][0] = temp
			}
		}
	}
	return podium
}
