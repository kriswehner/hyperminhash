package main

import (
	"fmt"
	"math/rand"
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

const maxRange = 5

var sketches []Sketch

func main() {
	sketches = make([]Sketch, maxRange+1)
	for i := 1; i <= maxRange; i++ {
		sketches[int(i)].Hash = hyperminhash.New()
	}

	// for i := 0; i < 10000; i++ {
	// 	newUser := user{UUID: uuid.NewV4().Bytes(), Segments: randomSegmentMembership()}
	// 	for _, segment := range newUser.Segments {
	// 		sketches[segment].Hash.Add(newUser.UUID)
	// 	}
	// }

	for i := 0; i < 10000000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Hash.Add(user)
		sketches[2].Hash.Add(user)
	}

	for i := 0; i < 5000000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Hash.Add(user)

	}

	for i := 0; i < 1000000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Hash.Add(user)
		sketches[3].Hash.Add(user)
	}

	for i := 0; i < 1000000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Hash.Add(user)
		sketches[2].Hash.Add(user)
		sketches[3].Hash.Add(user)
	}

	for i := 0; i < 300000; i++ {
		user := uuid.NewV4().Bytes()
		sketches[1].Hash.Add(user)
		sketches[2].Hash.Add(user)
		sketches[3].Hash.Add(user)
		sketches[4].Hash.Add(user)
	}

	fmt.Printf("Sketch 1 Cardinality: %d\n", sketches[1].Hash.Cardinality())
	fmt.Printf("Sketch 2 Cardinality: %d\n", sketches[2].Hash.Cardinality())
	fmt.Printf("Sketch 3 Cardinality: %d\n", sketches[3].Hash.Cardinality())
	fmt.Printf("Sketch 4 Cardinality: %d\n", sketches[4].Hash.Cardinality())
	fmt.Printf("Sketch 5 Cardinality: %d\n", sketches[5].Hash.Cardinality())

	start := time.Now()
	fmt.Printf("Intersection of Sketches 1 - 4: %v\n", hyperminhash.PolyIntersection([]*hyperminhash.Sketch{sketches[1].Hash, sketches[2].Hash, sketches[3].Hash, sketches[4].Hash}))
	fmt.Printf("Intersection took %s\n", time.Since(start))
	start = time.Now()
	fmt.Printf("Intersection of Sketches 2 - 4 : %d\n", hyperminhash.PolyIntersection([]*hyperminhash.Sketch{sketches[2].Hash, sketches[3].Hash, sketches[4].Hash}))
	fmt.Printf("Intersection took %s\n", time.Since(start))
	start = time.Now()
	fmt.Printf("Intersection of Sketches 1 - 3 : %d\n", hyperminhash.PolyIntersection([]*hyperminhash.Sketch{sketches[2].Hash, sketches[3].Hash, sketches[1].Hash}))
	fmt.Printf("Intersection took %s\n", time.Since(start))
	// fmt.Printf("Intersection of Sketch 1 & 3: %d\n", sketches[1].Hash.Intersection(sketches[3].Hash))

	// fmt.Printf("Union of All 3: %d\n", sketches[1].Hash.Merge(sketches[2].Hash.Merge(sketches[3].Hash)))

	// fmt.Printf("Intersection of Sketch 1, 2 & 3: %d\n", hyperminhash.PolyIntersection([]*hyperminhash.Sketch{sketches[1].Hash, sketches[2].Hash, sketches[3].Hash}))
	// fmt.Printf("Intersection of Sketch 1, 2, 3 & 4: %d\n", hyperminhash.PolyIntersection([]*hyperminhash.Sketch{sketches[1].Hash, sketches[2].Hash, sketches[3].Hash, sketches[4].Hash}))

	// fmt.Printf("Jaccard Index of Sketch 1, 2 & 3: %d\n", sketches[1].Hash.Intersection(sketches[2].Hash.Intersection(sketches[3].Hash))

	// fmt.Printf("Sketch 1 Size: %d\n", unsafe.Sizeof(*sketches[1].Hash))
	// fmt.Printf("Sketch 2 Size: %d\n", unsafe.Sizeof(*sketches[2].Hash))
	// fmt.Printf("Sketch 3 Size: %d\n", unsafe.Sizeof(*sketches[3].Hash))
	// fmt.Printf("Sketch 4 Size: %d\n", unsafe.Sizeof(*sketches[4].Hash))

}

func randomSegmentMembership() (segments []int) {
	for i := 1; i <= maxRange; i++ {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(2) == 1 {
			segments = append(segments, i)
		}
	}
	return
}
