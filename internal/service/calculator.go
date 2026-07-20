package service

import (
	"fmt"
	"hello-wails/internal/domain"
)

type CalculatorService struct {}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (c *CalculatorService) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("invalid b: %w", domain.ErrDivisionByZero)
	}

	return a / b, nil
}
