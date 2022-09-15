package dice_test

import (
	"dice"
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/montanaflynn/stats"
	"golang.org/x/exp/maps"
)

func percentDifference(a float64, b float64) float64 {
	mean, _ := stats.Mean([]float64{a, b})
	diff := math.Abs(a - b)
	return (diff / mean) * 100
}

func empiricalRuleTest(data []float64) {
	min, _ := stats.Min(data)
	max, _ := stats.Max(data)
	mean, _ := stats.Mean(data)
	median, _ := stats.Median(data)
	stdDev, _ := stats.StandardDeviation(data)
	skew := (3 * (mean - median)) / stdDev
	var kurtosisNumerator float64 = 0
	for idx := range data {
		f := data[idx]
		kurtosisNumerator += math.Pow((f - mean), 4)
	}
	kurtosis := (kurtosisNumerator / (float64(len(data)) * math.Pow(stdDev, 4))) - 3
	message := fmt.Sprintf("Stats:\nMin: %v Max: %v Mean: %v Median: %v StdDev: %v Skew: %v Kurtosis: %v",
		min, max, mean, median, stdDev, skew, kurtosis)
	fmt.Println(message)
	fmt.Println("68-95-99.7, aka empirical, rule checks")
	intervals := map[float64]float64{1: 64, 2: 95, 3: 99.7}
	// go intentionally randomizes the order that range iterates maps.
	keys := make([]float64, 0)
	for k, _ := range intervals {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	for _, k := range keys {
		var count float64 = 0
		stdDevMin := mean - (k * stdDev)
		stdDevMax := mean + (k * stdDev)
		for idx := range data {
			f := data[idx]
			if f >= stdDevMin && f <= stdDevMax {
				count += 1
			}
		}
		percentage := (count / float64(len(data))) * 100
		pctDiff := percentDifference(intervals[k], percentage)
		fmt.Println(fmt.Sprintf("percent of frequencies within %v standard deviations of mean: %v(%.2f)", k, percentage, pctDiff))
	}

}

func diceRollTest(numberOfDie int, numberOfSides int, modifier int, t *testing.T) bool {
	fmt.Println(fmt.Sprintf("Testing %dD%d+ %d", numberOfDie, numberOfSides, modifier))
	rolls := []float64{}
	frequency := make(map[int]float64)
	for x := 0; x < 10000; x++ {
		roll := dice.Roll(numberOfDie, numberOfSides, modifier)
		if roll < (numberOfDie + modifier) {
			t.Fatalf("Invalid roll! Got %d on %d sided die", roll, numberOfSides)
		}
		if roll > ((numberOfDie * numberOfSides) + modifier) {
			t.Fatalf("Invalid roll! Got %d on %d sided die", roll, numberOfSides)
		}
		rolls = append(rolls, float64(roll))
		frequency[roll] += 1
	}
	max, _ := stats.Max(rolls)
	min, _ := stats.Min(rolls)
	median, _ := stats.Median(rolls)
	mean, _ := stats.Mean(rolls)
	message := fmt.Sprintf("Dice: %d Sides: %d Modifier: %d Min: %v Max:%v Mean: %v Median: %v", numberOfDie, numberOfSides, modifier, min, max, mean, median)
	fmt.Println(message)
	freqs := maps.Values(frequency)
	fmt.Print("Frequency ")
	empiricalRuleTest(freqs)
	return true
}

func indexRollTest(arrayLen int, t *testing.T) bool {
	fmt.Println(fmt.Sprintf("Testing random index of %d items", arrayLen))
	rolls := []float64{}
	frequency := make(map[int]float64)
	for x := 0; x < 10000; x++ {
		roll := dice.GetRandomIndex(arrayLen)
		if roll < 0 {
			t.Fatalf("Invalid roll! Got %d on %d sided die", roll, arrayLen)
		}
		if roll > arrayLen-1 {
			t.Fatalf("Invalid roll! Got %d on %d sided die", roll, arrayLen)
		}
		rolls = append(rolls, float64(roll))
		frequency[roll] += 1
	}
	max, _ := stats.Max(rolls)
	min, _ := stats.Min(rolls)
	median, _ := stats.Median(rolls)
	mean, _ := stats.Mean(rolls)
	message := fmt.Sprintf("Sides: %d Min: %v Max:%v Mean: %v Median: %v", arrayLen, min, max, mean, median)
	fmt.Println(message)
	freqs := maps.Values(frequency)
	fmt.Print("Frequency ")
	empiricalRuleTest(freqs)
	return true
}

func elementRollTest(array []any, t *testing.T) bool {
	fmt.Println(fmt.Sprintf("Testing random index of %d items", array))
	frequency := make(map[any]float64)
	for x := 0; x < 10000; x++ {
		roll := dice.GetRandomElement(array)
		if roll == nil {
			t.Fatalf("Invalid roll! Got %d on %d sided die", roll, array)
		}
		frequency[roll] += 1
	}
	freqs := []float64{}
	for _, value := range frequency {
		freqs = append(freqs, value)
	}
	fmt.Print("Frequency ")
	empiricalRuleTest(freqs)
	return true
}

func TestDieRoll(t *testing.T) {
	tests := []struct {
		number   int
		sides    int
		modifier int
	}{
		{1, 2, 0},
		{1, 3, 0},
		{1, 4, 0},
		{1, 6, 0},
		{1, 8, 0},
		{1, 10, 0},
		{2, 10, 3},
		{1, 12, 0},
		{1, 20, 0},
		{1, 100, 0},
		{1, 200, 0},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v, %v, %v", tc.number, tc.sides, tc.modifier), func(t *testing.T) {
			if !diceRollTest(tc.number, tc.sides, tc.modifier, t) {
				t.Fatalf("Test Failed for %v, %v, %v", tc.number, tc.sides, tc.modifier)
			}
			fmt.Println("#############################################")
		})
	}
}

func TestGetIndex(t *testing.T) {
	tests := []struct {
		number int
	}{
		{200},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.number), func(t *testing.T) {
			if !indexRollTest(tc.number, t) {
				t.Fatalf("Test Failed for %v", tc.number)
			}
			fmt.Println("#############################################")
		})
	}
	indexRollTest(200, t)
}

func TestGetElement(t *testing.T) {
	tests := []struct {
		elements []interface{}
	}{
		{elements: []interface{}{24, 40, 46, 36, 12, 39, 16, 47, 41, 49, 25, 26, 50, 38, 13, 28, 19, 14, 35, 30}},
		{elements: []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.elements), func(t *testing.T) {
			if !elementRollTest(tc.elements, t) {
				t.Fatalf("Test failed for %v", tc.elements)
			}
			fmt.Println("#############################################")
		})
	}
}
