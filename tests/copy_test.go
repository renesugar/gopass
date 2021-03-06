package tests

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping test on windows.")
	}
	ts := newTester(t)
	defer ts.teardown()

	_, err := ts.run("copy")
	assert.Error(t, err)

	ts.initStore()

	out, err := ts.run("copy")
	assert.Error(t, err)
	assert.Equal(t, "\nError: Usage: "+filepath.Base(ts.Binary)+" cp <FROM> <TO>\n", out)

	out, err = ts.run("copy foo")
	assert.Error(t, err)
	assert.Equal(t, "\nError: Usage: "+filepath.Base(ts.Binary)+" cp <FROM> <TO>\n", out)

	out, err = ts.run("copy foo bar")
	assert.Error(t, err)
	assert.Equal(t, "\nError: foo does not exist\n", out)

	ts.initSecrets("")

	// recursive copy
	_, err = ts.run("copy foo/ bar")
	require.NoError(t, err)

	out, err = ts.run("copy foo/bar foo/baz")
	require.NoError(t, err)
	assert.Equal(t, "", out)

	orig, err := ts.run("show -f foo/bar")
	assert.NoError(t, err)

	copy, err := ts.run("show -f foo/baz")
	assert.NoError(t, err)

	assert.Equal(t, orig, copy)
}
