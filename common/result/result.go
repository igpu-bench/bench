package result

type Result struct {
	errors []error
}

func New() *Result {
	var res *Result

	return res
}
