package model

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestMachine_PrintState(t *testing.T) {
	type fields struct {
		Deposit float64
		Milk    uint
	}

	var wanted = "Coffee Machine has:\n%d ml of milk\n$%.2f of money"

	tests := []struct {
		name   string
		fields fields
	}{
		{"test1", fields{Deposit: 100, Milk: 50}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMachine(tt.fields.Deposit, tt.fields.Milk)
			got := m.PrintState()
			wanted := fmt.Sprintf(wanted, tt.fields.Milk, tt.fields.Deposit)
			if got != wanted {
				t.Errorf("PrintState() = `%v`.\nWant = `%v`", got, wanted)
			}
		})
	}
}

func TestNewMachine(t *testing.T) {
	type args struct {
		Deposit float64
		Milk    uint
	}

	tests := []struct {
		name string
		args args
		want *Machine
	}{
		{name: "Machine1", args: args{Deposit: 100, Milk: 0}, want: NewMachine(100, 0)},
		{name: "Machine2", args: args{Deposit: 0, Milk: 0}, want: NewMachine(0, 0)},
		{name: "Machine3", args: args{Deposit: 0, Milk: 50}, want: NewMachine(0, 50)},
		{name: "Machine4", args: args{}, want: NewMachine(0, 0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMachine(tt.args.Deposit, tt.args.Milk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMachine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMachine_AvailableActions(t *testing.T) {
	type fields struct {
		Deposit float64
		Milk    uint
		actions []string
	}

	var sep = ", "

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"actionsTwo", fields{
			Deposit: 100,
			Milk:    50,
			actions: []string{DepositAction, ExitAction},
		}, strings.Join([]string{DepositAction, ExitAction}, sep)},

		{"action1", fields{
			Deposit: 100,
			Milk:    50,
			actions: []string{ExitAction},
		}, strings.Join([]string{ExitAction}, sep)},

		{"actionNULL", fields{
			Deposit: 100,
			Milk:    50,
			actions: []string{},
		}, strings.Join([]string{}, sep)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Machine{
				Deposit: tt.fields.Deposit,
				Milk:    tt.fields.Milk,
				actions: tt.fields.actions,
			}
			if got := c.AvailableActions(sep); got != tt.want {
				t.Errorf("AvailableActions() = %v, want %v", got, tt.want)
			}
		})
	}
}
