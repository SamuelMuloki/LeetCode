package solutions

type numPair struct {
	num  int
	word string
}

var numPairs = []numPair{
	{num: 1000000000, word: "Billion"}, {num: 1000000, word: "Million"}, {num: 1000, word: "Thousand"}, {num: 100, word: "Hundred"},
	{num: 90, word: "Ninety"}, {num: 80, word: "Eighty"}, {num: 70, word: "Seventy"}, {num: 60, word: "Sixty"}, {num: 50, word: "Fifty"},
	{num: 40, word: "Forty"}, {num: 30, word: "Thirty"}, {num: 20, word: "Twenty"}, {num: 19, word: "Nineteen"}, {num: 18, word: "Eighteen"},
	{num: 17, word: "Seventeen"}, {num: 16, word: "Sixteen"}, {num: 15, word: "Fifteen"}, {num: 14, word: "Fourteen"}, {num: 13, word: "Thirteen"},
	{num: 12, word: "Twelve"}, {num: 11, word: "Eleven"}, {num: 10, word: "Ten"}, {num: 9, word: "Nine"}, {num: 8, word: "Eight"}, {num: 7, word: "Seven"},
	{num: 6, word: "Six"}, {num: 5, word: "Five"}, {num: 4, word: "Four"}, {num: 3, word: "Three"}, {num: 2, word: "Two"}, {num: 1, word: "One"},
}

func NumberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	for i := range numPairs {
		if num >= numPairs[i].num {
			ans := numPairs[i].word
			if num >= 100 {
				ans = NumberToWords(num/numPairs[i].num) + " " + ans
			}

			if num%numPairs[i].num != 0 {
				ans += " " + NumberToWords(num%numPairs[i].num)
			}

			return ans
		}
	}

	return ""
}
