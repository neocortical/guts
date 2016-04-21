package multierr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Push(t *testing.T) {
	var me MultiError

	me = me.Push(nil)
	assert.Nil(t, me)

	me = me.Push(errors.New("herp"))
	assert.NotNil(t, me)
	assert.Equal(t, 1, len(me))
	assert.Equal(t, "herp", me.Error())

	me = me.Push(errors.New("derp"))
	assert.NotNil(t, me)
	assert.Equal(t, 2, len(me))
	assert.Equal(t, `2 errors: ["herp", "derp"]`, me.Error())

	var me2 MultiError
	me2 = me2.Push(me)
	assert.Equal(t, 1, len(me2))
	assert.Equal(t, `2 errors: ["herp", "derp"]`, me2.Error())

	me2 = me2.Push(me)
	assert.Equal(t, 2, len(me2))
	assert.Equal(t, `2 errors: ["2 errors: ["herp", "derp"]", "2 errors: ["herp", "derp"]"]`, me2.Error())

	// instantiation edge case
	me = MultiError{}
	assert.Equal(t, "", me.Error())
}
