package solutions

type ProductOfNumbers struct {
	st     []int
	prefix int
}

func ProductOfNumbersConstructor() ProductOfNumbers {
	return ProductOfNumbers{
		st:     []int{},
		prefix: 1,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.prefix = 1
		this.st = []int{}
	} else {
		this.prefix *= num
		this.st = append(this.st, this.prefix)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.st)
	if k > n {
		return 0
	}

	startIdx, endIdx := n-k-1, n-1
	if startIdx < 0 {
		return this.st[endIdx]
	}

	return this.st[endIdx] / this.st[startIdx]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
