package cli

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCli(t *testing.T) {
	tests := []struct {
		name         string
		errStream    string
		cmdLineArg   string
		expectedOpts *Options
		expectedMsg  string
	}{
		{
			name:         "no opts",
			cmdLineArg:   "",
			expectedOpts: &Options{false, false, false},
		},
		{
			name:         "show help",
			cmdLineArg:   "-h",
			expectedOpts: &Options{false, false, false},
			expectedMsg:  "gjo -",
		},
		{
			name:         "show version",
			cmdLineArg:   "-v",
			expectedOpts: &Options{true, false, false},
		},
		{
			name:         "pretty print & array value",
			cmdLineArg:   "-p -a",
			expectedOpts: &Options{false, true, true},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ins := &bytes.Reader{}
			outs := &bytes.Buffer{}
			errs := &bytes.Buffer{}
			args := strings.Split(test.cmdLineArg, " ")

			c, err := newCli(ins, outs, errs, args)
			if test.expectedMsg != "" {
				assert.NotNil(t, err)
				assert.True(t, strings.Contains(errs.String(), test.expectedMsg), fmt.Sprintf("'%s' should contain '%s'", errs.String(), test.expectedMsg))
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, c)
				assert.Empty(t, outs.String())
				assert.Empty(t, errs.String())
				assert.Equal(t, test.expectedOpts, c.opts)
			}
		})
	}
}
