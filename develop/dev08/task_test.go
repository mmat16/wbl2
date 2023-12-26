package minishell

import "testing"

func Test_execCmd(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{"cd /"}, false},
		{"test2", args{"pwd"}, false},
		{"test3", args{"echo test"}, false},
		{"test4", args{"ps"}, false},
		{"test5", args{"we"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := execCmd(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("execCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cd(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{[]string{"cd", "/"}}, false},
		{"test2", args{[]string{"cd", "ejkrhwe"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cd(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("cd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pwd(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{[]string{"pwd"}}, false},
		{"test2", args{[]string{"pwd", "test"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := pwd(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("pwd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ps(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{[]string{"ps"}}, false},
		{"test2", args{[]string{"ps", "test"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ps(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("ps() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
