package enum

type Color string

const (
	ColorVerificationRequired Color = "#2baae8"
	ColorError                Color = "#ff0101"
	ColorSuccess              Color = "#00dd0b"
	ColorSpecialMessage       Color = "#6410ff"
	ColorDevelopment          Color = "#eb6200"
)

func (c Color) String() string { //* (c Color) is essentially this or self; it attaches itself to Color
	return string(c)
}
