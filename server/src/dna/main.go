package dna

type CPGIslandResult struct {
	Islands   int
	Basepairs int
}

const MIN_CPG_PERCENTAGE = 0.5
const MIN_CPG_OBSERVER_EXPECTED_RATIO = 0.6

func CalculateCPGIslands(dna string, basepairs int) int {
	result := 0

	if len(dna) < basepairs {
		return 0
	}

	counts := make(map[string]int)
	i := 0

	for j, v := range dna {
		counts[string(v)] += 1

		if j > 0 && string(dna[j-1])+string(dna[j]) == "CG" {
			counts["CG"] += 1
		}
		if j-i > basepairs {
			counts[string(dna[i])] -= 1
			if i > 0 && string(dna[i])+string(dna[i+1]) == "CG" {
				counts["CG"] -= 1
			}
			i += 1
		}

		if j-i < basepairs {
			continue
		}

		cpgPercentage := float64((counts["G"] + counts["C"])) / float64(basepairs)
		expectedCpg := float64((counts["C"] * counts["G"])) / float64(basepairs)
		if expectedCpg == 0 {
			continue
		}
		observedToExpectedCpgRatio := float64(counts["CG"]) / expectedCpg

		if cpgPercentage > MIN_CPG_PERCENTAGE && observedToExpectedCpgRatio > MIN_CPG_OBSERVER_EXPECTED_RATIO {
			result += 1
		}

	}

	return result
}

func CalculateCpGIslandsSimultaneously(dnaString string, minBasePairs int, maxBasePairs int, basePairsIncrement int) []CPGIslandResult {
	results := []CPGIslandResult{}
	channel := make(chan CPGIslandResult)
	for basepairs := minBasePairs; basepairs <= maxBasePairs; basepairs += basePairsIncrement {
		go func(pairs int) {
			res := CalculateCPGIslands(dnaString, pairs)
			channel <- CPGIslandResult{Islands: res, Basepairs: pairs}
		}(basepairs)
	}
	for basepairs := minBasePairs; basepairs <= maxBasePairs; basepairs += basePairsIncrement {
		res := <-channel
		results = append(results, res)
	}
	return results
}
