// boolean.go.

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
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

package boolean

import (
	"fmt"
	"strings"
)

const (
	TrueStringLc  = "true"
	FalseStringLc = "false"
)

const ErrfBadBoolean = "Bad boolean value: '%v'."

func StringToBoolean(
	str string,
) (bool, error) {

	var err error

	if len(str) < 4 {
		err = fmt.Errorf(ErrfBadBoolean, str)
		return false, err
	}

	str = strings.ToLower(str)

	switch str {

	case TrueStringLc:
		return true, nil

	case FalseStringLc:
		return false, nil

	default:
		err = fmt.Errorf(ErrfBadBoolean, str)
		return false, err
	}
}
