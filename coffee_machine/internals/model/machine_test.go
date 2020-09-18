package model

import (
	"reflect"
	"testing"
)

func TestMachine_PrintState(t *testing.T) {
	type fields struct {
		Deposit int
		Milk    int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test1", fields{Deposit: 100, Milk: 50}, "Coffee Machine has:\n50 ml of milk\n$100 of money\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMachine(tt.fields.Deposit, tt.fields.Milk)
			got := m.PrintState()
			if got != tt.want {
				t.Errorf("PrintState() = `%v`.\nWant = `%v`", got, tt.want)
			}
		})
	}
}

func TestNewMachine(t *testing.T) {
	type args struct {
		Deposit int
		Milk    int
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
