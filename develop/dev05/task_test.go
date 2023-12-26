package grep

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_defineMatcher(t *testing.T) {
	type args struct {
		pattern   string
		mockFixed bool
	}
	tests := []struct {
		name    string
		args    args
		want    matcher
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				pattern:   "test",
				mockFixed: true,
			},
			want: &stringMatcher{
				pattern: "test",
				line:    "",
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				pattern:   "test",
				mockFixed: false,
			},
			want: &reMatcher{
				pattern: regexp.MustCompile("test"),
				line:    "",
			},
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				pattern:   "[a-z",
				mockFixed: false,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixed = tt.args.mockFixed
			got, err := defineMatcher(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("defineMatcher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defineMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchInLine(t *testing.T) {
	type args struct {
		matcher      matcher
		linesBefore  *queue
		lnCounter    *int
		matchCounter *int
		mockFixed    bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				matcher: &reMatcher{
					pattern: regexp.MustCompile("test"),
					line:    "test test test",
				},
				linesBefore:  &queue{},
				lnCounter:    new(int),
				matchCounter: new(int),
				mockFixed:    false,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				matcher: &reMatcher{
					pattern: regexp.MustCompile("test"),
					line:    "jkh awkfe asdfjkh",
				},
				linesBefore:  &queue{},
				lnCounter:    new(int),
				matchCounter: new(int),
				mockFixed:    false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		fixed = tt.args.mockFixed
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInLine(tt.args.matcher, tt.args.linesBefore, tt.args.lnCounter, tt.args.matchCounter); got != tt.want {
				t.Errorf("searchInLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
