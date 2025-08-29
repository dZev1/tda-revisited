package solution

import (
	"reflect"
	"testing"
)

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

func TestBinaryTreePaths(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct{
		name string
		args args
		want []string
	}{
		{
			name: "set of n elements",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "empty tree",
			args: args{root: nil},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryTreePaths(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryTreePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubsetSum(t *testing.T) {
	type args struct {	
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "set with subset sum",
			args: args{nums: []int{3, 34, 4, 12, 5, 2}, target: 9},
			want: true,
		},
		{	
			name: "set without subset sum",
			args: args{nums: []int{3, 34, 4, 12, 5, 2}, target: 30},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubsetSum(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SubsetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

