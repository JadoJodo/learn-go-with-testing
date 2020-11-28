package iteration

func Repeat(initial string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += initial
	}
	return repeated
}