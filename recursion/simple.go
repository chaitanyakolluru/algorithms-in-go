package recursion

func Simple(n int) int {
	// base case
	if n == 1 {
		return 1
	}

	// recurse structuring:
	// pre: n + in this case
	// recurse: the recurse function call Simple(n - 1)
	// post which is the base case going back to previous recurse
	// function calls

	// base case should be clear
	// recurse then can be made simpler once we have base cases

	// structuring

	// base case
	// define your base cases clearly

	// actual recursing
	// pre
	// actual recurse function call
	// post

	// edge case
	if n == 0 {
		return 0
	}

	// we shall recurse!
	return n + Simple(n-1)
}
