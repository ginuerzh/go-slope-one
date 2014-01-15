// slope_one
package slopeone

type SlopeOne struct {
	diffMatrix map[string]map[string]float32
	freqMatrix map[string]map[string]int
}

func NewSlopeOne(users []map[string]float32) *SlopeOne {
	so := &SlopeOne{}
	so.diffMatrix = make(map[string]map[string]float32)
	so.freqMatrix = make(map[string]map[string]int)

	so.buildDiffMatrix(users)
	return so
}

/**
 * Based on existing data, and using weights,
 * try to predict all missing ratings.
 * The trick to make this more scalable is to consider
 * only mDiffMatrix entries having a large  (>1) mFreqMatrix
 * entry.
 *
 * It will output the prediction 0 when no prediction is possible.
 */
func (so *SlopeOne) Predict(user map[string]float32) map[string]float32 {
	predictions := make(map[string]float32)
	frequencies := make(map[string]int)

	for item, _ := range so.diffMatrix {
		frequencies[item] = 0
		predictions[item] = 0
	}

	for itemj, rating := range user {
		for itemk, _ := range so.diffMatrix {
			if _, ok := so.diffMatrix[itemk][itemj]; !ok {
				continue
			}

			newVal := (so.diffMatrix[itemk][itemj] + rating) * float32(so.freqMatrix[itemk][itemj])
			predictions[itemk] = predictions[itemk] + newVal
			frequencies[itemk] = frequencies[itemk] + so.freqMatrix[itemk][itemj]
		}
	}

	cleanPredictions := make(map[string]float32)
	for item, pred := range predictions {
		if freq := frequencies[item]; freq > 0 {
			cleanPredictions[item] = pred / float32(freq)
		}
	}
	for item, pred := range user {
		cleanPredictions[item] = pred
	}

	return cleanPredictions
}

func (so *SlopeOne) buildDiffMatrix(users []map[string]float32) {
	diffMatrix := so.diffMatrix
	freqMatrix := so.freqMatrix

	// first iterate through users
	for _, user := range users {
		// then iterate through user data
		for item, rating := range user {
			if _, ok := diffMatrix[item]; !ok {
				diffMatrix[item] = make(map[string]float32)
				freqMatrix[item] = make(map[string]int)
			}

			for item2, rating2 := range user {
				oldCount := 0
				oldDiff := float32(0.0)

				if freq, ok := freqMatrix[item][item2]; ok {
					oldCount = freq
				}
				if diff, ok := diffMatrix[item][item2]; ok {
					oldDiff = diff
				}
				observedDiff := rating - rating2
				freqMatrix[item][item2] = oldCount + 1
				diffMatrix[item][item2] = oldDiff + observedDiff
			}
		}
	}
	for itemj, diffMap := range diffMatrix {
		for itemi, diff := range diffMap {
			oldValue := diff
			count := freqMatrix[itemj][itemi]
			diffMatrix[itemj][itemi] = oldValue / float32(count)
		}
	}
}
