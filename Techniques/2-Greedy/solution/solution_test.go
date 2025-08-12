package solution

import (
	"greedy/hashset"
	"testing"
)

func TestRunAll(t *testing.T) {
	t.Run("TestParejasDeBaile", TestParejasDeBaile)
	t.Run("TestSumaSelectiva", TestSumaSelectiva)
	t.Run("TestSumaGolosa", TestSumaGolosa)
	t.Run("TestLongestPalindrome", TestLongestPalindrome)
	t.Run("TestCanJump", TestCanJump)
}

func TestParejasDeBaile(t *testing.T) {
	type args struct {
		b1 []int
		b2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "slice b1 mayor de longitud que b2",
			args: args{
				b1: []int{1, 2, 4, 6},
				b2: []int{1, 5, 5, 7, 9},
			},
			want: 3,
		},
		{
			name: "no hay parejas",
			args: args{
				b1: []int{3, 5, 7, 9},
				b2: []int{1, 1, 1, 1},
			},
			want: 0,
		},
		{
			name: "slice b2 mayor de longitud que b1",
			args: args{
				b1: []int{1, 1, 1, 1, 1},
				b2: []int{1, 2, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParejasDeBaile(tt.args.b1, tt.args.b2); got != tt.want {
				t.Errorf("ParejasDeBaile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumaSelectiva(t *testing.T) {
	type args struct {
		X hashset.HashSet[int]
		k int
	}
	tests := []struct {
		name  string
		args  args
		want  *hashset.HashSet[int]
		want1 int
	}{
		{
			name: "X conjunto vacio no hay S",
			args: args{X: *hashset.New[int](), k: 0},
			want: hashset.New[int](),
			want1: 0,
		},
		{
			name: "X con longitud mayor que k",
			args: args{X: *makeHashset([]int{10,39,20,14,58,9,35}), k: 4},
			want: makeHashset([]int{58,39,35,20}),
			want1: 152,
		},
		{
			name: "X con longitud igual a k",
			args: args{X: *makeHashset([]int{1,2,3,4,5,6,7,8,9,10}), k: 10},
			want: makeHashset([]int{1,2,3,4,5,6,7,8,9,10}),
			want1: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SumaSelectiva(tt.args.X, tt.args.k)
			if !got.Equals(tt.want) {
				t.Errorf("SumaSelectiva() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SumaSelectiva() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSumaGolosa(t *testing.T) {
	type args struct {
		set []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "conjunto sin numeros repetidos de n elementos",
			args: args{[]int{1,2,5}},
			want: 11,
		},
		{
			name: "conjunto con numeros repetidos de n elementos",
			args: args{[]int{1,1,2,3,4,4,5}},
			want: 53,
		},
		{
			name: "conjunto con solo un numero",
			args: args{[]int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumaGolosa(tt.args.set); got != tt.want {
				t.Errorf("SumaGolosa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "string of even length",
			args: args{s: "abccccdd"},
			want: 7,
		},
		{
			name: "string of odd length",
			args: args{s: "abcad"},
			want: 3,
		},
		{
			name: "string of length 1",
			args: args{s: "a"},
			want: 1,
		},
		{
			name: "string of all different chars",
			args: args{s:"abcdefghijk"},
			want: 1,
		},
		{
			name: "empty string",
			args: args{s: ""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "has a path to the end",
			args: args{nums: []int{2,3,1,1,4}},
			want: true,
		},
		{
			name: "does not have a path to the end",
			args: args{nums: []int{3,2,1,0,4}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanJump(tt.args.nums); got != tt.want {
				t.Errorf("CanJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinPartitions(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "very large number",
			args: args{n: "27346209830709182346"},
			want: 9,
		},
		{
			name: "small number",
			args: args{n:"82734"},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinPartitions(tt.args.n); got != tt.want {
				t.Errorf("MinPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinMovesToSeat(t *testing.T) {
	type args struct {
		seats    []int
		students []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n students - n seats",
			args: args{seats: []int{3,1,5}, students: []int{2,7,4}},
			want: 4,
		},
		{
			name: "no students - no seats",
			args: args{seats: []int{}, students: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMovesToSeat(tt.args.seats, tt.args.students); got != tt.want {
				t.Errorf("MinMovesToSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}













// funcion para generar HashSets a partir de slices
func makeHashset(nums []int) *hashset.HashSet[int] {
	set := hashset.New[int]()

	for _, num := range nums {
		set.Add(num)
	}

	return set
}