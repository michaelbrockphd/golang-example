package collections

import "testing"

func testNodeQueue(t *testing.T, items []interface{}) {
	subject := new(NodeQueue)

	for _, v := range items {
		subject.Enqueue(v)

		if v != subject.Tail() {
			t.Errorf(
				"Enqueue failed: expected %v but got %v",
				v,
				subject.Tail())
		}
	}

	var h interface{}

	for _, v := range items {
		h = subject.Head()

		if h != v {
			t.Errorf(
				"Head failed: expected %v but got %v",
				v,
				h)
		}

		err := subject.Dequeue()

		if err != nil {
			t.Errorf("Dequeue failed: %v", err)
		}
	}
}

func Test_NodeQueue_ints(t *testing.T) {
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

		testNodeQueue(t, arr)
	}
}

func Test_NodeQueue_strings(t *testing.T) {
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

		testNodeQueue(t, arr)
	}
}

func benchmarkNodeQueueInts(min int, max int, b *testing.B) {
	s := new(NodeQueue)

	v := int64(min)

	m := int64(max)

	for v <= m {
		s.Enqueue(v)

		v += 1
	}

	for 0 < s.length {
		s.Dequeue()
	}
}

func Benchmark_NodeQueue_10(b *testing.B) {
	benchmarkNodeQueueInts(0, 10, b)
}

func Benchmark_NodeQueue_100(b *testing.B) {
	benchmarkNodeQueueInts(0, 100, b)
}

func Benchmark_NodeQueue_1000(b *testing.B) {
	benchmarkNodeQueueInts(0, 1000, b)
}

func Benchmark_NodeQueue_4G(b *testing.B) {
	benchmarkNodeQueueInts(0, 1024*4, b)
}

func Benchmark_NodeQueue_1Gsq(b *testing.B) {
	benchmarkNodeQueueInts(0, 1024*1024, b)
}
