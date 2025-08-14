package solution

import "testing"

func TestSubsetXORSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct{
		name string
		args args
		want int
	}{
		{
			name: "set of n elements",
			args: args{nums: []int{3,4,5,6,7,8}},
			want: 480,
		},
		{
			name: "empty set",
			args: args{nums: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubsetXORSum(tt.args.nums); got != tt.want {
				t.Errorf("SubsetXORSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountMaxOrSubsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct{
		name string
		args args
		want int
	}{
		{
			name: "set of n elements",
			args: args{nums: []int{3,2,1,5}},
			want: 6,
		},
		{
			name: "empty set",
			args: args{nums: []int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountMaxOrSubsets(tt.args.nums); got != tt.want {
				t.Errorf("CountMaxOrSubsets() = %v, want %v", got, tt.want)
			}
		})
	}

	
}