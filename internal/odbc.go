// +build all,!no_odbc odbc,!no_odbc

package internal

import (
	// odbc driver
	_ "github.com/knq/usql/drivers/odbc"
)
