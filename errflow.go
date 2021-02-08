package errflow

type errflowThrow struct {
	errs []error
}

func ImplementCheck(recoverObj interface{}, err error) error {
	globalErrflowValidator.validate()
	var errs []error
	if recoverObj != nil {
		errflowThrowObj, ok := recoverObj.(errflowThrow)
		if ok {
			errs = errflowThrowObj.errs
		} else {
			panic(recoverObj)
		}
	}
	if err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		panic(errflowThrow{errs: errs})
	}
	return nil
}

func Check(err error) error {
	return ImplementCheck(recover(), err)
}

func C(err error) error {
	return ImplementCheck(recover(), err)
}

func Untyped(value interface{}, err error) interface{} {
	ImplementCheck(recover(), err)
	return value
}

func U(value interface{}, err error) interface{} {
	ImplementCheck(recover(), err)
	return value
}

func IgnoreReturnValue(value interface{}, err error) error {
	return ImplementCheck(recover(), err)
}

func I(value interface{}, err error) error {
	return ImplementCheck(recover(), err)
}
