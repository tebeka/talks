package main

// START_SEQ OMIT
type Seq[V any] func(yield func(V) bool)

type Seq2[K, V any] func(yield func(K, V) bool)

// END_SEQ OMIT

// START_PULL OMIT
func Pull[V any](seq Seq[V]) (next func() (V, bool), stop func())

func Pull2[K, V any](seq Seq2[K, V]) (next func() (K, V, bool), stop func())

// END_PULL OMIT
