package anagrams

import (
	"reflect"
	"slices"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	type args struct {
		words *[]string
	}
	tests := []struct {
		name string
		args args
		want *map[string]*[]string
	}{
		{
			name: "regular test",
			args: args{
				words: &[]string{
					"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "клуб", "булк",
					"каблук", "клубок", "клуб", "букл", "клуб", "клубок", "клубок", "клубок", "клубок",
				},
			},
			want: &map[string]*[]string{
				"листок": {"листок", "слиток", "столик"},
				"букл":   {"букл", "булк", "клуб"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name: "test for empty slice",
			args: args{
				words: &[]string{},
			},
			want: &map[string]*[]string{},
		},
		{
			name: "test for slice with one element",
			args: args{
				words: &[]string{"пятак"},
			},
			want: &map[string]*[]string{},
		},
		{
			name: "test for slice with same words in different cases",
			args: args{
				words: &[]string{"пятак", "Пятак", "ПЯТАК", "Тяпка", "тяпка", "ТЯПКА"},
			},
			want: &map[string]*[]string{
				"пятак": {"пятак", "тяпка"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAnagrams(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_order(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "regular test",
			args: args{
				word: "пятак",
			},
			want: "акптя",
		},
		{
			name: "test for empty string",
			args: args{
				word: "",
			},
			want: "",
		},
		{
			name: "test for string with one letter",
			args: args{
				word: "п",
			},
			want: "п",
		},
		{
			name: "test for string with same letters",
			args: args{
				word: "пппп",
			},
			want: "пппп",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := order(tt.args.word); got != tt.want {
				t.Errorf("order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "regular test",
			args: args{
				words: []string{"пятак", "пятак", "пятка", "пятка", "тяпка", "тяпка"},
			},
			want: []string{"пятак", "пятка", "тяпка"},
		},
		{
			name: "test for empty slice",
			args: args{
				words: []string{},
			},
			want: []string{},
		},
		{
			name: "test for slice with one element",
			args: args{
				words: []string{"пятак"},
			},
			want: []string{"пятак"},
		},
		{
			name: "test for slice with duplicates only",
			args: args{
				words: []string{"пятак", "пятак", "пятак", "пятак", "пятак", "пятак"},
			},
			want: []string{"пятак"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.words)
			slices.Sort(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
