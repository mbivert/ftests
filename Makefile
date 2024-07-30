.PHONY: all
all: tests js-tests perl-tests nix-tests

.PHONY: help
help:
	@echo "all            : run all tests"
	@echo "tests          : run Go tests"
	@echo "js-tests       : run JS tests"
	@echo "perl-tests     : run Perl tests"
	@echo "nix-tests      : run Nix tests"

.PHONY: tests
tests:
	@echo Running Go tests...
	@go test -v .

.PHONY: js-tests
js-tests: js/full.js
	@echo Running JS tests...
	@node js/full.js

js/full.js: js/tests.js js/tests_test.js
	@echo Creating $@...
	@(cat $^; echo "Tests.run(TestsTests.tests);") > $@

.PHONY: perl-tests
perl-tests: perl/ftests.pl
	@echo Running Perl tests...
	@perl $^

.PHONY: nix-tests
nix-tests: nix/ftests_test.nix
	@echo Running Nix tests...
	@nix-instantiate --eval $^
