package util

import "sync/atomic"

type ModelId uint32

var modelIdCounter atomic.Uint32

func NewModelId() ModelId {
	return ModelId(modelIdCounter.Add(1))
}
