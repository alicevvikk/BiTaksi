package errs

type errs struct {
	Message string
	Code	int64
}

func (er *errs) New(m string, c, int64) *errs{
	return &errs{
		Message:	m,
		Code:		c,
	}
}
