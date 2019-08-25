////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

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
