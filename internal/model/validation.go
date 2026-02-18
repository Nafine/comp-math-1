package model

import (
	"fmt"
	"strconv"
	"strings"
)

func nValidator(s string) error {
	if len(s) > 2 {
		return fmt.Errorf("matrix size is too big")
	}

	c := strings.ReplaceAll(s, " ", "")
	n, err := strconv.ParseInt(c, 10, 64)

	if err != nil {
		return fmt.Errorf("matrix size is invalid")
	}

	if n < 1 || n > 20 {
		return fmt.Errorf("matrix size must be in [1;20]")
	}

	return nil
}

func epsValidator(s string) error {
	c := strings.ReplaceAll(s, " ", "")
	eps, err := strconv.ParseFloat(c, 64)

	if err != nil {
		return fmt.Errorf("eps is invalid")
	} else if eps <= 0 {
		return fmt.Errorf("eps must be greater than 0")
	}

	return nil
}

func matrixCellValidator(s string) error {
	c := strings.ReplaceAll(s, " ", "")
	_, err := strconv.ParseInt(c, 10, 64)

	return err
}
