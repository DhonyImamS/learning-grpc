package assertions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
)

// asserter is used to be able to retrieve the error reported by the called assertion
type asserter struct {
	err error
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// Errorf is used by the called assertion to report an error
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func GodogExpectedAndActual(customGodogExpect expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter
	customGodogExpect(&t, expected, actual, msgAndArgs...)
	return t.err
}