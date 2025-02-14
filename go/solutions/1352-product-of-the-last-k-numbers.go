package solutions

type ProductOfNumbers struct {
	st              []int
	lastZeroFoundAt int
}

func ProductOfNumbersConstructor() ProductOfNumbers {
	return ProductOfNumbers{
		st:              []int{},
		lastZeroFoundAt: -1,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	product := num
	if len(this.st) > 0 {
		if this.st[len(this.st)-1] == 0 {
			product *= 1
		} else {
			product *= this.st[len(this.st)-1]
		}
	}

	this.st = append(this.st, product)
	if num == 0 {
		this.lastZeroFoundAt = len(this.st) - 1
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.st)
	if n-k <= this.lastZeroFoundAt {
		return 0
	}

	idx := n - k - 1
	if idx < 0 || this.st[idx] == 0 {
		return this.st[n-1]
	}

	return this.st[n-1] / this.st[idx]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
