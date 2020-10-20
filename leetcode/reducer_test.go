package leetcode

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createMaxIntFunc() ReduceFunction {
	return func (a, b interface{}) (interface{}, error) {
		intA, aOk := a.(int)
		intB, bOk := b.(int)
		if !aOk || !bOk {
			return nil, errors.New("expected an int")
		}

		if intA > intB {
			return intA, nil
		} else {
			return intB, nil
		}
	}
}

func TestReducer_MaxIntFunc(t *testing.T) {
	reducer := NewReducer(createMaxIntFunc())

	err := reducer.Reduce(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, reducer.Get())

	err = reducer.Reduce(2)
	assert.NoError(t, err)
	assert.Equal(t, 2, reducer.Get())

	err = reducer.Reduce(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, reducer.Get())

	err = reducer.Reduce(42)
	assert.NoError(t, err)
	assert.Equal(t, 42, reducer.Get())
}

func TestReducer_WithNilArgument(t *testing.T) {
	reducer := NewReducer(createMaxIntFunc())

	err := reducer.Reduce(nil)
	assert.EqualError(t, err, "can't compare nil")
}

func TestReducer_WhenGivenReducerFunctionReturnsError(t *testing.T) {
	reducer := NewReducer(createMaxIntFunc())

	err := reducer.Reduce(1)
	err = reducer.Reduce("2")
	assert.EqualError(t, err, "expected an int")
}