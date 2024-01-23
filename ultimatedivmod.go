package piscine

func UltimateDivMod(a *int, b *int) {
	tmp := *a
	*a = tmp / *b
	*b = tmp % *b
}
