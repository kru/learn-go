package core

// different ways to climb stairs of n with 1,2,3 step respectively
// example:
//   n = 3 => 1,1,1
//            2,1 1,2 3: answer 4
//   n = 4 =>  1,1,1,1
//             2,2
//             2,1,1
//             1,2,1
//             1,1,2
//             3,1
//             1,3: answer 7

// n = 5 answer: 13
func ClimbStairs(n int64) int64 {
	var count int64
	if n%3 >= 0 {
		count += 1
		if n/3 == 1 {
			count += 2
		}
		if n/3 == 2 {
			count += 2
		}
	}
	if n%2 == 0 {
		count += 1
	}

	return count + 1
}
