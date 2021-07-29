package collections

import "testing"

func testQueue(t *testing.T, items []interface{}) {
	subject := new(SliceQueue)

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
				v)
		}

		err := subject.Dequeue()

		if err != nil {
			t.Errorf("Dequeue failed: %v", err)
		}
	}
}

func Test_SliceQueue_ints(t *testing.T) {
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

		testQueue(t, arr)
	}
}

func Test_SliceQueue_strings(t *testing.T) {
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

		testQueue(t, arr)
	}
}
