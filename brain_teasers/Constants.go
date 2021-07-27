package brain_teasers

// Max int value taken from below:
// https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const MAX_UINT_16 = ^uint16(0)

const MAX_INT_16 = int16(MAX_UINT_16 >> 1)
