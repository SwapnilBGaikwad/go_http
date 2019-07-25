package random

import "testing"

func TestGetRandom(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Sample Test", args{2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandom(tt.args.n); len(got) != tt.want {
				t.Errorf("GetRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}
