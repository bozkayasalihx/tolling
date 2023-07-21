package aggr

import (
	"errors"

	"github.com/bozkayasalihx/paid_road/types"
)

var DuplicateValue = errors.New("already in it")

type InMemStore struct {
	data map[int]float64
}

func NewInMemStore() *InMemStore {
	return &InMemStore{
		data: make(map[int]float64),
	}
}

func (m *InMemStore) Insert(data types.Distance) error {
	m.data[data.ID] += data.Distance
	return nil
}
