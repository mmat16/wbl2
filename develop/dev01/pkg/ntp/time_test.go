package ntp

import (
	"testing"
)

func TestGetPreciseTime(t *testing.T) {
	type args struct {
		ntpPool string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{"0.beevik-ntp.pool.ntp.org"}, false},
		{"test2", args{"hagfrhgqwehuihasdfbasdvuqiwehf"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetPreciseTime(tt.args.ntpPool)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPreciseTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
