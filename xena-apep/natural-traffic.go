package main

import (
	"math"
	"math/rand"
)

/*
	1. Divide payload into multiple chunks of random size.
	2. Chose identifier for each chunk.
*/
func NaturalTrafficOut(rootSeed int64, payload string) map[string]string {
	maxIdentifiers := 16

	// Output.
	var jsonObject = make(map[string]string)

	rand.Seed(rootSeed)

	// Assume the chunks count.
findCorrectChunkCount:
	var chunksMaxCount int = rand.Intn(maxIdentifiers) + 3
	for chunksMaxCount >= len(payload)/2 || chunksMaxCount%2 != 0 {
		chunksMaxCount = rand.Intn(maxIdentifiers) + 3
	}

	// Length of each chunk.
	var chunkLength float64 = float64(len(payload)) / float64(chunksMaxCount)

	_, isValidChunkCount := math.Modf(chunkLength)
	if isValidChunkCount != 0 {
		goto findCorrectChunkCount
	}

	payloadIndex := 0

	for chunkIndex := 0; chunkIndex < chunksMaxCount; chunkIndex++ {
		// Generate an identifier for this chunk.
		rand.Seed(rootSeed + int64(chunkIndex) + 1)
		identifier := EnglishCommon[rand.Intn(len(EnglishCommon))]

		var chunkEnd int = payloadIndex + int(chunkLength)

		if payloadIndex >= len(payload) {
			jsonObject[identifier] = payload[payloadIndex:]
		} else {
			jsonObject[identifier] = payload[payloadIndex:chunkEnd]
		}

		payloadIndex = chunkEnd
	}

	return jsonObject
}
