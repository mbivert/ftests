# Note
I'm publishing this mainly because I "need" it as a Go module
(I'm starting to need to copy ``ftests.go`` multiple times within
the same project, for different packages).

That's to say, copying the file is probably what you'd want
instead of adding a dependency. I'm including implementations
for other languages for completeness.

# Examples

Some places where it's being used:

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

# Introduction
This directory contains, in their respective sub-directories,
per-programming language implementation examples of a rustic
function-based test framework.

Loosely, the main data-structure, a `test`, is mainly:

  - a function (pointer to function);
  - some input: function arguments ([tuple][wp-en-tuple]);
  - some (expected) output: return values ([tuple][wp-en-tuple]);
  - a description (string);

Such a test is run by calling the related function with
the given arguments, and comparing its output to the
expected values.

Most modern languages have [tuples][wp-en-tuple], and
allow, at least in most cases, for some kind of deep-comparison,
which are the two main requirements for an implementation.

Because of the regularity of the test format, it
should be easy to move away from this pattern if need be. Or
the other way around: it's a great way to progressively add
tests to an existing codebase with almost no overhead.

For more, see this [blog post][tales-ftest].

Implementations are provided as references only. The main
takeaway is the pattern; implementation details can vary from
project to project, or language to language.

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
