package solutions

import "strconv"

func GetLucky(s string, k int) int {
    str := ""
    for i := range s {
        pos := strconv.Itoa(int(s[i]-'a') + 1)
        str += pos
    }

    str1, ans := str, 0
    for k > 0 {
        ans = 0
        for i := range str1 {
            toInt, _ := strconv.Atoi(string(str1[i]))
            ans += toInt
        }
        str1 = strconv.Itoa(ans)
        k--
    }

    return ans
}
