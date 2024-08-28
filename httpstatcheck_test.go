package httpstatcheck_test

import (
	"fmt"
	"regexp"
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

func BenchmarkHTTPStatCheck(b *testing.B) {
	var checker httpstatcheck.Checker
	checker.Insert("2XX")
	assert := testr.New(b)
	for i := 0; i < b.N; i++ {
		match := checker.Check(200)
		assert.Equal(match, true)
	}
}

func BenchmarkRegex(b *testing.B) {
	regex := regexp.MustCompile(`^2\d{2}$`)
	assert := testr.New(b)
	for i := 0; i < b.N; i++ {
		match := regex.MatchString("200")
		assert.Equal(match, true)
	}
}

func ExampleChecker() {
	var checker httpstatcheck.Checker
	checker.Insert("2XX", "400", "500", "3X1")
	fmt.Println(checker.Check(200))
	fmt.Println(checker.Check(201))
	fmt.Println(checker.Check(400))
	fmt.Println(checker.Check(401))
	fmt.Println(checker.Check(500))
	fmt.Println(checker.Check(501))
	fmt.Println(checker.Check(300), "3X1 will be considered as 3XX")
	// Output:
	// true
	// true
	// true
	// false
	// true
	// false
	// true 3X1 will be considered as 3XX
}

func TestIsEmpty(t *testing.T) {
	is := testr.New(t)

	var checker httpstatcheck.Checker
	is.Equal(checker.IsEmpty(), true, testr.WithMessage("IsEmpty before insert"))

	checker.Insert()
	is.Equal(checker.IsEmpty(), true, testr.WithMessage("IsEmpty after insert empty"))

	checker.Check(200)
	is.Equal(checker.IsEmpty(), true, testr.WithMessage("IsEmpty after check against empty rules"))

	checker.Insert("2XX")
	is.Equal(checker.IsEmpty(), false, testr.WithMessage("IsEmpty after insert"))
}
