#!/usr/bin/perl

use strict;
use warnings;

use File::Basename;
use lib File::Basename::dirname (__FILE__);

use Test::More;
use FTests;

sub id { return $_[0]; }

FTests::run([
	{
		"f"        => sub { return $_[0]; },
		"fn"       => '(\\x = x)',
		"args"     => [1],
		"expected" => 1,
		"descr"    => "(\\x = x) 1",
	},
	{
		"f"        => \&id,
		"args"     => [{foo=>"bar"}],
		"expected" => {foo=>"bar"},
		"descr"    => "id({foo=>'bar'}) = {foo=>'bar'}",
	},
	{
		"f"        => \&FTests::run,
		"args"     => [[{
			"f"        => \&id,
			"fn"       => "FTests::run([id])",
			"args"     => [{foo=>"baz"}],
			"expected" => {foo=>"baz"},
			"descr"    => "id({foo=>'baz'}) = {foo=>'baz'}",
		}]],
		"expected" => 1,
		"descr"    => "going meta",
	},
]);

Test::More::done_testing();
