package piscine

func Unmatch(a []int) int {
    n := -1
    for _, nb := range a {
        cnt := 0
        for_, com := range a {
            if nb == com {
                cnt++
            }
        }
        if cnt%2 != 0 {
            return nb
        }
    }
    return n
}
