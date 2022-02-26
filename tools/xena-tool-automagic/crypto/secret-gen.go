package crypto

import (
	"fmt"
	"main/data"
	"math/rand"
	"time"
)

func UniqueSecret() string {
	seed := time.Now().UnixNano()
	mili := time.Now().UnixMilli()

	secret := data.RandomPopularWordBySeed(seed + 12%mili + time.Now().UnixNano()/3)
	rand.Seed(time.Now().UnixNano())
	secret += fmt.Sprint(rand.Intn(12798093))
	secret += data.RandomPopularWordBySeed(seed - 24%mili + time.Now().UnixNano()/3)
	secret += data.RandomPopularWordBySeed(seed/68%mili + time.Now().UnixNano()/3)
	secret += data.RandomPopularWordBySeed(seed*3%mili + time.Now().UnixNano()/3)
	secret += data.RandomPopularWordBySeed(seed - 20%mili + time.Now().UnixNano()/3)
	secret += data.RandomPopularWordBySeed(seed/22%mili + time.Now().UnixNano()/3)
	rand.Seed(time.Now().UnixNano())
	secret += fmt.Sprint(rand.Intn(56689707))
	secret += data.RandomPopularWordBySeed(seed + 12%mili + time.Now().UnixNano()/3)
	secret += data.RandomPopularWordBySeed(seed - 24%mili + time.Now().UnixNano()/3)
	secret += data.RandomPopularWordBySeed(seed/68%mili + time.Now().UnixNano()/3)
	secret += data.RandomPopularWordBySeed(seed*3%mili + time.Now().UnixNano()/3)
	rand.Seed(time.Now().UnixNano())
	secret += fmt.Sprint(rand.Intn(59679807))
	secret += data.RandomPopularWordBySeed(seed - 20%mili + time.Now().UnixNano()/3)
	secret += data.RandomPopularWordBySeed(seed - 2622%mili + time.Now().UnixNano()/3)
	secret += data.RandomPopularWordBySeed(seed + 28672%mili + time.Now().UnixNano()/3)
	secret += fmt.Sprint(rand.Intn(78665765))
	secret += data.RandomPopularWordBySeed(seed - 678786%mili + time.Now().UnixNano()/3)
	secret += data.RandomPopularWordBySeed(seed/13%mili + time.Now().UnixNano()/3)

	return secret
}
