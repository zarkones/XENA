package main

import (
	"math/rand"
)

/*
	1. Divide payload into multiple chunks of random size.
	2. Chose identifier for each chunk.
*/
func NaturalTrafficOut(rootSeed int64, payload string) map[string]string {
	rand.Seed(rootSeed)
	var jsonObject = make(map[string]string)

	// Assume the chunks count.
	var chunksMaxCount int = rand.Intn(50) + 3
	for chunksMaxCount >= len(payload) || chunksMaxCount%2 != 0 {
		chunksMaxCount = rand.Intn(50) + 3
	}

	chunkSize := len(payload) / chunksMaxCount
	payloadIndex := 0

	for chunkIndex := 0; chunkIndex < chunksMaxCount; chunkIndex++ {
		rand.Seed(rootSeed + int64(chunkIndex) + 1)
		identifier := EnglishCommon[rand.Intn(len(EnglishCommon))]

		nextChunkStart := payloadIndex + chunkSize
		if payloadIndex >= len(payload) {
			jsonObject[identifier] = payload[payloadIndex:]
		} else {
			jsonObject[identifier] = payload[payloadIndex:nextChunkStart]
		}

		if payloadIndex == 0 {
			payloadIndex = chunkSize
		} else {
			payloadIndex = nextChunkStart
		}
	}

	return jsonObject
}
