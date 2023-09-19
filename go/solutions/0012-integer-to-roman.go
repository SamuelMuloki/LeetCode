package solutions

func IntToRoman(num int) string {
	var ones = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	var tens = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	var hrns = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	var ths = []string{"", "M", "MM", "MMM"}

	return ths[num/1000] + hrns[(num%1000)/100] + tens[(num%100)/10] + ones[num%10]
}
