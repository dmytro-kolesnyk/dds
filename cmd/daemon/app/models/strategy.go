package strategy

import "fmt"

type Strategy struct {
	name string
}

func ParseStrategy(value string) (*Strategy, error) {
	switch value {
	case "copy":
		return &Strategy{value}, nil
	case "fragment":
		return &Strategy{value}, nil
	case "fragment_copy":
		return &Strategy{value}, nil
	case "default":
		return &Strategy{value}, nil
	default:
		return nil, fmt.Errorf("unknown strategy: %v", value)
	}
}
