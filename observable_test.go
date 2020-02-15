package rxgo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Observable_All(t *testing.T) {
	predicateAllInt := func(i interface{}) bool {
		switch i.(type) {
		case int:
			return true
		default:
			return false
		}
	}

	AssertSingle(context.Background(), t, testObservable(1, 2, 3).All(context.Background(), predicateAllInt),
		HasItem(true), HasNotRaisedError())
	AssertSingle(context.Background(), t, testObservable(1, "x", 3).All(context.Background(), predicateAllInt),
		HasItem(false), HasNotRaisedError())
}

func Test_Observable_AverageFloat32(t *testing.T) {
	AssertSingle(context.Background(), t, testObservable(float32(1), float32(2), float32(3)).AverageFloat32(context.Background()), HasItem(float32(2)))
	AssertSingle(context.Background(), t, testObservable(float32(1), float32(20)).AverageFloat32(context.Background()), HasItem(float32(10.5)))
	AssertSingle(context.Background(), t, Empty().AverageFloat32(context.Background()), HasItem(0))
	AssertSingle(context.Background(), t, testObservable("x").AverageFloat32(context.Background()), HasRaisedAnError())
}

func Test_Observable_AverageFloat64(t *testing.T) {
	AssertSingle(context.Background(), t, testObservable(float64(1), float64(2), float64(3)).AverageFloat64(context.Background()), HasItem(float64(2)))
	AssertSingle(context.Background(), t, testObservable(float64(1), float64(20)).AverageFloat64(context.Background()), HasItem(10.5))
	AssertSingle(context.Background(), t, Empty().AverageFloat64(context.Background()), HasItem(0))
	AssertSingle(context.Background(), t, testObservable("x").AverageFloat64(context.Background()), HasRaisedAnError())
}

func Test_Observable_AverageInt(t *testing.T) {
	AssertSingle(context.Background(), t, testObservable(1, 2, 3).AverageInt(context.Background()), HasItem(2))
	AssertSingle(context.Background(), t, testObservable(1, 20).AverageInt(context.Background()), HasItem(10))
	AssertSingle(context.Background(), t, Empty().AverageInt(context.Background()), HasItem(0))
	AssertSingle(context.Background(), t, testObservable(1.1, 2.2, 3.3).AverageInt(context.Background()), HasRaisedAnError())
}

func Test_Observable_AverageInt8(t *testing.T) {
	AssertSingle(context.Background(), t, testObservable(int8(1), int8(2), int8(3)).AverageInt8(context.Background()), HasItem(int8(2)))
	AssertSingle(context.Background(), t, testObservable(int8(1), int8(20)).AverageInt8(context.Background()), HasItem(int8(10)))
	AssertSingle(context.Background(), t, Empty().AverageInt8(context.Background()), HasItem(0))
	AssertSingle(context.Background(), t, testObservable(1.1, 2.2, 3.3).AverageInt8(context.Background()), HasRaisedAnError())
}

func Test_Observable_AverageInt16(t *testing.T) {
	AssertSingle(context.Background(), t, testObservable(int16(1), int16(2), int16(3)).AverageInt16(context.Background()), HasItem(int16(2)))
	AssertSingle(context.Background(), t, testObservable(int16(1), int16(20)).AverageInt16(context.Background()), HasItem(int16(10)))
	AssertSingle(context.Background(), t, Empty().AverageInt16(context.Background()), HasItem(0))
	AssertSingle(context.Background(), t, testObservable(1.1, 2.2, 3.3).AverageInt16(context.Background()), HasRaisedAnError())
}

func Test_Observable_AverageInt32(t *testing.T) {
	AssertSingle(context.Background(), t, testObservable(int32(1), int32(2), int32(3)).AverageInt32(context.Background()), HasItem(int32(2)))
	AssertSingle(context.Background(), t, testObservable(int32(1), int32(20)).AverageInt32(context.Background()), HasItem(int32(10)))
	AssertSingle(context.Background(), t, Empty().AverageInt32(context.Background()), HasItem(0))
	AssertSingle(context.Background(), t, testObservable(1.1, 2.2, 3.3).AverageInt32(context.Background()), HasRaisedAnError())
}

func Test_Observable_AverageInt64(t *testing.T) {
	AssertSingle(context.Background(), t, testObservable(int64(1), int64(2), int64(3)).AverageInt64(context.Background()), HasItem(int64(2)))
	AssertSingle(context.Background(), t, testObservable(int64(1), int64(20)).AverageInt64(context.Background()), HasItem(int64(10)))
	AssertSingle(context.Background(), t, Empty().AverageInt64(context.Background()), HasItem(0))
	AssertSingle(context.Background(), t, testObservable(1.1, 2.2, 3.3).AverageInt64(context.Background()), HasRaisedAnError())
}

func Test_Observable_BufferWithCount_CountAndSkipEqual(t *testing.T) {
	obs := testObservable(1, 2, 3, 4, 5, 6).BufferWithCount(context.Background(), 3, 3)
	AssertObservable(context.Background(), t, obs, HasItems([]interface{}{1, 2, 3}, []interface{}{4, 5, 6}))
}

func Test_Observable_BufferWithCount_CountAndSkipNotEqual(t *testing.T) {
	obs := testObservable(1, 2, 3, 4, 5, 6).BufferWithCount(context.Background(), 2, 3)
	AssertObservable(context.Background(), t, obs, HasItems([]interface{}{1, 2}, []interface{}{4, 5}))
}

func Test_Observable_BufferWithCount_IncompleteLastItem(t *testing.T) {
	obs := testObservable(1, 2, 3, 4).BufferWithCount(context.Background(), 2, 3)
	AssertObservable(context.Background(), t, obs, HasItems([]interface{}{1, 2}, []interface{}{4}))
}

func Test_Observable_BufferWithCount_Error(t *testing.T) {
	obs := testObservable(1, 2, 3, 4, errFoo).BufferWithCount(context.Background(), 3, 3)
	AssertObservable(context.Background(), t, obs, HasItems([]interface{}{1, 2, 3}, []interface{}{4}), HasRaisedError(errFoo))
}

//func Test_Observable_BufferWithInvalidInputs(t *testing.T) {
//	obs := Just(1, 2, 3, 4).BufferWithCount(0, 5)
//	AssertObservable(t, obs, HasRaisedAnError())
//
//	obs = Just(1, 2, 3, 4).BufferWithCount(5, 0)
//	AssertObservable(t, obs, HasRaisedAnError())
//}

func Test_Observable_Filter(t *testing.T) {
	obs := testObservable(1, 2, 3, 4).Filter(context.Background(),
		func(i interface{}) bool {
			return i.(int)%2 == 0
		})
	AssertObservable(context.Background(), t, obs, HasItems(2, 4), HasNotRaisedError())
}

func Test_Observable_ForEach(t *testing.T) {
	count := 0
	var gotErr error
	done := make(chan struct{})

	obs := testObservable(1, 2, 3, errFoo)
	obs.ForEach(context.Background(), func(i interface{}) {
		count += i.(int)
	}, func(err error) {
		gotErr = err
		done <- struct{}{}
	}, func() {})

	// We avoid using the assertion API on purpose
	<-done
	assert.Equal(t, 6, count)
	assert.Equal(t, errFoo, gotErr)
}

func Test_Observable_Map_One(t *testing.T) {
	obs := testObservable(1, 2, 3).Map(context.Background(), func(i interface{}) (interface{}, error) {
		return i.(int) + 1, nil
	})
	AssertObservable(context.Background(), t, obs, HasItems(2, 3, 4), HasNotRaisedError())
}

func Test_Observable_Map_Multiple(t *testing.T) {
	obs := testObservable(1, 2, 3).Map(context.Background(), func(i interface{}) (interface{}, error) {
		return i.(int) + 1, nil
	}).Map(context.Background(), func(i interface{}) (interface{}, error) {
		return i.(int) * 10, nil
	})
	AssertObservable(context.Background(), t, obs, HasItems(20, 30, 40), HasNotRaisedError())
}

func Test_Observable_Map_Error(t *testing.T) {
	obs := testObservable(1, 2, 3, errFoo).Map(context.Background(), func(i interface{}) (interface{}, error) {
		return i.(int) + 1, nil
	})
	AssertObservable(context.Background(), t, obs, HasItems(2, 3, 4), HasRaisedError(errFoo))
}

func Test_Observable_Map_Cancel(t *testing.T) {
	next := make(chan Item)

	ctx, cancel := context.WithCancel(context.Background())
	obs := FromChannel(next).Map(ctx, func(i interface{}) (interface{}, error) {
		return i.(int) + 1, nil
	})
	cancel()
	AssertObservable(context.Background(), t, obs, HasNoItems(), HasNotRaisedError())
}

func Test_Observable_Observe(t *testing.T) {
	got := make([]int, 0)
	ch := testObservable(1, 2, 3).Observe()
	for item := range ch {
		got = append(got, item.Value.(int))
	}
	assert.Equal(t, []int{1, 2, 3}, got)
}

func Test_Observable_SkipWhile(t *testing.T) {
	obs := testObservable(1, 2, 3, 4, 5).SkipWhile(context.Background(), func(i interface{}) bool {
		switch i := i.(type) {
		case int:
			return i != 3
		default:
			return true
		}
	})

	AssertObservable(context.Background(), t, obs, HasItems(3, 4, 5), HasNotRaisedError())
}
