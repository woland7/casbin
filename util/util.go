// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"regexp"
	"strings"
	"sort"
)

// Escape the dots in the assertion, because the expression evaluation doesn't support such variable names.
func EscapeAssertion(s string) string {
	return strings.Replace(s, ".", "_", -1)
}

// Translate the ABAC attributes into functions.
func FixAttribute(s string) string {
	reg := regexp.MustCompile("(r|p)\\.(sub|obj|act)\\.([A-Za-z0-9]*)")
	res := reg.ReplaceAllString(s, "${2}Attr($1.$2, \"$3\")")

	return res
}

// Determine whether two string arrays are identical.
func ArrayEquals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Determine whether two 2-dimensional string arrays are identical.
func Array2DEquals(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !ArrayEquals(v, b[i]) {
			return false
		}
	}
	return true
}

// Remove any duplicated elements in a string array.
func ArrayRemoveDuplicates(s *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *s {
		if !found[x] {
			found[x] = true
			(*s)[j] = (*s)[i]
			j++
		}
	}
	*s = (*s)[:j]
}

// Get a printable string for a string array.
func ArrayToString(s []string) string {
	return strings.Join(s, ", ")
}

// Determine whether two string sets are identical.
func SetEquals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}