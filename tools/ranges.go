package tools

type intRange struct {
	start int
	end   int
}

func RangesOverlap(a, b intRange) bool {
	return a.start <= b.end && b.start <= a.end
}

// this will fail if a & b do not overlap, so don't do that
func MergeRanges(a, b intRange) intRange {
	return intRange{MinInt(a.start, b.start), MaxInt(a.end, b.end)}
}
