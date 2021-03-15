package errf

type logStrategy int

const (
	logStrategyDefault = logStrategy(iota)
	logStrategyNever
	logStrategyIfSuppressed
	logStrategyAlways
)

func setLogStrategy(ef *Errflow, ls logStrategy) *Errflow {
	newEf := ef.copy()
	if ef.logStrategy == logStrategyDefault {
		newEf.logStrategy = ls
	}
	return newEf
}

// LogStrategyNever configures Errflow instance to never log errors.
// This is default behavior.
func LogStrategyNever(ef *Errflow) *Errflow {
	return setLogStrategy(ef, logStrategyNever)
}

// LogStrategyIfSuppressed configures Errflow instance to log errors only if
// they are not present in resulting error value.
func LogStrategyIfSuppressed(ef *Errflow) *Errflow {
	return setLogStrategy(ef, logStrategyIfSuppressed)
}

// LogStrategyAlways configures Errflow instance to always log errors.
func LogStrategyAlways(ef *Errflow) *Errflow {
	return setLogStrategy(ef, logStrategyAlways)
}
