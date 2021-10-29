package solution

import (
	"sort"
)

// Failed, 100% correctness 0% performace
// https://app.codility.com/demo/results/trainingJNZD8R-6UT/
func Solution1(K int, M int, A []int) []int {
	maxSlide := K
	// maxIntVal := M

	lenList := len(A)

	mapCountInt := make(map[int]int)

	for _, val := range A {
		mapCountInt[val]++
	}

	answer := []int{}
	isMarked := make(map[int]bool)
	for i := 0; i < lenList; i++ {
		if i+maxSlide > lenList {
			break
		}

		afterCounterMap := copyMap(mapCountInt)
		for j := 0; j < maxSlide; j++ {
			idx := i + j

			afterCounterMap[A[idx]]--
			afterCounterMap[A[idx]+1]++
		}

		for k, val := range afterCounterMap {
			if val > lenList/2 && !isMarked[k] {
				isMarked[k] = true
				answer = append(answer, k)
				break
			}
		}

	}

	sort.Ints(answer)
	return answer
}

func copyMap(m1 map[int]int) map[int]int {
	m2 := make(map[int]int)
	for k, v := range m1 {
		m2[k] = v
	}
	return m2
}

// Failed, got 20% correctness and 75% performance
// https://app.codility.com/demo/results/trainingBZ8NJZ-KM2/
func Solution2(K int, M int, A []int) []int {
	maxSlide := K
	// maxIntVal := M

	lenList := len(A)
	mapCountInt := make(map[int]int)

	for _, val := range A {
		mapCountInt[val]++
	}

	answer := []int{}
	isMarked := make(map[int]bool)

	for j := 0; j < maxSlide; j++ {
		idx := j

		mapCountInt[A[idx]]--
		mapCountInt[A[idx]+1]++

		// mistake here, should not count before the initial sliding done
		if mapCountInt[A[idx]+1] > lenList/2 && !isMarked[A[idx]+1] {
			isMarked[A[idx]+1] = true
			answer = append(answer, A[idx]+1)
		}
	}

	for i := maxSlide; i < lenList; i++ {
		prevIdx := i - maxSlide
		mapCountInt[A[prevIdx]]++
		mapCountInt[A[prevIdx]+1]--

		currIdx := i
		mapCountInt[A[currIdx]]--
		mapCountInt[A[currIdx]+1]++

		if mapCountInt[A[currIdx]+1] > lenList/2 && !isMarked[A[currIdx]+1] {
			isMarked[A[currIdx]+1] = true
			answer = append(answer, A[currIdx]+1)
		}
	}

	sort.Ints(answer)
	return answer
}

// Failed, 40% correctness & 75% performance
// https://app.codility.com/demo/results/training566VMS-BSF/
func Solution3(K int, M int, A []int) []int {
	maxSlide := K

	lenList := len(A)
	mapCountInt := make(map[int]int)

	for _, val := range A {
		mapCountInt[val]++
	}

	answer := []int{}
	isMarked := make(map[int]bool)

	for j := 0; j < maxSlide; j++ {
		idx := j

		mapCountInt[A[idx]]--
		mapCountInt[A[idx]+1]++

	}

	for k, val := range mapCountInt {
		if val > lenList/2 && !isMarked[k] {
			isMarked[k] = true
			answer = append(answer, k)
			break
		}
	}

	for i := maxSlide; i < lenList; i++ {
		prevIdx := i - maxSlide
		mapCountInt[A[prevIdx]]++
		mapCountInt[A[prevIdx]+1]--

		currIdx := i
		mapCountInt[A[currIdx]]--
		mapCountInt[A[currIdx]+1]++

		// same mistake
		if mapCountInt[A[currIdx]+1] > lenList/2 && !isMarked[A[currIdx]+1] {
			isMarked[A[currIdx]+1] = true
			answer = append(answer, A[currIdx]+1)
		}
	}

	sort.Ints(answer)
	return answer
}

// Accepted
// Same approach with Solution1, the only difference is instead of sliding the array from idx + 1 ... idx + K
// It will start the slider from the first element until idx + K
// Then in the last loop we can just revert back the counter of oldest value and increase the counter of newest value
// https://app.codility.com/demo/results/training64M8X2-8VK/
func Solution4(K int, M int, A []int) []int {
	maxSlide := K

	lenList := len(A)
	mapCountInt := make(map[int]int)

	for _, val := range A {
		mapCountInt[val]++
	}

	answer := []int{}
	isMarked := make(map[int]bool)

	for j := 0; j < maxSlide; j++ {
		idx := j

		mapCountInt[A[idx]]--
		mapCountInt[A[idx]+1]++
	}

	for k, val := range mapCountInt {
		if val > lenList/2 && !isMarked[k] {
			isMarked[k] = true
			answer = append(answer, k)
			break
		}
	}

	for i := maxSlide; i < lenList; i++ {
		prevIdx := i - maxSlide
		mapCountInt[A[prevIdx]]++
		mapCountInt[A[prevIdx]+1]--

		currIdx := i
		mapCountInt[A[currIdx]]--
		mapCountInt[A[currIdx]+1]++

		if mapCountInt[A[currIdx]+1] > lenList/2 && !isMarked[A[currIdx]+1] {
			isMarked[A[currIdx]+1] = true
			answer = append(answer, A[currIdx]+1)
		}

		if mapCountInt[A[prevIdx]] > lenList/2 && !isMarked[A[prevIdx]] {
			isMarked[A[prevIdx]] = true
			answer = append(answer, A[prevIdx])
		}
	}

	sort.Ints(answer)
	return answer
}

// Accepted, using if in last loop
// https://app.codility.com/demo/results/trainingKNPQMP-5B5/
func Solution5(K int, M int, A []int) []int {
	maxSlide := K

	lenList := len(A)
	mapCountInt := make(map[int]int)

	for _, val := range A {
		mapCountInt[val]++
	}

	answer := []int{}
	isMarked := make(map[int]bool)

	for j := 0; j < maxSlide; j++ {
		idx := j

		mapCountInt[A[idx]]--
		mapCountInt[A[idx]+1]++
	}

	for k, val := range mapCountInt {
		if val > lenList/2 && !isMarked[k] {
			isMarked[k] = true
			answer = append(answer, k)
			break
		}
	}

	for i := maxSlide; i < lenList; i++ {
		prevIdx := i - maxSlide
		mapCountInt[A[prevIdx]]++
		mapCountInt[A[prevIdx]+1]--

		currIdx := i
		mapCountInt[A[currIdx]]--
		mapCountInt[A[currIdx]+1]++

		if mapCountInt[A[currIdx]+1] > lenList/2 && !isMarked[A[currIdx]+1] {
			isMarked[A[currIdx]+1] = true
			answer = append(answer, A[currIdx]+1)
		}

		if mapCountInt[A[prevIdx]] > lenList/2 && !isMarked[A[prevIdx]] {
			isMarked[A[prevIdx]] = true
			answer = append(answer, A[prevIdx])
		}
	}

	sort.Ints(answer)
	return answer
}
