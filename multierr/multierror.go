package multierr

import (
	"bytes"
	"fmt"
)

// MultiError accumulates errors into a single error. It is useful when a caller wants to
// know everything that went wrong even if some error conditions shouldn't abort an operation.
// NOTE: This type is not threadsafe.
type MultiError []error

// Push an error into a MultiError type. If the err parameter is nil, then no action is taken.
func (me MultiError) Push(err error) MultiError {
	if err == nil {
		return me
	}

	return append(me, err)
}

// Error implements the error interface.
func (me MultiError) Error() string {
	return MultiErrorDefaultFunc(me)
}

// MultiErrorDefaultFunc is the default behavior of MultiError's Error() method. It is exposed
// by the package so it can be overridden if desired.
var MultiErrorDefaultFunc = func(me MultiError) string {
	if len(me) == 0 {
		return ""
	}
	if len(me) == 1 {
		return me[0].Error()
	}
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%d errors: [", len(me))
	for i, err := range me {
		if i != 0 {
			fmt.Fprint(buf, ", ")
		}
		fmt.Fprintf(buf, `"%v"`, err)
	}
	fmt.Fprint(buf, "]")
	return buf.String()
}
