package brain_teasers

type quickIntTracker struct {
	tally   int16
	numbers []bool
}

func (qit *quickIntTracker) Run(r Randomizer) {
	n := int16(0)

	for qit.tally < MAX_INT_16 {
		n = int16(r.Next())

		if !qit.numbers[n] {
			qit.numbers[n] = true

			qit.tally += 1
		}
	}
}

func NewIntTrackerQuick() IntTracker {
	rtn := new(quickIntTracker)

	rtn.tally = 0

	rtn.numbers = make([]bool, MAX_INT_16)

	return (rtn)
}
