package pattern

type RunePattern struct {
	BasePattern
}

func (runePattern *RunePattern) Accept(pattern Pattern, orientation Orientation) bool {
	return true
}
