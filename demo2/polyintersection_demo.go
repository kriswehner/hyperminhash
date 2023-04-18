package main

import (
	"fmt"
	"time"

	"github.com/crodwell/hyperminhash"
	uuid "github.com/satori/go.uuid"
)

type user struct {
	UUID     []byte
	Segments []int
}

type Sketch struct {
	Hash   *hyperminhash.Sketch
	Actual int
}

// Add is a test method
func (sk *Sketch) Add(b []byte) {
	sk.Actual++
	sk.Hash.Add(b)
}

const maxRange = 15

var sketches []Sketch

func main() {
	sketches = make([]Sketch, maxRange+1)
	for i := 1; i <= maxRange; i++ {
		sketches[int(i)].Hash = hyperminhash.New()
	}

	for i := 0; i < 1000000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Add(user)
		sketches[2].Add(user)
	}

	for i := 0; i < 500000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Add(user)

	}

	for i := 0; i < 100000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Add(user)
		sketches[3].Add(user)
	}

	for i := 0; i < 100000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Add(user)
		sketches[2].Add(user)
		sketches[3].Add(user)
	}

	for i := 0; i < 30000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Add(user)
		sketches[2].Add(user)
		sketches[3].Add(user)
		sketches[4].Add(user)
	}

	for i := 0; i < 10000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Add(user)
		sketches[2].Add(user)
		sketches[3].Add(user)
		sketches[4].Add(user)
		sketches[5].Add(user)
	}

	for i := 0; i < 1000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Add(user)
		sketches[2].Add(user)
		sketches[3].Add(user)
		sketches[4].Add(user)
		sketches[5].Add(user)
		sketches[6].Add(user)
		sketches[7].Add(user)
		sketches[8].Add(user)
		sketches[9].Add(user)
		sketches[10].Add(user)
		sketches[11].Add(user)
		sketches[12].Add(user)
		sketches[13].Add(user)
		sketches[14].Add(user)
		sketches[15].Add(user)

	}
	start := time.Now()
	card := sketches[1].Hash.Cardinality()
	fmt.Printf("Set 1 HLL Cardinality: %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", card, sketches[1].Actual, accuracy(card, sketches[1].Actual), time.Since(start))
	start = time.Now()
	card = sketches[2].Hash.Cardinality()
	fmt.Printf("Set 2 HLL Cardinality: %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", card, sketches[2].Actual, accuracy(card, sketches[2].Actual), time.Since(start))
	start = time.Now()
	card = sketches[3].Hash.Cardinality()
	fmt.Printf("Set 3 HLL Cardinality: %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", card, sketches[3].Actual, accuracy(card, sketches[3].Actual), time.Since(start))
	start = time.Now()
	card = sketches[4].Hash.Cardinality()
	fmt.Printf("Set 4 HLL Cardinality: %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", card, sketches[4].Actual, accuracy(card, sketches[4].Actual), time.Since(start))
	start = time.Now()
	card = sketches[5].Hash.Cardinality()
	fmt.Printf("Set 5 HLL Cardinality: %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", card, sketches[5].Actual, accuracy(card, sketches[5].Actual), time.Since(start))
	start = time.Now()
	card = sketches[6].Hash.Cardinality()
	fmt.Printf("Set 6-15 HLL Cardinality: %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", card, sketches[6].Actual, accuracy(card, sketches[6].Actual), time.Since(start))

	start = time.Now()
	sketch1to2ActualIntersection := 140000
	sketch1to2Intersection := hyperminhash.PolyIntersection([]*hyperminhash.Sketch{sketches[2].Hash, sketches[3].Hash})
	fmt.Printf("S2 ∩ S3 : %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", sketch1to2Intersection, sketch1to2ActualIntersection, accuracy(sketch1to2Intersection, sketch1to2ActualIntersection), time.Since(start))

	start = time.Now()
	sketch1to3ActualIntersection := 140000
	sketch1to3Intersection := hyperminhash.PolyIntersection([]*hyperminhash.Sketch{sketches[1].Hash, sketches[2].Hash, sketches[3].Hash})
	fmt.Printf("S1 ∩ S2 ∩ S3 : %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", sketch1to3Intersection, sketch1to3ActualIntersection, accuracy(sketch1to3Intersection, sketch1to3ActualIntersection), time.Since(start))

	start = time.Now()
	sketch2to4ActualIntersection := 40000
	sketch2to4Intersection := hyperminhash.PolyIntersection([]*hyperminhash.Sketch{sketches[2].Hash, sketches[3].Hash, sketches[4].Hash})
	fmt.Printf("S2 ∩ S3 ∩ S4 : %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", sketch2to4Intersection, sketch2to4ActualIntersection, accuracy(sketch2to4Intersection, sketch2to4ActualIntersection), time.Since(start))

	start = time.Now()
	sketch1to4ActualIntersection := 40000
	sketch1to4Intersection := hyperminhash.PolyIntersection([]*hyperminhash.Sketch{sketches[1].Hash, sketches[2].Hash, sketches[3].Hash, sketches[4].Hash})
	fmt.Printf("S1 ∩ S2 ∩ S3 ∩ S4 : %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", sketch1to4Intersection, sketch1to4ActualIntersection, accuracy(sketch1to4Intersection, sketch1to4ActualIntersection), time.Since(start))

	start = time.Now()
	sketch1to5ActualIntersection := 11000
	sketch1to5Intersection := hyperminhash.PolyIntersection([]*hyperminhash.Sketch{sketches[1].Hash, sketches[2].Hash, sketches[3].Hash, sketches[4].Hash, sketches[5].Hash})
	fmt.Printf("S1 ∩ S2 ∩ S3 ∩ S4 ∩ S5 : %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", sketch1to5Intersection, sketch1to5ActualIntersection, accuracy(sketch1to5Intersection, sketch1to5ActualIntersection), time.Since(start))

	start = time.Now()
	sketch1to15ActualIntersection := 1000
	sketch1to15Intersection := hyperminhash.PolyIntersection([]*hyperminhash.Sketch{sketches[1].Hash, sketches[2].Hash, sketches[3].Hash, sketches[4].Hash, sketches[5].Hash,
		sketches[6].Hash, sketches[7].Hash, sketches[8].Hash, sketches[9].Hash, sketches[10].Hash, sketches[11].Hash, sketches[12].Hash, sketches[13].Hash, sketches[14].Hash, sketches[15].Hash})
	fmt.Printf("S1 ∩ S2 ∩ S3 ∩ S4 ∩ S5 ∩ S6 ∩ S7 ∩ S8 ∩ S4 ∩ S9 ∩ S10 ∩ S11 ∩ S12 ∩ S13 ∩ S14 ∩ S15: %d, Actual: %d,  Accuracy: %.2f%%, Execution Time: %s\n", sketch1to15Intersection, sketch1to15ActualIntersection, accuracy(sketch1to15Intersection, sketch1to15ActualIntersection), time.Since(start))

}

func accuracy(a uint64, b int) (acc float64) {
	return ((float64(a) / float64(b)) - 1) * 100
}
