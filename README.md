# Overview
This package provides a simple and systematic way to test
functions: we merely check that on given inputs, functions indeed
compute the expected outputs; tests are still managed via the
standard [``testing``][godoc-testing] package.

Tests can then be wrapped in a simple ``struct{}``:

    type Test struct {
    	Name string // Test name
    	Fun any // Function to test
    	Args []any // Inputs
    	Expected []any // Expected outputs
    }

Arrays of tests may then be tested using ``ftests.Run()``:

    import (
    	"testing"
    	"github.com/mbivert/ftests"
    )

    func TestSomething(t *testing.T) {
    	ftests.Run(t, []ftests.Test{
    		{
    			"(x -> x)(42) = 42",
    			func(x any) any { return x },
    			[]interface{}{42},
    			[]interface{}{42},
    		},
    		{
        		...
    		},
    		...
    	})
    }

"Weird" cases (e.g. environment alteration) can be managed
by wrapping the code with a few extra functions. See
this [blog post][tales-ftest] for some examples.

**<u>Note:</u>** Because of the regularity of the tests format,
it's easy to move towards or away from this pattern, to programatically
alter the tests, to run the same tests in different settings, etc.
In particular, it's a great way to progressively add
tests to an existing codebase with almost no overhead.

# Other languages (Nix, JavaScript, Perl, etc.)
Implementations for other languages are also provided. The
main requirements are rather standard nowadays:

  - [tuples][wp-en-tuple];
  - deep-comparison mechanism.

# Usage examples

  - ``go(1)``:
    - [github.com/mbivert/auth][github.com/mbivert/auth];
    - [github.com/mbivert/mkshort][github.com/mbivert/mkshort];
    - [github.com/mbivert/log-tools][github.com/mbivert/log-tools];
    - [github.com/mbivert/ovh-tools][github.com/mbivert/ovh-tools];
  - ``perl(1)``:
    - [github.com/mbivert/log-tools/tree/master/old/perl/tests][github.com/mbivert/log-tools/tree/master/old/perl/tests]
  - ``sh(1)`` / ``cc(1)``:
    - [github.com/mbivert/qsplit][github.com/mbivert/qsplit];
  - JavaScript:
    - [github.com/mbivert/zm/tree/master/tests][github.com/mbivert/zm/tree/master/tests];
  - ``nix(1)``:
    - [github.com/mbivert/nix-series-code][github.com/mbivert/nix-series-code].

[tales-ftest]: https://tales.mbivert.com/on-function-based-test-framework/

[wp-en-tuple]: https://en.wikipedia.org/wiki/Tuple

[github.com/mbivert/auth]:    https://github.com/mbivert/auth
[github.com/mbivert/mkshort]: https://github.com/mbivert/mkshort
[github.com/mbivert/qsplit]: https://github.com/mbivert/qsplit
[github.com/mbivert/log-tools]: https://github.com/mbivert/log-tools
[github.com/mbivert/ovh-tools]: https://github.com/mbivert/ovh-tools

[github.com/mbivert/log-tools/tree/master/old/perl/tests]: https://github.com/mbivert/log-tools/tree/master/old/perl/tests

[github.com/mbivert/zm/tree/master/tests]: https://github.com/mbivert/zm/tree/master/tests

[github.com/mbivert/nix-series-code]: https://github.com/mbivert/nix-series-code

[godoc-testing]: https://pkg.go.dev/testing
