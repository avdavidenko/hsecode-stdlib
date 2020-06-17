//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

package cmp

// good IDEs will put the import statement automatically
import "github.com/cheekybits/genny/generic"

type ValueType generic.Number

func Max(values ...ValueType) ValueType {
	if len(values) == 0 {
		panic("at least one argument required")
	}
	m := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > m {
			m = values[i]
		}
	}
	return m
}

func Min(values ...ValueType) ValueType {
	if len(values) == 0 {
		panic("at least one argument required")
	}
	m := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < m {
			m = values[i]
		}
	}
	return m
}
