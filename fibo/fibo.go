package fibo

func Fibonacci() func() int {
	prev, next := 0, 1
	return func() int {
		ans:=prev
		prev, next = next, prev+ next
		return ans
	}
}
