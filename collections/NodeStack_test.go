package collections

import "testing"

func testNodeStackPushPop(t *testing.T, items []interface{}) {
	subject := new(NodeStack)

	for _, v := range items {
		subject.Push(v)

		if v != subject.Top() {
			t.Errorf(
				"Push failed: expected %v but got %v",
				v,
				subject.Top())
		}
	}

	var v interface{}

	i := len(items)

	for 0 < i {
		i -= 1

		v = subject.Top()

		if v != items[i] {
			t.Errorf(
				"Top failed: expected %v but got %v",
				items[i],
				v)
		}

		err := subject.Pop()

		if err != nil {
			t.Errorf("Pop failed: %v", err)
		}
	}
}

func Test_NodeStack_ints(t *testing.T) {
	testCases := []struct {
		slice []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, testCase := range testCases {
		arr := make([]interface{}, len(testCase.slice))

		for i, v := range testCase.slice {
			arr[i] = v
		}

		testNodeStackPushPop(t, arr)
	}
}

func Test_NodeStack_strings(t *testing.T) {
	testCases := []struct {
		slice []string
	}{
		{
			[]string{"Cheese", "Grommit"},
		},
	}

	for _, testCase := range testCases {
		arr := make([]interface{}, len(testCase.slice))

		for i, v := range testCase.slice {
			arr[i] = v
		}

		testNodeStackPushPop(t, arr)
	}
}

func benchmarkNodeStackInts(min int, max int, b *testing.B) {
	s := new(NodeStack)

	v := int64(min)

	m := int64(max)

	for v <= m {
		s.Push(v)

		v += 1
	}

	for 0 < s.length {
		s.Pop()
	}
}

func Benchmark_NodeStack_10(b *testing.B) {
	benchmarkNodeStackInts(0, 10, b)
}

func Benchmark_NodeStack_100(b *testing.B) {
	benchmarkNodeStackInts(0, 100, b)
}

func Benchmark_NodeStack_1000(b *testing.B) {
	benchmarkNodeStackInts(0, 1000, b)
}

func Benchmark_NodeStack_4G(b *testing.B) {
	benchmarkNodeStackInts(0, 1024*4, b)
}

func Benchmark_NodeStack_1Gsq(b *testing.B) {
	benchmarkNodeStackInts(0, 1024*1024, b)
}
