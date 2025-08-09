package solution

import (
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("TestMoreLeftArr", TestMoreLeftArr)
	t.Run("TestMirrorIndex", TestMirrorIndex)
	t.Run("TestPowLog", TestPowLog)
	t.Run("TestSumaDePotencias", TestSumaDePotencias)
	t.Run("TestMergeSort", TestMergeSort)
	t.Run("TestConstructMaximuimBinaryTree", TestConstructMaximumBinaryTree)
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
			args: args{arr: []int{1, 3, 5, 7, 9, 11}},
			want: true,
		},
		{
			name: "has mirror index to the right",
			args: args{arr: []int{-2, -1, 0, 1, 5}},
			want: true,
		},
		{
			name: "has mirror index in the center",
			args: args{arr: []int{-1, 0, 1, 4, 9, 10, 13}},
			want: true,
		},
		{
			name: "does not have mirror index",
			args: args{arr: []int{-1, 0, 1, 2, 3, 4, 5, 6}},
			want: false,
		},
		{
			name: "empty array",
			args: args{arr: []int{}},
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

func TestSumaDePotencias(t *testing.T) {
	type args struct {
		A [4][4]int
		n int
	}
	tests := []struct {
		name string
		args args
		want [4][4]int
	}{
		{
			name: "identity matrix sum",
			args: args{A: [4][4]int{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			}, n: 4},
			want: [4][4]int{
				{4, 0, 0, 0},
				{0, 4, 0, 0},
				{0, 0, 4, 0},
				{0, 0, 0, 4},
			},
		},
		{
			name: "some Matrix to the 2^kth power",
			args: args{A: [4][4]int{
				{1, 2, 3, 4},
				{8, 7, 6, 5},
				{-1, -2, -3, -4},
				{-8, -7, -6, -5},
			}, n: 8},
			want: [4][4]int{
				{-17, -16, -15, -14},
				{26, 25, 24, 23},
				{17, 16, 15, 14},
				{-26, -25, -24, -23},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumaDePotencias(tt.args.A, tt.args.n); got != tt.want {
				t.Errorf("SumaDePotencias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "all equal slice",
			args: args{arr: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "merge and Sort",
			args: args{arr: []int{1, 2934, -343, 19, 3, 359, -123, -32422342, 43}},
			want: []int{-32422342, -343, -123, 1, 3, 19, 43, 359, 2934},
		},
		{
			name: "one element array",
			args: args{arr: []int{1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "simple case",
			args: args{nums: []int{3, 2, 1, 6, 0, 5}},
			want: []int{6, 3, 5, 2, 0, 1},
		},
		{
			name: "single element",
			args: args{nums: []int{42}},
			want: []int{42},
		},
		{
			name: "empty array",
			args: args{nums: []int{}},
			want: []int{},
		},
		{
			name: "all increasing",
			args: args{nums: []int{1, 2, 3, 4}},
			want: []int{4, 3, 2, 1},
		},
		{
			name: "all decreasing",
			args: args{nums: []int{4, 3, 2, 1}},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := ConstructMaximumBinaryTree(tt.args.nums)
			got := GetLevelOrder(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConstructMaximumBinaryTree() level order = %v, want %v", got, tt.want)
			}
		})
	}
}
