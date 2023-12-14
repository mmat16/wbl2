package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_readLines(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				input: strings.NewReader("1\n2\n3\n"),
			},
			want:    []string{"1", "2", "3"},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				input: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readLines(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("readLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeLines(t *testing.T) {
	type args struct {
		lines  []string
		output io.Writer
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			name: "test1",
			args: args{
				lines:  []string{"1", "2", "3"},
				output: &bytes.Buffer{},
			},
			wantOutput: "1\n2\n3\n",
			wantErr:    false,
		},
		{
			name: "test2",
			args: args{
				lines:  []string{"1", "2", "3"},
				output: nil,
			},
			wantOutput: "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := writeLines(tt.args.lines, tt.args.output)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) == tt.wantErr {
				return
			}
			if gotOutput := tt.args.output.(*bytes.Buffer).String(); gotOutput != tt.wantOutput {
				t.Errorf("writeLines() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_determineAndReadInput(t *testing.T) {
	tests := []struct {
		name      string
		testInput string
		want      []string
	}{
		{
			name:      "test1",
			testInput: "1\n2\n3\n",
			want:      []string{"1", "2", "3"},
		},
	}
	args = []string{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock Stdin
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			r, w, _ := os.Pipe()
			os.Stdin = r
			w.Write([]byte(tt.testInput))
			w.Close()

			got, err := determineAndReadInput()
			if err != nil {
				t.Errorf("determineAndReadInput() error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("determineAndReadInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortLines(t *testing.T) {
	type args struct {
		lines []string
		want  []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				lines: []string{"10", "4", "1", "9"},
				want:  []string{"1", "10", "4", "9"},
			},
		},
		{
			name: "test2",
			args: args{
				lines: []string{"b", "d", "a", "e"},
				want:  []string{"a", "b", "d", "e"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortLines(tt.args.lines)
			if !reflect.DeepEqual(tt.args.lines, tt.args.want) {
				t.Errorf("sortLines() = %v, want %v", tt.args.lines, tt.args.want)
			}
		})
	}
}

func Test_numericSort(t *testing.T) {
	type args struct {
		lines []string
		want  []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				lines: []string{"10", "4", "1", "9"},
				want:  []string{"1", "4", "9", "10"},
			},
		},
		{
			name: "test2",
			args: args{
				lines: []string{"10", "4", "1", "9", "a"},
				want:  []string{"a", "1", "4", "9", "10"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numericSort(tt.args.lines)
			if !reflect.DeepEqual(tt.args.lines, tt.args.want) {
				t.Errorf("numericSort() = %v, want %v", tt.args.lines, tt.args.want)
			}
		})
	}
}

func Test_reverseSort(t *testing.T) {
	type args struct {
		lines []string
		want  []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				lines: []string{"10", "4", "1", "9"},
				want:  []string{"9", "1", "4", "10"},
			},
		},
		{
			name: "test2",
			args: args{
				lines: []string{"10", "4", "1", "9", "a"},
				want:  []string{"a", "9", "1", "4", "10"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseSort(tt.args.lines)
			if !reflect.DeepEqual(tt.args.lines, tt.args.want) {
				t.Errorf("reverseSort() = %v, want %v", tt.args.lines, tt.args.want)
			}
		})
	}
}

func Test_uniqueSort(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				lines: []string{"10", "4", "1", "9", "10", "4", "1", "9"},
			},
			want: []string{"1", "10", "4", "9"},
		},
		{
			name: "test2",
			args: args{
				lines: []string{"a", "a", "a", "a"},
			},
			want: []string{"a"},
		},
		{
			name: "test3",
			args: args{
				lines: []string{"a", "b", "c", "d"},
			},
			want: []string{"a", "b", "c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniqueSort(tt.args.lines)
			sortLines(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniqueSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_items_Len(t *testing.T) {
	tests := []struct {
		name string
		i    items
		want int
	}{
		{
			name: "test1",
			i: items{
				{[]string{"1", "2", "3"}},
				{[]string{"1", "3", "2"}},
				{[]string{"2", "3", "1"}},
			},
			want: 3,
		},
		{
			name: "test2",
			i: items{
				{[]string{"1", "2", "3"}},
				{[]string{"1", "3", "2"}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Len(); got != tt.want {
				t.Errorf("items.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_items_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		it   items
		args args
		want bool
	}{
		{
			name: "test1",
			it: items{
				{[]string{"1", "2", "3"}},
				{[]string{"1", "3", "2"}},
				{[]string{"2", "3", "1"}},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want: true,
		},
		{
			name: "test2",
			it: items{
				{[]string{"1", "2", "3"}},
				{[]string{"1", "3", "2"}},
				{[]string{"2", "3", "1"}},
			},
			args: args{
				i: 2,
				j: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		column = 1
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.it.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("items.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_items_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		it   items
		args args
		want items
	}{
		{
			name: "test1",
			it: items{
				{[]string{"1", "2", "3"}},
				{[]string{"1", "3", "2"}},
				{[]string{"2", "3", "1"}},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want: items{
				{[]string{"1", "2", "3"}},
				{[]string{"2", "3", "1"}},
				{[]string{"1", "3", "2"}},
			},
		},
		{
			name: "test2",
			it: items{
				{[]string{"1", "2", "3"}},
				{[]string{"1", "3", "2"}},
				{[]string{"2", "3", "1"}},
				{[]string{"2", "3", "1"}},
			},
			args: args{
				i: 0,
				j: 3,
			},
			want: items{
				{[]string{"2", "3", "1"}},
				{[]string{"1", "3", "2"}},
				{[]string{"2", "3", "1"}},
				{[]string{"1", "2", "3"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.it.Swap(tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.it, tt.want) {
				t.Errorf("items.Swap() = %v, want %v", tt.it, tt.want)
			}
		})
	}
}

func Test_columnSort(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				lines: []string{"1 2 3", "1 3 2", "2 3 1"},
			},
			want: []string{"2 3 1", "1 3 2", "1 2 3"},
		},
		{
			name: "test2",
			args: args{
				lines: []string{"4 5 6", "1 3 0", "2 3 1"},
			},
			want: []string{"1 3 0", "2 3 1", "4 5 6"},
		},
	}
	for _, tt := range tests {
		column = 3
		t.Run(tt.name, func(t *testing.T) {
			columnSort(tt.args.lines)
			if !reflect.DeepEqual(tt.args.lines, tt.want) {
				t.Errorf("columnSort() = %v, want %v", tt.args.lines, tt.want)
			}
		})
	}
}
