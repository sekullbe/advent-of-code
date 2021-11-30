package advent14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bitmask_applyBitmaskToValue(t *testing.T) {

	mask := createEmptyBitmask()
	mask[1] = ZERO
	mask[6] = ONE

	assert.Equal(t, 73, mask.applyBitmaskToValue(11))
	assert.Equal(t, 101, mask.applyBitmaskToValue(101))
	assert.Equal(t, 64, mask.applyBitmaskToValue(0))
	assert.Equal(t, 64, mask.applyBitmaskToValue(2))
}

func Test_bitmask_createFromString(t *testing.T) {

	mask := createBitmask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	assert.Equal(t, ONE, mask[6])
	assert.Equal(t, ZERO, mask[1])
	assert.Equal(t, X, mask[0])

	assert.Equal(t, 73, mask.applyBitmaskToValue(11))
	assert.Equal(t, 101, mask.applyBitmaskToValue(101))
	assert.Equal(t, 64, mask.applyBitmaskToValue(0))
	assert.Equal(t, 64, mask.applyBitmaskToValue(2))


}

func Test_parseAddress(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name     string
		args     args
		wantAddr int
		wantVal  int
	}{
		{
			name: "1",
			args: args{line:"mem[8] = 11"},
			wantAddr: 8,
			wantVal: 11,
		},
		{
			name: "2",
			args: args{line:"mem[6532] = 103013119"},
			wantAddr: 6532,
			wantVal: 103013119,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddr, gotVal := parseAddress(tt.args.line)
			if gotAddr != tt.wantAddr {
				t.Errorf("parseAddress() gotAddr = %v, want %v", gotAddr, tt.wantAddr)
			}
			if gotVal != tt.wantVal {
				t.Errorf("parseAddress() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}
