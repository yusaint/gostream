package functions

import "golang.org/x/exp/constraints"

func Gte[T constraints.Ordered](e1 T, e2 T) bool { return e1 >= e2 }
func Lte[T constraints.Ordered](e1 T, e2 T) bool { return !(e1 >= e2) }

func Sum[T constraints.Ordered](e1, e2 T) T { return e1 + e2 }

var (
	IntGte = Gte[int]
	IntLte = Lte[int]
	IntSum = Sum[int]
)
