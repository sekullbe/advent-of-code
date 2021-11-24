package advent4

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/oleiade/reflections"
)

type passportValues struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func newPassportValues(data string) *passportValues {
	p := new(passportValues)
	for _, s := range strings.Fields(data) {
		chunks := strings.Split(s, ":")
		reflections.SetField(p, strings.Title(chunks[0]), chunks[1])
	}

	return p
}

func (pass *passportValues) isValid0() bool {
	return pass.Byr != "" && pass.Iyr != "" && pass.Eyr != "" && pass.Hgt != "" && pass.Hcl != "" && pass.Ecl != "" && pass.Pid != ""
}
func (pass *passportValues) isValid() bool {
	byr, err := strconv.Atoi(pass.Byr)
	if err != nil || byr < 1920 || byr > 2002 {
		return false
	}

	iyr, err := strconv.Atoi(pass.Iyr)
	if err != nil || iyr < 2010|| iyr > 2020 {
		return false
	}

	eyr, err := strconv.Atoi(pass.Eyr)
	if err != nil || eyr < 2020|| eyr > 2030 {
		return false
	}

	if !validateHeight(pass.Hgt) {
		return false
	}

	hclok, err := regexp.MatchString(`^#[a-f0-9]{6}$`, pass.Hcl)
	if err != nil || !hclok {
		return false
	}

	eclok, err := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, pass.Ecl)
	if err != nil || !eclok {
		return false
	}

	pidok, err := regexp.MatchString(`^\d{9}$`, pass.Pid)
	if err != nil || !pidok {
		return false
	}
	return true
}

func validateHeight(hgt string) bool {
	// from https://github.com/StefanSchroeder/Golang-Regex-Tutorial/blob/master/01-chapter2.markdown
	// I should probably try to understand this better
	re := regexp.MustCompile(`^(?P<value>\d+)(?P<unit>cm|in)$`)
	n1 := re.SubexpNames()
	// if it doesn't work, get out early
	if !re.MatchString(hgt) {
		return false
	}
	r2 := re.FindAllStringSubmatch(hgt, -1)[0]
	md := map[string]string{}
	for i, n := range r2 {
		//fmt.Printf("%d. match='%s'\tname='%s'\n", i, n, n1[i])
		md[n1[i]] = n
	}

	unit := md["unit"]
	if unit!= "cm" && unit != "in" {
		return false
	}
	val, err := strconv.Atoi(md["value"])
	if err != nil {
		return false
	}
	switch unit {
	case "in":
		return val >= 59 && val <= 76
	case "cm":
		return val >= 150 && val <= 193
	}
	return false
}
