package grep

import (
	"reflect"
	"regexp"
	"testing"
)

func Test_reMatcher_match(t *testing.T) {
	type fields struct {
		pattern *regexp.Regexp
		line    string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "test1",
			fields: fields{
				pattern: regexp.MustCompile(`test`),
				line:    "test",
			},
			want: true,
		},
		{
			name: "test2",
			fields: fields{
				pattern: regexp.MustCompile(`test`),
				line:    "kjsh",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := &reMatcher{
				pattern: tt.fields.pattern,
				line:    tt.fields.line,
			}
			if got := rm.match(); got != tt.want {
				t.Errorf("reMatcher.match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reMatcher_addLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				line: "test1",
			},
		},
		{
			name: "test2",
			args: args{
				line: "test2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rm reMatcher
			rm.addLine(tt.args.line)
			if rm.line != tt.args.line {
				t.Errorf("reMatcher.addLine() = %v, want %v", rm.line, tt.args.line)
			}
		})
	}
}

func Test_reMatcher_getLine(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test1",
			want: "test1",
		},
		{
			name: "test2",
			want: "test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rm reMatcher
			rm.line = tt.want
			if got := rm.getLine(); got != tt.want {
				t.Errorf("reMatcher.getLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringMatcher_match(t *testing.T) {
	type fields struct {
		pattern string
		line    string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "test1",
			fields: fields{
				pattern: "test",
				line:    "test",
			},
			want: true,
		},
		{
			name: "test2",
			fields: fields{
				pattern: "test",
				line:    "kjsh",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &stringMatcher{
				pattern: tt.fields.pattern,
				line:    tt.fields.line,
			}
			if got := sm.match(); got != tt.want {
				t.Errorf("stringMatcher.match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringMatcher_addLine(t *testing.T) {
	type fields struct {
		pattern string
		line    string
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test1",
			fields: fields{
				pattern: "test1",
				line:    "test1",
			},
			args: args{
				line: "test1",
			},
		},
		{
			name: "test2",
			fields: fields{
				pattern: "test2",
				line:    "test2",
			},
			args: args{
				line: "test2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &stringMatcher{
				pattern: tt.fields.pattern,
				line:    tt.fields.line,
			}
			sm.addLine(tt.args.line)
			if sm.line != tt.args.line {
				t.Errorf("stringMatcher.addLine() = %v, want %v", sm.line, tt.args.line)
			}
		})
	}
}

func Test_stringMatcher_getLine(t *testing.T) {
	type fields struct {
		pattern string
		line    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test1",
			fields: fields{
				pattern: "test1",
				line:    "test1",
			},
			want: "test1",
		},
		{
			name: "test2",
			fields: fields{
				pattern: "test2",
				line:    "test2",
			},
			want: "test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &stringMatcher{
				pattern: tt.fields.pattern,
				line:    tt.fields.line,
			}
			if got := sm.getLine(); got != tt.want {
				t.Errorf("stringMatcher.getLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newQueue(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want *queue
	}{
		{
			name: "test1",
			args: args{
				limit: 1,
			},
			want: &queue{lines: []string{}, limit: 2},
		},
		{
			name: "test2",
			args: args{
				limit: 2,
			},
			want: &queue{lines: []string{}, limit: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newQueue(tt.args.limit)
			if got.limit != tt.want.limit || len(got.lines) != len(tt.want.lines) {
				t.Errorf("newQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_push(t *testing.T) {
	type fields struct {
		lines []string
		limit int
	}
	type args struct {
		line string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "test1",
			fields: fields{
				lines: []string{},
				limit: 2,
			},
			args: args{
				line: "test1",
			},
			want: []string{"test1"},
		},
		{
			name: "test2",
			fields: fields{
				lines: []string{"test1"},
				limit: 2,
			},
			args: args{
				line: "test2",
			},
			want: []string{"test1", "test2"},
		},
		{
			name: "test3",
			fields: fields{
				lines: []string{"test1", "test2"},
				limit: 2,
			},
			args: args{
				line: "test3",
			},
			want: []string{"test2", "test3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				lines: tt.fields.lines,
				limit: tt.fields.limit,
			}
			q.push(tt.args.line)
			if !reflect.DeepEqual(q.lines, tt.want) {
				t.Errorf("queue.push() = %v, want %v", q.lines, tt.want)
			}
		})
	}
}

func Test_queue_popAll(t *testing.T) {
	type fields struct {
		lines []string
		limit int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "test1",
			fields: fields{
				lines: []string{"test1", "test2"},
				limit: 2,
			},
			want: []string{"test1", "test2"},
		},
		{
			name: "test2",
			fields: fields{
				lines: []string{"test1", "test2", "test3"},
				limit: 3,
			},
			want: []string{"test1", "test2", "test3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				lines: tt.fields.lines,
				limit: tt.fields.limit,
			}
			if got := q.popAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queue.popAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
