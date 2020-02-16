// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Package cmp defines Min and Max functions for a variable number of int arguments.
package cmp

// int is a generic type for Min and Max (imported from github.com/cheekybits/genny/generic). Concrete types have to support comparison via operators < and >.
//
// The package contains subpackages where int is automatically replaced with concrete types.
//
// The int subpackage is required for tests. Make sure to include the following comment in your source code:
//
// //go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "Int=int"

// Min returns the minimum of a variable number of int arguments. Panics if called with no arguments.
func Min(values ...int) int {
	if len(values) < 1 {
		panic("Min requires at least one argument")
	}
	m := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < m {
			m = values[i]
		}
	}
	return m
}

// Max returns the maximum of a variable number of int arguments. Panics if called with no arguments.
func Max(values ...int) int {
	if len(values) < 1 {
		panic("Min requires at least one argument")
	}
	m := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > m {
			m = values[i]
		}
	}
	return m
}
