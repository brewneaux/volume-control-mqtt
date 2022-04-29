package volume_control

import (
	"github.com/itchyny/volume-go"
)

var (
	IncreaseVolume = volume.IncreaseVolume
	GetVolume      = volume.GetVolume
	changeVolume   = volumeChange
)

func GetCurrentVolume() (int, error) {
	vol, err := GetVolume()
	if err != nil {
		return -1, err
	}
	return vol, nil
}

func volumeChange(howMuch int) (int, error) {
	currentVolume, _ := GetCurrentVolume()
	errIncrease := IncreaseVolume(howMuch)

	currentVolume, _ = GetCurrentVolume()

	if errIncrease != nil {
		return currentVolume, errIncrease
	}
	return currentVolume, nil
}

func TurnUpVolume(increment int) (int, error) {
	currentVolume, err := changeVolume(increment)
	return currentVolume, err
}

func TurnDownVolume(increment int) (int, error) {
	currentVolume, err := changeVolume(-increment)
	return currentVolume, err
}
