package advent4

import (
	"reflect"
	"testing"
)



func Test_newPassportValues(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *passportValues
	}{
		{
			name: "basic",
			args: args{data: "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"},
			want: &passportValues{Ecl: "gry", Pid:"860033327", Eyr:"2020", Hcl: "#fffffd", Byr:"1937", Iyr:"2017", Cid:"147", Hgt: "183cm" },
		},
		{
			name: "newlines",
			args: args{data: "ecl:gry\npid:860033327\neyr:2020\nhcl:#fffffd\nbyr:1937\niyr:2017\ncid:147\nhgt:183cm"},
			want: &passportValues{Ecl: "gry", Pid:"860033327", Eyr:"2020", Hcl: "#fffffd", Byr:"1937", Iyr:"2017", Cid:"147", Hgt: "183cm" },
		},
		{
			name: "actual",
			args: args{data: "byr:1937\neyr:2030 pid:154364481\nhgt:158cm iyr:2015 ecl:brn hcl:#c0946f cid:155"},
			want: &passportValues{Ecl: "brn", Pid:"154364481", Eyr:"2030", Hcl: "#c0946f", Byr:"1937", Iyr:"2015", Cid:"155", Hgt: "158cm" },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPassportValues(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPassportValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsePassports(t *testing.T) {
	type args struct {
		inputs string
	}
	tests := []struct {
		name string
		args args
		want []passportValues
	}{
		{
			name: "basic",
			args: args{inputs: "ecl:gry pid:860033327 pyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm eyr:2020"},
			want: []passportValues{{Ecl: "gry", Pid:"860033327", Eyr:"2020", Hcl: "#fffffd", Byr:"1937", Iyr:"2017", Cid:"147", Hgt: "183cm" }},
		},
		{
			name: "newline",
			args: args{inputs: "Ecl:gry Pid:860033327 Eyr:2020 Hcl:#fffffd\nByr:1937 Iyr:2017 Cid:147 Hgt:183cm"},
			want: []passportValues{{Ecl: "gry", Pid:"860033327", Eyr:"2020", Hcl: "#fffffd", Byr:"1937", Iyr:"2017", Cid:"147", Hgt: "183cm" }},
		},
		{
			name: "two records",
			args: args{inputs: "Ecl:gry Pid:860033327 Eyr:2020 Hcl:#fffffd\nByr:1937 Iyr:2017 Cid:147 Hgt:183cm\n\nEcl:blu Pid:8675309 Eyr:2020 Hcl:#aaaaaa Byr:1950 Iyr:2015 Hgt:190cm"},
			want: []passportValues{{Ecl: "gry", Pid: "860033327", Eyr: "2020", Hcl: "#fffffd", Byr: "1937", Iyr: "2017", Cid: "147", Hgt: "183cm"},
				{Ecl: "blu", Pid: "8675309", Eyr: "2020", Hcl: "#aaaaaa", Byr: "1950", Iyr: "2015", Cid: "", Hgt: "190cm"}},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePassports(tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePassports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateHeight(t *testing.T) {
	type args struct {
		hgt string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "good cm",
			args: args{hgt:"150cm"},
			want: true,
		},
		{
			name: "good in",
			args: args{hgt:"68in"},
			want: true,
		},
		{
			name: "bad unit",
			args: args{hgt:"150mi"},
			want: false,
		},
		{
			name: "no unit",
			args: args{hgt:"150"},
			want: false,
		},
		{
			name: "low cm",
			args: args{hgt:"15cm"},
			want: false,
		},
		{
			name: "high cm ",
			args: args{hgt:"250cm"},
			want: false,
		},
		{
			name: "low in",
			args: args{hgt:"12in"},
			want: false,
		},
		{
			name: "high in",
			args: args{hgt:"80in"},
			want: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateHeight(tt.args.hgt); got != tt.want {
				t.Errorf("validateHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passportValues_isValid(t *testing.T) {
	type fields struct {
		Byr string
		Iyr string
		Eyr string
		Hgt string
		Hcl string
		Ecl string
		Pid string
		Cid string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "basic",
			fields: fields{
				Byr: "1950",
				Iyr: "2015",
				Eyr: "2025",
				Hgt: "68in",
				Hcl: "#012345",
				Ecl: "blu",
				Pid: "123456789",
				Cid: "111",
			},
			want: true,
		},
		{
			name: "variant 1",
			fields: fields{
				Byr: "1950",
				Iyr: "2015",
				Eyr: "2025",
				Hgt: "150cm",
				Hcl: "#ffffff",
				Ecl: "grn",
				Pid: "000000000",
			},
			want: true,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pass := &passportValues{
				Byr: tt.fields.Byr,
				Iyr: tt.fields.Iyr,
				Eyr: tt.fields.Eyr,
				Hgt: tt.fields.Hgt,
				Hcl: tt.fields.Hcl,
				Ecl: tt.fields.Ecl,
				Pid: tt.fields.Pid,
				Cid: tt.fields.Cid,
			}
			if got := pass.isValid(); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
