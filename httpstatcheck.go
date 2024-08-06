// Package httpstatcheck provides wildcard HTTP Status Code checking, e.g. 200 == "2XX".
package httpstatcheck

import (
	"strconv"
	"strings"
)

// Checker provides wildcard HTTP Status Code checking, e.g. 200 == "2XX".
// The zero value for Checker is an empty rule checker ready to use.
type Checker struct {
	root *node
}

// Insert inserts rules around HTTP Status Code.
//   - Support digits '0-9', e.g. 200 == "200".
//   - Support wildcard using char 'x' or 'X' (case insensitive), e.g. 200 == "2XX".
//   - Any rules that contains non 3 digits rule will be ignored, e.g. 200 != "2000".
//   - Any rules that contains non digits or wildcard will be ignored, e.g. 200 != "2NN".
//   - Digits after wildcard will considered as wildcard, e.g. 200 == "2X1".
func (c *Checker) Insert(rules ...string) {
	if c.root == nil {
		c.root = &node{}
	}

	for _, rule := range rules {
		if len(rule) != 3 {
			continue
		}
		rule = strings.ToLower(rule)

		n := c.root
		for _, char := range rule {
			index, isWildcard, isValid := charToIndex(char)
			if !isValid {
				break
			}

			if isWildcard {
				n.isWildcard = true
				break
			}

			if n.children[index] == nil {
				n.children[index] = &node{}
			}
			n = n.children[index]
		}
		n.isEnd = true
	}
}

// Check checks the status code to the inserted rules.
// Returns true if the status code is matched with the rules.
func (c *Checker) Check(statusCode int) bool {
	if c.root == nil {
		c.root = &node{}
	}

	code := strconv.Itoa(statusCode)
	if len(code) != 3 {
		return false
	}

	n := c.root
	for _, char := range code {
		if n.isWildcard {
			return true
		}
		index, _, _ := charToIndex(char)
		if n.children[index] != nil {
			n = n.children[index]
		} else {
			return false
		}
	}
	return n.isEnd
}

type node struct {
	children   [10]*node // index for '0-9'
	isWildcard bool      // wildcard 'x' or 'X'
	isEnd      bool
}

func charToIndex(char rune) (index int, isWildcard bool, isValid bool) {
	if char >= '0' && char <= '9' {
		return int(char - '0'), false, true // index for '0-9'
	}
	if char == 'x' {
		return 0, true, true
	}
	return 0, false, false
}
