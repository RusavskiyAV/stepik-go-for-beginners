package main

import (
	batteryPackage "Task_1/battery"
	"fmt"
)

const powerStatusLen = 10

func main() {
	var power string

	if _, err := fmt.Scan(&power); err != nil || len(power) != powerStatusLen {
		panic("ошибка ввода статуса батареи")
	}

	battery := batteryPackage.NewBattery(powerStatusLen)
	err := battery.Charge(power)

	if err != nil {
		panic(err)
	}

	fmt.Print(battery)
}
