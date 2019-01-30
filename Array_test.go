package Array

import "testing"

func TestSlice(t *testing.T) {
	arr := Array{1, 2, 3}

	expected := Array{2}
	actual := arr.Slice(1, 2)

	if !isEq(actual, expected) {
		throwErr(t, "Slice works not correct", expected, actual)
	}
}

func TestSplice(t *testing.T) {
	arr := Array{1, 2, 3, 4, 5}

	expected := Array{3, 4, 5}
	actual := arr.Splice(2, 3)

	if !isEq(actual, expected) {
		throwErr(t, "Splice works not correct", expected, actual)
	}

	expected = Array{1, 2}
	actual = arr

	if !isEq(actual, expected) {
		throwErr(t, "Splice works not correct (origin slice)", expected, actual)
	}
}

func TestEvery(t *testing.T) {
	expected := true
	actual := Array{1, 3, 5}.Every(isEven);

	if actual != expected {
		throwErr(t, "Every works not correct", expected, actual)
	}
}

func TestSome(t *testing.T) {
	expected := true
	actual := Array{1, 4, 6}.Some(isEven)

	if actual != expected {
		throwErr(t, "Some works not correct", expected, actual)
	}
}

func TestMap(t *testing.T) {
	expected := Array{2, 3, 4}
	actual := Array{1, 2, 3}.Map(func (item Any, _ int, _ Array) Any {
		newItem := toInt(item)

		return newItem + 1
	})

	if !isEq(actual, expected) {
		throwErr(t, "Map works not correct", expected, actual)
	}
}

func TestPush(t *testing.T) {
	expected := Array{1, 2, 3, 4}
	arr := Array{1, 2}
	arr.Push(3, 4)

	actual := arr

	if !isEq(actual, expected) {
		throwErr(t, "Push works not correct", expected, actual)
	}
}

func TestPop(t *testing.T) {
	expected := Array{1, 2, 3}
	arr := Array{1, 2, 3, 4}
	arr.Pop()

	actual := arr

	if !isEq(actual, expected) {
		throwErr(t, "Pop works not correct", expected, actual)
	}
}

func TestShift(t *testing.T) {
	expected := Array{2, 3}
	arr := Array{1, 2, 3}
	arr.Shift()

	actual := arr

	if !isEq(actual, expected) {
		throwErr(t, "Shift works not correct", expected, actual)
	}
}

func TestReduce(t *testing.T) {
	expected := 6
	actual := Array{1, 2, 3}.Reduce(func(prev Any, cur Any, _ int, _ Array) Any {
		newPrev := toInt(prev)
		newCur := toInt(cur)

		return newPrev + newCur
	}, 0)

	if actual != expected {
		throwErr(t, "Reduce works not correct", expected, actual)
	}
}

func TestReverse(t *testing.T) {
	expected := Array{3, 2, 1}
	actual := Array{1, 2, 3}.Reverse()

	if !isEq(actual, expected) {
		throwErr(t, "Reverse works not correct", expected, actual)
	}
}

func TestFilter(t *testing.T) {
	expected := Array{1, 3}
	actual := Array{1, 2, 3}.Filter(isEven)

	if !isEq(actual, expected) {
		throwErr(t, "Filter works not correct", expected, actual)
	}
}

func TestForEach(t *testing.T) {
	expected := 6
	actual := 0

	Array{1, 2, 3}.ForEach(func(item Any, _ int, _ Array) {
		actual += toInt(item)
	})

	if actual != expected {
		throwErr(t, "ForEach works not correct", expected, actual)
	}
}
