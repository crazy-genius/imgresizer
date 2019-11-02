package resizer

// Dimenstions represents img dimenstions struct
type Dimenstions struct {
	Height uint
	Width  uint
}

// AspectPreserveType describes what should to be preserved
type AspectPreserveType int

const (
	// NoPreserve means don't preserve any aspect rato
	NoPreserve AspectPreserveType = iota
	// PreserveWidth means programm should preserve aspectrato and calculate height by width
	PreserveWidth
	// PreserveHeight means programm should preserve aspectrato and calculate widht by height
	PreserveHeight
)

func preserveAspectRatoDimensions(currentDimenstions, desiredDimenstions Dimenstions, preserveType AspectPreserveType) Dimenstions {
	width := desiredDimenstions.Width
	hight := desiredDimenstions.Height

	if preserveType != NoPreserve && currentDimenstions.compareAspectRatio(desiredDimenstions) {
		switch preserveType {
		case PreserveHeight:
			width = currentDimenstions.calculateWidthBy(desiredDimenstions.Height)
		case PreserveWidth:
			hight = currentDimenstions.calculateHeightBy(desiredDimenstions.Width)
		}
	}

	return Dimenstions{
		hight,
		width,
	}
}

func (d Dimenstions) compareAspectRatio(withDimenstions Dimenstions) bool {
	return d.calculateAspectRatio() != withDimenstions.calculateAspectRatio()
}

func (d Dimenstions) calculateAspectRatio() uint {
	return (d.Width / d.Height)
}

func (d Dimenstions) calculateHeightBy(desiredWidth uint) uint {
	return uint(desiredWidth / d.calculateAspectRatio())
}

func (d Dimenstions) calculateWidthBy(desiredHeigh uint) uint {
	return uint(d.Height * d.calculateAspectRatio())
}
