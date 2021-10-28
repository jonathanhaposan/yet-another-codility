package solution

import (
	"fmt"
	"sort"
)

// Solution1 Failed
// https://app.codility.com/demo/results/trainingGTCS7E-VYM/
func Solution1(A []int, B []int, F int) int {
	maxFE := F
	maxBE := len(B) - F

	mapContribFE := make(map[int][]int)
	mapContribBE := make(map[int][]int)

	var listUniqueValFE, listUniqueValBE []int

	for i := 0; i < len(A); i++ {
		if _, exist := mapContribFE[A[i]]; !exist {
			listUniqueValFE = append(listUniqueValFE, A[i])
		}
		mapContribFE[A[i]] = append(mapContribFE[A[i]], i)

		if _, exist := mapContribBE[B[i]]; !exist {
			listUniqueValBE = append(listUniqueValBE, B[i])
		}
		mapContribBE[B[i]] = append(mapContribBE[B[i]], i)
	}

	sort.Ints(listUniqueValFE)
	sort.Ints(listUniqueValBE)

	totalAvailableFE := len(listUniqueValFE)
	totalAvailableBE := len(listUniqueValBE)

	idxTotalAvailableFE := totalAvailableFE - 1
	idxTotalAvailableBE := totalAvailableBE - 1

	fmt.Println("max", maxFE, maxBE)
	fmt.Println("map", mapContribFE, mapContribBE)
	fmt.Println("list", listUniqueValFE, listUniqueValBE)

	totalContrib := 0

	usedIdx := make(map[int]bool)
	for {
		fmt.Println("compare: FE", listUniqueValFE[idxTotalAvailableFE], "> BE", listUniqueValBE[idxTotalAvailableBE])
		if (listUniqueValFE[idxTotalAvailableFE] > listUniqueValBE[idxTotalAvailableBE] && maxFE != 0) || maxBE == 0 {
			fmt.Println("Masuk Atas")

			listIdxFE := mapContribFE[listUniqueValFE[idxTotalAvailableFE]]
			if len(listIdxFE) != 0 {
				if usedIdx[listIdxFE[0]] {
					mapContribFE[listUniqueValFE[idxTotalAvailableFE]] = listIdxFE[1:]
					continue
				}

				totalContrib += A[listIdxFE[0]]
				usedIdx[listIdxFE[0]] = true
				mapContribFE[listUniqueValFE[idxTotalAvailableFE]] = listIdxFE[1:]
				maxFE--
			} else {
				if idxTotalAvailableFE != 0 {
					idxTotalAvailableFE--
				}
			}
		} else {
			fmt.Println("Masuk Bawah")

			listIdxBE := mapContribBE[listUniqueValBE[idxTotalAvailableBE]]
			if len(listIdxBE) != 0 {
				if usedIdx[listIdxBE[0]] {
					mapContribBE[listUniqueValBE[idxTotalAvailableBE]] = listIdxBE[1:]
					continue
				}

				usedIdx[listIdxBE[0]] = true
				totalContrib += B[listIdxBE[0]]
				mapContribBE[listUniqueValBE[idxTotalAvailableBE]] = listIdxBE[1:]
				maxBE--
			} else {
				if idxTotalAvailableBE != 0 {
					idxTotalAvailableBE--
				}
			}
		}

		fmt.Println(totalContrib, maxFE, maxBE)

		if maxBE == 0 && maxFE == 0 {
			break
		}
	}

	return totalContrib
}

// Solution2 Accepted
// https://app.codility.com/demo/results/trainingHFYEE3-VM8/
func Solution2(A []int, B []int, F int) int {
	differenceList := []int{}
	sumAllBE := 0
	for i := 0; i < len(A); i++ {
		differenceList = append(differenceList, B[i]-A[i])
		sumAllBE += B[i]
	}

	sort.Ints(differenceList)

	sumWeakContributionFE := 0
	for i := 0; i < F; i++ {
		sumWeakContributionFE += differenceList[i]
	}

	return sumAllBE - sumWeakContributionFE
}
