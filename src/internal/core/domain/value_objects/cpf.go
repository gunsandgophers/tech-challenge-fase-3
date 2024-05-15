package valueobjects

type CPF struct {
	digits string
}

func NewCPF(digits string) *CPF {
	return &CPF{digits: digits}
}

func (c *CPF) Formated() string {
	return c.digits
}
