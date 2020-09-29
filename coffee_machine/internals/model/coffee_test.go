package model

import (
	"reflect"
	"testing"
)

func TestNewCoffee(t *testing.T) {
	type args struct {
		water uint
		milk  uint
		sugar uint
		beans uint
		price float64
	}
	tests := []struct {
		name string
		args args
		want *Coffee
	}{
		{"test", args{water: 50, milk: 50, sugar: 50, beans: 50, price: 1.00}, &Coffee{
			Water: 50,
			Milk:  50,
			Sugar: 50,
			Beans: 50,
			Price: 1.00,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCoffee(tt.args.water, tt.args.milk, tt.args.sugar, tt.args.beans, tt.args.price); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCoffee() = %v, want %v", got, tt.want)
			}
		})
	}
}
