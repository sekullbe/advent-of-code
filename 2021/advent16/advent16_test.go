package main

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_decodeHex(t *testing.T) {
	type args struct {
		inputText string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "basic",
			args: args{inputText: "8"},
			want: []byte("1000"),
		},
		{
			name: "multichar",
			args: args{inputText: "8F01"},
			want: []byte("1000111100000001"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeHex(tt.args.inputText); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryStringInBytesToInt(t *testing.T) {
	type args struct {
		bin []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{bin: []byte("000")},
			want: 0,
		},
		{
			name: "1",
			args: args{bin: []byte("001")},
			want: 1,
		},
		{
			name: "4",
			args: args{bin: []byte("100")},
			want: 4,
		},
		{
			name: "7",
			args: args{bin: []byte("111")},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryStringInBytesToInt(tt.args.bin); got != tt.want {
				t.Errorf("binaryStringInBytesToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsePacketHeader(t *testing.T) {
	type args struct {
		buf *bytes.Buffer
	}
	tests := []struct {
		name        string
		args        args
		wantVersion int
		wantTypeId  int
	}{
		{
			name:        "0 0",
			args:        args{buf: bytes.NewBufferString("000000")},
			wantVersion: 0,
			wantTypeId:  0,
		},
		{
			name:        "7 4",
			args:        args{buf: bytes.NewBufferString("111100")},
			wantVersion: 7,
			wantTypeId:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVersion, gotTypeId := parsePacketHeader(tt.args.buf)
			if gotVersion != tt.wantVersion {
				t.Errorf("parsePacketHeader() gotVersion = %v, want %v", gotVersion, tt.wantVersion)
			}
			if gotTypeId != tt.wantTypeId {
				t.Errorf("parsePacketHeader() gotTypeId = %v, want %v", gotTypeId, tt.wantTypeId)
			}
		})
	}
}

func Test_parseLiteralValue(t *testing.T) {
	type args struct {
		buf *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{buf: bytes.NewBufferString("00000")},
			want: 0,
		},
		{
			name: "multipart 0",
			args: args{buf: bytes.NewBufferString("1000000000")},
			want: 0,
		},
		{
			name: "8",
			args: args{buf: bytes.NewBufferString("01000")},
			want: 8,
		},
		{
			name: "multipart 8",
			args: args{buf: bytes.NewBufferString("1000001000")},
			want: 8,
		},
		{
			name: "light em up (255)",
			args: args{buf: bytes.NewBufferString("1111101111")},
			want: 255,
		},
		{
			name: "3 chunks",
			args: args{buf: bytes.NewBufferString("111111111101111")},
			want: 4095,
		},
		{
			name: "4 chunks",
			args: args{buf: bytes.NewBufferString("11000100001000000000")},
			want: 32768,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLiteralValue(tt.args.buf); got != tt.want {
				t.Errorf("parseLiteralValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

// One set of examples only gave version sum w/o value and the other only gave value,
// so instead of bothering to compute it myself I just split the tests

func Test_parsePacketIgnoringValue(t *testing.T) {
	type args struct {
		buf *bytes.Buffer
	}
	tests := []struct {
		name           string
		args           args
		wantVersionSum int
		wantValue      int
		wantErr        bool
	}{
		{
			name:           "short packet, error",
			args:           args{buf: bytes.NewBufferString("1111110000")},
			wantValue:      0,
			wantVersionSum: 0,
			wantErr:        true,
		},
		{
			name:           "example D2FE28, just a value",
			args:           args{buf: bytes.NewBufferString("110100101111111000101000")},
			wantValue:      2021,
			wantVersionSum: 6,
			wantErr:        false,
		},
		{
			name:           "example 38006F45291200, operator with 2 subpackets, bit-length type",
			args:           args{buf: bytes.NewBufferString("00111000000000000110111101000101001010010001001000000000")},
			wantValue:      0,
			wantVersionSum: 9,
			wantErr:        false,
		},
		{
			name:           "example EE00D40C823060, operator with 3 subpackets, packet number type",
			args:           args{buf: bytes.NewBufferString("11101110000000001101010000001100100000100011000001100000")},
			wantValue:      0,
			wantVersionSum: 14,
			wantErr:        false,
		},
		{
			name:           "example 8A004A801A8002F478",
			args:           args{buf: bytes.NewBuffer(decodeHex("8A004A801A8002F478"))},
			wantValue:      0,
			wantVersionSum: 16,
			wantErr:        false,
		},
		{
			name:           "example 620080001611562C8802118E3",
			args:           args{buf: bytes.NewBuffer(decodeHex("620080001611562C8802118E34"))},
			wantValue:      0,
			wantVersionSum: 12,
			wantErr:        false,
		},
		{
			name:           "example C0015000016115A2E0802F182340",
			args:           args{buf: bytes.NewBuffer(decodeHex("C0015000016115A2E0802F182340"))},
			wantValue:      0,
			wantVersionSum: 23,
			wantErr:        false,
		},
		{
			name:           "example A0016C880162017C3686B18A3D4780",
			args:           args{buf: bytes.NewBuffer(decodeHex("A0016C880162017C3686B18A3D4780"))},
			wantValue:      0,
			wantVersionSum: 31,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVersionSum, gotValue, err := parsePacket(tt.args.buf)
			if gotVersionSum != tt.wantVersionSum {
				t.Errorf("parsePacket() gotVersionSum = %v, want %v", gotVersionSum, tt.wantVersionSum)
			}
			if gotValue != tt.wantValue {
				//t.Errorf("parsePacket() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_parsePacketIgnoringVersion(t *testing.T) {
	type args struct {
		buf *bytes.Buffer
	}
	tests := []struct {
		name           string
		args           args
		wantVersionSum int
		wantValue      int
		wantErr        bool
	}{
		{
			name:           "example C200B40A82",
			args:           args{buf: bytes.NewBuffer(decodeHex("C200B40A82"))},
			wantValue:      3,
			wantVersionSum: 0,
			wantErr:        false,
		},
		{
			name:           "example 04005AC33890",
			args:           args{buf: bytes.NewBuffer(decodeHex("04005AC33890"))},
			wantValue:      54,
			wantVersionSum: 0,
			wantErr:        false,
		},
		{
			name:           "example 880086C3E88112",
			args:           args{buf: bytes.NewBuffer(decodeHex("880086C3E88112"))},
			wantValue:      7,
			wantVersionSum: 0,
			wantErr:        false,
		},
		{
			name:           "example CE00C43D881120",
			args:           args{buf: bytes.NewBuffer(decodeHex("CE00C43D881120"))},
			wantValue:      9,
			wantVersionSum: 0,
			wantErr:        false,
		},
		{
			name:           "example D8005AC2A8F0",
			args:           args{buf: bytes.NewBuffer(decodeHex("D8005AC2A8F0"))},
			wantValue:      1,
			wantVersionSum: 0,
			wantErr:        false,
		},
		{
			name:           "example F600BC2D8F",
			args:           args{buf: bytes.NewBuffer(decodeHex("F600BC2D8F"))},
			wantValue:      0,
			wantVersionSum: 0,
			wantErr:        false,
		},
		{
			name:           "example 9C005AC2F8F0",
			args:           args{buf: bytes.NewBuffer(decodeHex("9C005AC2F8F0"))},
			wantValue:      0,
			wantVersionSum: 0,
			wantErr:        false,
		},
		{
			name:           "example 9C0141080250320F1802104A08",
			args:           args{buf: bytes.NewBuffer(decodeHex("9C0141080250320F1802104A08"))},
			wantValue:      1,
			wantVersionSum: 0,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVersionSum, gotValue, err := parsePacket(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotVersionSum != tt.wantVersionSum {
				//t.Errorf("parsePacket() gotVersionSum = %v, want %v", gotVersionSum, tt.wantVersionSum)
			}
			if gotValue != tt.wantValue {
				t.Errorf("parsePacket() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
