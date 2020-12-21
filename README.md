# golang-testing-util-functions
Testing functions I use in my GO testing files

Usage:
Add repo to your '_test.go' file(s):

	import (
		"testing"

		tt "github.com/john2exonets/golang-testing-util-functions"
	)

Then call the functions as needed.

**tt.NotErr(t, err)**		-- Make sure there was no error returned.<br>
**tt.IsErr(t, err)**		-- Make sure an error was returned.<br>
**tt.Assert(t, a, b)**		-- Is 'a' equilivliant to 'b'. Works with any variable Type.<br>
**tt.AssertNot(t, a, b)**	-- Check to make sure 'a' is not equilivliant to 'b'. Works with any variable Type.<br>
**tt.Contains(t, a, b)**	-- Is 'a' a member of 'b'? Works with Strings and Slices.<br>
