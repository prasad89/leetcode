func getSumOfSqaures(n int) int {
	total := 0
	for n != 0 {
		d := n % 10
		total += d * d
		n = n / 10
	}
	return total
}

func isHappy(n int) bool {
	slow := n
	fast := getSumOfSqaures(n)
	for slow != fast {
		slow = getSumOfSqaures(slow)
		fast = getSumOfSqaures(getSumOfSqaures(fast))
	}
	return slow == 1
}
