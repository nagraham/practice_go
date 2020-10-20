package leetcode

import "errors"

/**
 * NOTE:
 * I'm going to leave this as a monument to complexity.
 *
 * This was an attempt to make some code more modular, by extracting out the functionality to
 * track a maximum integer. But honestly, because the ReduceFunction requires type checking, you
 * end up with more code.
 *
 * Additionally, there's already an orthogonal and easy function to help with maximum integers: math.Max
 * (although you have to convert to/from float64(), that isn't hard).
 *
 * This was a situation where just because I could decompose the problem into a different module didn't
 * mean this was the *correct* module, or the simplest available. You need to stop, think, and question
 * the design choice. Try to list out different alternatives for solving a problem with a module. Pick
 * the one with the right balance of tradeoffs.
 *
 * That said, this "Reducer" may be useful in a different problem.
 */

// this is not a leetcode solution, but more of a general solution
// not ready to put into a module until I have a few more like it.

// A function that takes two objects and reduces them into a value.
type ReduceFunction = func(reduction, newValue interface{}) (interface{}, error)

// A Reducer allows you to execute a reduce function. Ideally, the implementation should
// cache the reduction, so this can be executed statefully.
//
// This could be used in a case where reduction functionality (e.g. get Sum, find Max, find Min)
// is needed, but cannot be executed in its own iterative loop. Perhaps it is distributed over
// one or more other orthogonal iterative loops or recursive chains, even on different go threads.
type Reducer interface {
	Reduce(x interface{}) error
	Get() interface{}
}

// Implementation struct for the Reducer interface
type reducer struct {
	reduceFunction ReduceFunction
	reduction      interface{}
}

func NewReducer(compareFunc ReduceFunction) Reducer {
	return &reducer{
		reduction:      nil,
		reduceFunction: compareFunc,
	}
}

func (cs *reducer) Reduce(x interface{}) error {
	if x == nil {
		return errors.New("can't compare nil")
	} else if cs.reduction == nil {
		cs.reduction = x
		return nil
	}

	newBest, err := cs.reduceFunction(cs.reduction, x)
	if err != nil {
		return err
	}
	cs.reduction = newBest
	return nil
}

func (cs *reducer) Get() interface{} {
	return cs.reduction
}