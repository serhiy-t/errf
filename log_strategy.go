package errflow

type logStrategy int

const (
	logStrategyDefault = logStrategy(iota)
	logStrategyNever
	logStrategyIfSuppressed
	logStrategyAlways
)

func setLogStrategy(ef *Errflow, ls logStrategy) *Errflow {
	newEf := ef.Copy()
	if ef.logStrategy == logStrategyDefault {
		newEf.logStrategy = ls
	}
	return newEf
}

func LogStrategyNever(ef *Errflow) *Errflow {
	return setLogStrategy(ef, logStrategyNever)
}

func LogStrategyIfSuppressed(ef *Errflow) *Errflow {
	return setLogStrategy(ef, logStrategyIfSuppressed)
}

func LogStrategyAlways(ef *Errflow) *Errflow {
	return setLogStrategy(ef, logStrategyAlways)
}
