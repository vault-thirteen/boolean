// +build test

package boolean

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_StringToBoolean(t *testing.T) {

	var aTest *tester.Test
	var err error
	var result bool

	aTest = tester.New(t)

	// Test #1.
	result, err = StringToBoolean("TRuE")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, true)

	// Test #2.
	result, err = StringToBoolean("FalsE")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, false)

	// Test #3.
	result, err = StringToBoolean("x")
	aTest.MustBeAnError(err)

	// Test #4.
	result, err = StringToBoolean("123456789")
	aTest.MustBeAnError(err)
}
