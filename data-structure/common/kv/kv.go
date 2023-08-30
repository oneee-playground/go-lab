package kv

type KV[TKey, TValue any] struct {
	Key   TKey
	Value TValue
}
