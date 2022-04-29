package volume_control

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVolume(t *testing.T) {
	volumeGetter := func() (int, error) {
		return 27, nil
	}
	GetVolume = volumeGetter
	res, err := GetCurrentVolume()
	assert.Equal(t, res, 27)
	assert.Nil(t, err)

	volumeGetter = func() (int, error) {
		return 0, errors.New("I failed")
	}
	GetVolume = volumeGetter
	res, err = GetCurrentVolume()
	assert.Equal(t, res, -1)
	assert.Error(t, err, "I failed")
}

func TestChangeVolume(t *testing.T) {
	volumeGetter := func() (int, error) {
		return 27, nil
	}
	GetVolume = volumeGetter
	mockIncreaseVolume := func(i int) error {
		if i > 10 {
			return errors.New("Throwing an error")
		}
		return nil
	}
	IncreaseVolume = mockIncreaseVolume
	type args struct {
		increment int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "no error", args: args{1}, wantErr: false},
		{name: "error", args: args{11}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := changeVolume(tt.args.increment)
			assert.Equal(t, got, 27)
			if (err != nil) != tt.wantErr {
				t.Errorf("TurnUpVolume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTurnUpVolume(t *testing.T) {
	volumeChanger := func(howMuch int) (int, error) {
		if howMuch > 10 {
			return -1, errors.New("Error")
		}
		return howMuch + 20, nil
	}
	changeVolume = volumeChanger
	type args struct {
		increment int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Change Volume No Error", args{1}, 21, false},
		{"Change Volume Error", args{11}, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TurnUpVolume(tt.args.increment)
			if (err != nil) != tt.wantErr {
				t.Errorf("TurnUpVolume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TurnUpVolume() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTurnDownVolume(t *testing.T) {
	volumeChanger := func(howMuch int) (int, error) {
		if howMuch > 0 {
			return -1, errors.New("Error")
		}
		return howMuch + 20, nil
	}
	changeVolume = volumeChanger
	type args struct {
		increment int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Change Volume No Error", args{1}, 19, false},
		{"Change Volume Error", args{-11}, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TurnDownVolume(tt.args.increment)
			if (err != nil) != tt.wantErr {
				t.Errorf("TurnUpVolume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TurnUpVolume() got = %v, want %v", got, tt.want)
			}
		})
	}
}
