package brain_teasers

import "testing"

func Benchmark_IntQuickTrackerQuick_Run(b *testing.B) {
	r := NewRandomizer()

	subject := NewIntTrackerQuick()

	subject.Run(r)
}
