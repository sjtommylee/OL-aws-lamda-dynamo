package service

import "math/rand"

type Rand struct {
	random *rand.Rand
}

func (r *Rand) Get() *rand.Rand {
	return r.random
}
