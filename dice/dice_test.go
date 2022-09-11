package dice_test

import (
	"dice"
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
	"golang.org/x/exp/maps"
)

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
	intervals := []float64{1, 2, 3}
	for m := range intervals {
		var count float64 = 0
		i := intervals[m]
		stdDevMin := mean - (i * stdDev)
		stdDevMax := mean + (i * stdDev)
		for idx := range data {
			f := data[idx]
			if f >= stdDevMin && f <= stdDevMax {
				count += 1
			}
		}
		percentage := (count / float64(len(data))) * 100
		fmt.Println(fmt.Sprintf("percent of frequencies within %v of mean: %v", i, percentage))
	}

}

func diceRollTest(numberOfDie int, numberOfSides int, modifier int, t *testing.T) {
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
}

func indexRollTest(arrayLen int, t *testing.T) {
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
}

func TestD2(t *testing.T) {
	diceRollTest(1, 2, 0, t)
}

func TestD3(t *testing.T) {
	diceRollTest(1, 3, 0, t)
}

func TestD4(t *testing.T) {
	diceRollTest(1, 4, 0, t)
}

func TestD6(t *testing.T) {
	diceRollTest(1, 6, 0, t)
}

func TestD8(t *testing.T) {
	diceRollTest(1, 8, 0, t)
}

func TestD10(t *testing.T) {
	diceRollTest(1, 10, 0, t)
}

func TestD12(t *testing.T) {
	diceRollTest(1, 12, 0, t)
}

func TestD20(t *testing.T) {
	diceRollTest(1, 20, 0, t)
}

func TestD100(t *testing.T) {
	diceRollTest(1, 100, 0, t)
}

func TestD200(t *testing.T) {
	diceRollTest(1, 200, 0, t)
}
func TestModifiedDieRoll(t *testing.T) {
	diceRollTest(2, 10, 3, t)
}

func TestGetIndex(t *testing.T) {
	indexRollTest(200, t)
}
