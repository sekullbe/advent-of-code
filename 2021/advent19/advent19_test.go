package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name         string
		args         args
		wantScanners []*scanner
	}{
		{
			name: "simple",
			args: args{inputText: "--- scanner 0 ---\n0,2,0\n4,1,0\n3,3,0\n\n--- scanner 1 ---\n-1,-1,0\n-5,0,0\n-2,1,0"},
			wantScanners: []*scanner{
				&scanner{
					id:          0,
					location:    beacon{0, 0, 0},
					beacons:     []beacon{beacon{0, 2, 0}, {4, 1, 0}, {3, 3, 0}},
					fingerprint: nil,
				},
				&scanner{
					id:          1,
					location:    beacon{0, 0, 0},
					beacons:     []beacon{beacon{-1, -1, 0}, {-5, 0, 0}, {-2, 1, 0}},
					fingerprint: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := parseInput(tt.args.inputText); !reflect.DeepEqual(got, tt.wantScanners) {
			//	t.Errorf("parseInput() = %v, want %v", got, tt.wantScanners)
			//}
			scanners := parseInput(tt.args.inputText)
			assert.Equalf(t, tt.wantScanners, scanners, "parseInput(%v)", tt.args.inputText)
		})
	}
}

func TestCalculateDistanceFingerprint(t *testing.T) {
	scanners := parseInput("--- scanner 0 ---\n0,2,0\n4,1,0\n3,3,0")
	s := scanners[0]
	fp := s.calculateDistanceFingerprint(false)
	_ = fp
}

func TestMatchFingerprints_SimpleData(t *testing.T) {
	scanners := parseInput("--- scanner 0 ---\n0,2,0\n4,1,0\n3,3,0\n\n--- scanner 1 ---\n-1,-1,0\n-5,0,0\n-2,1,0")
	setFingerprints(scanners)
	f0 := scanners[0].fingerprint
	f1 := scanners[1].fingerprint
	match, _ := matchFingerprints(f0, f1, 3)
	assert.True(t, match)

}

func TestMatchFingerprints_ComplexData(t *testing.T) {
	scanners := parseInput(testText)
	setFingerprints(scanners)
	f0 := scanners[0].fingerprint
	var matchingScanner *scanner
	for _, s := range scanners[1:] {
		log.Printf("Checking s0 vs s%d", s.id)
		match, _ := matchFingerprints(f0, s.fingerprint, 12)
		if match {
			matchingScanner = s
			log.Printf("Overmap match: s0 and s%d", s.id)
		}
	}
	assert.NotNil(t, matchingScanner)
	//assert.Equal(t, 1, matchingScanner.id)
}

func TestRun1(t *testing.T) {
	run1(testText)
}
