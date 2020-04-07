package accumulate

// Words is the main type for the Accumulate function
type Words []string

// Accumulate transforms the strings with the associated function
func Accumulate(input Words, function func(string) string) (output Words) {
	for _, word := range input {
		output = append(output, function(word))
	}
	return
}
