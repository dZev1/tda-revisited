package solution

import "testing"

func TestAll(t *testing.T) {
	t.Run("TestMoreLeftArr", TestMoreLeftArr)
	t.Run("TestMirrorIndex", TestMirrorIndex)
	t.Run("TestPowLog", TestPowLog)
}

func TestMoreLeftArr(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is more left",
			args: args{arr: []int{8, 6, 7, 4, 5, 1, 3, 2}},
			want: true,
		},
		{
			name: "is not more left",
			args: args{arr: []int{8, 4, 7, 6, 5, 1, 3, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoreLeftArr(tt.args.arr); got != tt.want {
				t.Errorf("MoreLeftArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMirrorIndex(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has mirror index to the left",
			args: args{arr: []int{1,3,5,7,9,11}},
			want: true,
		},
		{
			name: "has mirror index to the right",
			args: args{arr: []int{-2,-1,0,1,5}},
			want: true,
		},
		{
			name: "has mirror index in the center",
			args: args{arr: []int{-1,0,1,4,9,10,13}},
			want: true,
		},
		{
			name: "does not have mirror index",
			args: args{arr: []int{-1,0,1,2,3,4,5,6}},
			want: false,
		},
		{
			name: "empty array",
			args : args{arr: []int {}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MirrorIndex(tt.args.arr); got != tt.want {
				t.Errorf("MirroIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestPowLog(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2^3 = 8",
			args: args{n: 2, k: 3},
			want: 8,
		},
		{
			name: "2^32 = 4294967296",
			args: args{n: 2, k: 32},
			want: 4294967296,
		},
		{
			name: "N^0 = 1",
			args: args{n: 30, k: 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PowLog(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("PowLog() = %v, want %v", got, tt.want)
			}
		})
	}
}