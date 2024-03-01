package piscine

func ReverseMenuIndex(menu []string) []string {
	reversedMenu := make([]string, len(menu))
	for i := 0; i <= len(menu)-1; i++ {
		reversedMenu[i] = menu[len(menu)-1-i]
	}
	return reversedMenu
}
