package sum_zero

import (
	"reflect"
	"testing"
)

func Test_sumZero(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				n: 5,
			},
			want: []int{1, 2, 0, -1, -2},
		},
		{
			name: "Example2",
			args: args{
				n: 3,
			},
			want: []int{1, 0, -1},
		},
		{
			name: "Example3",
			args: args{
				n: 1,
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumZero(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
