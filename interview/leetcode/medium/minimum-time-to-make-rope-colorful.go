package medium

func minCost(colors string, neededTime []int) int {
	totalCost := 0
	n := len(neededTime)

	for groupStart := 0; groupStart < n; groupStart++ {
		groupEnd := groupStart
		groupSum := 0
		maxTimeInGroup := 0

		for groupEnd < n && colors[groupEnd] == colors[groupStart] {
			groupSum += neededTime[groupEnd]
			if neededTime[groupEnd] > maxTimeInGroup {
				maxTimeInGroup = neededTime[groupEnd]
			}
			groupEnd++
		}

		if groupEnd-groupStart > 1 {
			totalCost += groupSum - maxTimeInGroup
		}

		groupStart = groupEnd - 1
	}

	return totalCost
}
