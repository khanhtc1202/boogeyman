package domain

type RankerStrategyType int

const (
	ALL RankerStrategyType = iota
	TOP
	CROSS
)

func (s RankerStrategyType) String() string {
	switch s {
	case ALL:
		return "ALL"
	case CROSS:
		return "CROSS"
	case TOP:
		return "TOP"
	}
	return "ALL"
}
