package domain

type StrategyType int

const (
	ALL StrategyType = iota
	TOP
	CROSS
)

func (s StrategyType) String() string {
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
