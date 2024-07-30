package ftests

// I'm being light here; code is simple and should have
// been battletested long enough anyway.

import (
	"testing"
)

func TestRun(t *testing.T) {
	Run(t, []Test{
		{
			"(x -> x)(42) = 42",
			func(x any) any { return x },
			[]any{42},
			[]any{42},
		},
		{
			"Meta (x -> x)(42) = 42 (expect success)",
			Run1,
			[]any{
				t,
				func(x any) any { return x },
				[]any{42},
				[]any{42},
			},
			nil,
		},
		// A bit annoying to test the failure case,
		// because we're calling t.Fatalf() on purpose.
		//
		// Anyway.
	})
}
