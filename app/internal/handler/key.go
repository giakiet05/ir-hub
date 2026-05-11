package handler

import (
	"fmt"
)

type IrKey struct {
	Address uint16
	Command uint16
}

func (k *IrKey) String() string {
	return fmt.Sprintf("A%d-C%d", k.Address, k.Command)
}
