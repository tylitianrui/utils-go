package expr

type Exp struct {
	*Interpreter
}

func New(expr string) (*Exp, error) {
	inptr, err := newInterpreter(expr, nil)

	return &Exp{inptr}, err
}
