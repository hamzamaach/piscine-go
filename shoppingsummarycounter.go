package piscine

func ShoppingSummaryCounter(str string) map[string]int {
    maps := make(map[string]int)
    var word string

    for _, v := range str {
        if v == ' ' {
            maps[word]++
            word = ""
        } else if v != ' ' {
            word = word + string(v)
        }
    }
    maps[word]++
    return maps
}
