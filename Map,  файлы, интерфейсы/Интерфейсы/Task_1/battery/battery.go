package battery

import (
	"errors"
	"fmt"
	"strings"
)

const chargeSymbol = "1"

type Battery struct {
	maxPower, power int
}

func (b *Battery) String() string {
	return fmt.Sprintf("[%10s]", strings.Repeat("X", b.power))
}

func (b *Battery) Charge(power string) error {
	tmp := strings.Count(power, chargeSymbol)

	if tmp > b.maxPower {
		return errors.New("переизбыток заряда")
	}

	b.power = tmp

	return nil
}

func NewBattery(maxPower int) *Battery {
	return &Battery{maxPower: maxPower}
}
