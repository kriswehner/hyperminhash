package main

import (
	"fmt"
	"math/rand"

	"github.com/spaolacci/murmur3"
)

const (
	numHashes = 128
)

type MinHash struct {
	hashes []uint64
}

func NewMinHash() *MinHash {
	hashes := make([]uint64, numHashes)
	for i := 0; i < numHashes; i++ {
		hashes[i] = rand.Uint64()
	}
	return &MinHash{hashes}
}

func (m *MinHash) Compute(set []string) []uint64 {
	signatures := make([]uint64, numHashes)
	for i := 0; i < numHashes; i++ {
		minHash := ^uint64(0)
		for _, s := range set {
			hasher := murmur3.New64WithSeed(uint32(i))
			hasher.Write([]byte(s))
			hash := hasher.Sum64()
			if hash < minHash {
				minHash = hash
			}
		}
		signatures[i] = minHash
	}
	return signatures
}

func Jaccard(set1, set2 []string) float64 {
	if len(set1) == 0 && len(set2) == 0 {
		return 1.0
	}
	set1Map := make(map[string]bool)
	set2Map := make(map[string]bool)
	union := make(map[string]bool)
	for _, s := range set1 {
		set1Map[s] = true
		union[s] = true
	}
	for _, s := range set2 {
		set2Map[s] = true
		union[s] = true
	}
	intersection := 0
	for s := range union {
		if set1Map[s] && set2Map[s] {
			intersection++
		}
	}
	return float64(intersection) / float64(len(union))
}

func IntersectionMinHash(sets [][]string) float64 {
	if len(sets) < 2 {
		return 0.0
	}
	minhashes := make([][]uint64, len(sets))
	for i := range sets {
		mh := NewMinHash()
		minhashes[i] = mh.Compute(sets[i])
	}
	intersection := 0
	for i := 0; i < numHashes; i++ {
		hashValues := make([]uint64, len(sets))
		for j, mh := range minhashes {
			hashValues[j] = mh[i]
		}
		if len(hashValues) != len(sets) {
			continue
		}
		allEqual := true
		for j := 1; j < len(hashValues); j++ {
			if hashValues[j] != hashValues[0] {
				allEqual = false
				break
			}
		}
		if allEqual {
			intersection++
		}
	}
	jaccard := Jaccard(sets[0], sets[1])
	for i := 2; i < len(sets); i++ {
		jaccard = Jaccard(sets[i], sets[0]) // sets[0] should be union
	}
	return float64(intersection) / float64(numHashes) / jaccard
}

func main() {
	set1 := []string{"apple", "banana", "orange"}
	set2 := []string{"apple", "kiwi", "pear"}
	set3 := []string{"banana", "orange", "watermelon"}
	sets := [][]string{set1, set2, set3}
	intersection := IntersectionMinHash(sets)
	fmt.Println("Intersection:", intersection)
}
