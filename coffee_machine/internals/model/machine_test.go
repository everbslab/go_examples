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
		{"test1", fields{Deposit: 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New(tt.fields.Deposit, nil)
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
		{name: "Machine1", args: args{Deposit: 100, Milk: 0}, want: New(100, nil)},
		{name: "Machine2", args: args{Deposit: 0, Milk: 0}, want: New(0, nil)},
		{name: "Machine3", args: args{Deposit: 0, Milk: 50}, want: New(0, nil)},
		{name: "Machine4", args: args{}, want: New(0, nil)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.Deposit, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMachine_AvailableActions(t *testing.T) {
	type fields struct {
		Deposit float64
		Milk    uint
		actions []Command
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
			actions: []Command{depositCmd, exitCmd},
		}, strings.Join([]string{"deposit", "exit"}, sep)},

		{"action1", fields{
			Deposit: 100,
			Milk:    50,
			actions: []Command{exitCmd},
		}, strings.Join([]string{"exit"}, sep)},

		{"actionNULL", fields{
			Deposit: 100,
			Milk:    50,
			actions: []Command{},
		}, strings.Join([]string{}, sep)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Machine{
				Deposit: tt.fields.Deposit,
				actions: tt.fields.actions,
			}
			if got := c.AvailableActions(sep); got != tt.want {
				t.Errorf("AvailableActions() = %v, want %v", got, tt.want)
			}
		})
	}
}
