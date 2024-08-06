package httpstatcheck_test

import (
	"testing"

	"github.com/minizilla/httpstatcheck"
	"github.com/minizilla/testr"
)

func TestChecker(t *testing.T) {
	t.Run("invalid rule", testChecker(false, 200, "2000"))
	t.Run("invalid wildcard", testChecker(false, 200, "2NN"))
	t.Run("invalid search", testChecker(false, 2000, "2XX"))
	t.Run("no rules", testChecker(false, 200))

	t.Run("match: all wildcard", testChecker(true, 200, "XXX"))
	t.Run("match: no wildcard", testChecker(true, 200, "200"))
	t.Run("match: with wildcard", testChecker(true, 200, "2XX"))
	t.Run("match: lowercase wildcard", testChecker(true, 200, "2xx"))
	t.Run("match: multiple rules", testChecker(true, 200, "200", "400"))
	t.Run("match: multiple rules with wildcard", testChecker(true, 200, "2XX", "400"))
	t.Run("match: redundant rules", testChecker(true, 200, "200", "2XX"))
	t.Run("match: 2XN==2XX", testChecker(true, 200, "200", "2X1"))

	t.Run("no match: no wildcard", testChecker(false, 500, "200"))
	t.Run("no match: with wildcard", testChecker(false, 500, "2XX"))
	t.Run("no match: lowercase wildcard", testChecker(false, 500, "2xx"))
	t.Run("no match: multiple rules", testChecker(false, 500, "200", "400"))
	t.Run("no match: multiple rules with wildcard", testChecker(false, 500, "2XX", "400"))
	t.Run("no match: redundant rules", testChecker(false, 500, "200", "2XX"))
}

func testChecker(match bool, statusCode int, rules ...string) func(t *testing.T) {
	return func(t *testing.T) {
		var checker httpstatcheck.Checker
		if len(rules) != 0 {
			checker.Insert(rules...)
		}
		gotMatch := checker.Check(statusCode)
		testr.New(t).Equal(gotMatch, match)
	}
}
