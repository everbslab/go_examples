package model

import (
	"reflect"
	"testing"
)

func TestNewCoffee(t *testing.T) {
	type args struct {
		w     Ingredient
		m     Ingredient
		s     Ingredient
		b     Ingredient
		price float64
	}
	tests := []struct {
		name string
		args args
		want *Coffee
	}{
		{"test1", args{}, &Coffee{
			water: 0,
			milk:  0,
			sugar: 0,
			beans: 0,
			price: 0,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCoffee(tt.args.w, tt.args.m, tt.args.s, tt.args.b, tt.args.price); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCoffee() = %v, want %v", got, tt.want)
			}
		})
	}
}
