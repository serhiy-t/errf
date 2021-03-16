package errf

// Opts combines mulitple options into a single.
func Opts(options ...ErrflowOption) ErrflowOption {
	return func(ef *Errflow) *Errflow {
		for _, opt := range options {
			ef = opt(ef)
		}
		return ef
	}
}

// OptsFrom creates ErrflowOption which applies all options
// from provided Errflow.
func OptsFrom(ef *Errflow) ErrflowOption {
	return Opts(ef.Opts()...)
}
