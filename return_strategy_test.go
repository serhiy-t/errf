package errf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getReturnStrategyImpl_unknown(t *testing.T) {
	assert.PanicsWithError(t, "unknown errflow return strategy: 1", func() {
		getReturnStrategyImpl(returnStrategyUnknown)
	})
}

func Test_returnStrategyCombinedImpl_no_errors(t *testing.T) {
	supp1, supp2, err := returnStrategyCombinedImpl(nil, nil)
	assert.NoError(t, err)
	assert.False(t, supp1)
	assert.False(t, supp2)
}

func Test_returnStrategyCombinedImpl_single_error(t *testing.T) {
	originalErr := fmt.Errorf("error 1")
	supp1, supp2, err := returnStrategyCombinedImpl(originalErr, nil)
	assert.Same(t, originalErr, err)
	assert.False(t, supp1)
	assert.False(t, supp2)
}

func Test_returnStrategyCombinedImpl_multiple_errors(t *testing.T) {
	supp1, supp2, err := returnStrategyCombinedImpl(fmt.Errorf("err1"), fmt.Errorf("err2"))
	assert.EqualError(t, err, "combined error {err1; err2}")
	assert.False(t, supp1)
	assert.False(t, supp2)
}

func Test_GetCombinedErrors(t *testing.T) {
	cerr := CombinedError{errs: []error{fmt.Errorf("err1"), fmt.Errorf("err2")}}
	werr := fmt.Errorf("wrapped: %w", cerr)
	assert.EqualError(t, werr, "wrapped: combined error {err1; err2}")
	assert.Equal(t,
		[]error{fmt.Errorf("err1"), fmt.Errorf("err2")},
		GetCombinedErrors(werr))
}
