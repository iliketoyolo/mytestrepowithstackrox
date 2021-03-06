package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	t.Parallel()
	_, pass := valuesAreSet([]string{"hi"}, "", "", "")
	require.True(t, pass)

	_, pass = valuesAreSet([]string{}, "", "", "")
	require.False(t, pass)
}

func TestUnset(t *testing.T) {
	t.Parallel()
	_, pass := unset([]string{"hi"}, "", "", "")
	require.False(t, pass)

	_, pass = unset([]string{}, "", "", "")
	require.True(t, pass)
}

func TestMatchesFunctions(t *testing.T) {
	t.Parallel()

	var cases = []struct {
		name       string
		values     []string
		target     string
		defaultVal string
		pass       bool
	}{
		{
			name:   "No values",
			target: "random",
			pass:   false,
		},
		{
			name:   "Does match",
			target: "random",
			values: []string{
				"blah",
				"random",
			},
			pass: true,
		},
		{
			name:       "Matches default",
			target:     "random",
			defaultVal: "random",
			pass:       true,
		},
		{
			name:       "Does not match default",
			target:     "random",
			defaultVal: "notrandom",
			pass:       false,
		},
	}

	for _, c := range cases {
		t.Run(c.name+"-pass", func(t *testing.T) {
			_, pass := matches(c.values, "key", c.target, c.defaultVal)
			assert.Equal(t, c.pass, pass)
		})
		t.Run(c.name+"-notmatches", func(t *testing.T) {
			_, pass := notMatches(c.values, "key", c.target, c.defaultVal)
			assert.Equal(t, c.pass, !pass)
		})
	}
}

func TestContainsFunctions(t *testing.T) {
	t.Parallel()

	var cases = []struct {
		name       string
		values     []string
		target     string
		defaultVal string
		pass       bool
	}{
		{
			name:   "No values",
			target: "random",
			pass:   false,
		},
		{
			name:   "Does contain",
			target: "ran",
			values: []string{
				"blah",
				"random",
			},
			pass: true,
		},
		{
			name:       "Default contains target",
			target:     "ran",
			defaultVal: "random",
			pass:       true,
		},
		{
			name:       "Does not contain default",
			target:     "nope",
			defaultVal: "notrandom",
			pass:       false,
		},
	}

	for _, c := range cases {
		t.Run(c.name+"-contains", func(t *testing.T) {
			_, pass := contains(c.values, "key", c.target, c.defaultVal)
			assert.Equal(t, c.pass, pass)
		})
		t.Run(c.name+"-notcontains", func(t *testing.T) {
			_, pass := notContains(c.values, "key", c.target, c.defaultVal)
			assert.Equal(t, c.pass, !pass)
		})
	}
}

func TestOnlyContainsFunction(t *testing.T) {
	t.Parallel()

	var cases = []struct {
		name       string
		values     []string
		target     string
		defaultVal string
		pass       bool
	}{
		{
			name:   "No values",
			target: "random1,and,random2",
			pass:   false,
		},
		{
			name:   "Only contains",
			target: "random1,and,random2",
			values: []string{
				"random1",
				"random2",
			},
			pass: true,
		},
		{
			name:   "Contain alien",
			target: "random1,and,random2",
			values: []string{
				"random1",
				"notrandom2",
			},
			pass: false,
		},
		{
			name:       "Default contains target",
			target:     "random1",
			defaultVal: "random1",
			pass:       true,
		},
		{
			name:       "Does not contain default",
			target:     "random1",
			defaultVal: "notrandom1",
			pass:       false,
		},
	}

	for _, c := range cases {
		t.Run(c.name+"-onlycontains", func(t *testing.T) {
			_, pass := onlyContains(c.values, "key", c.target, c.defaultVal)
			assert.Equal(t, c.pass, pass)
		})
	}
}
