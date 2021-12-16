package main

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"strconv"
)

//go:embed input.txt
var inputText string

const ZERO = byte(48)
const ONE = byte(49)

func main() {
	fmt.Printf("Magic number: %d\n", run1(inputText))
	fmt.Println("-------------")
	fmt.Printf("Magic number: %d\n", run2(inputText))
}

func run1(inputText string) int {

	bits := decodeHex(inputText)
	buf := bytes.NewBuffer(bits)
	versionSum, _, _ := parsePacket(buf)

	return versionSum
}

func run2(inputText string) int {
	bits := decodeHex(inputText)
	buf := bytes.NewBuffer(bits)
	_, value, _ := parsePacket(buf)
	return value
}

//consumes 6 bits
func parsePacketHeader(buf *bytes.Buffer) (version, typeId int) {
	version = binaryStringInBytesToInt(buf.Next(3))
	typeId = binaryStringInBytesToInt(buf.Next(3))
	return
}

func parsePacket(buf *bytes.Buffer) (versionSum, value int, err error) {
	if buf.Len() < 11 {
		return 0, 0, errors.New("not enough bytes for another packet")
	}

	version, typeId := parsePacketHeader(buf)
	versionSum = version
	if typeId == 4 {
		value = parseLiteralValue(buf)
	} else {
		// operator packet - next 16 or 12 bits indicate subpackets
		ver, val := parseOperatorPacketContents(typeId, buf)
		versionSum += ver
		value = val
	}
	return versionSum, value, nil
}

func parseOperatorPacketContents(typeId int, buf *bytes.Buffer) (versionSum, value int) {
	subpacketIndicator, _ := buf.ReadByte()
	var values []int
	if subpacketIndicator == ZERO {
		// the next 15 bits are a number, NOT using the literal system
		subpacketLength := binaryStringInBytesToInt(buf.Next(15))
		// it's a number of bits so parse packets until we're out of bits
		subpacketBytes := bytes.NewBuffer(buf.Next(subpacketLength))
		for {
			ver, val, err := parsePacket(subpacketBytes)
			if err != nil {
				break
			}
			versionSum += ver
			values = append(values, val)
		}
	} else {
		// the next 11 bits are the number of packets
		subpacketLength := binaryStringInBytesToInt(buf.Next(11))
		// iterate, call parsePacket N times
		// it'll recurse if any subpacket contains subpackets
		for i := 0; i < subpacketLength; i++ {
			ver, val, err := parsePacket(buf)
			if err != nil {
				panic(fmt.Sprintf("expected %d subpackets but could only read %d", subpacketLength, i))
			}
			versionSum += ver
			values = append(values, val)
		}
	}
	// now we have the values of the subpackets, do the operation
	if len(values) == 0 {
		log.Println("ERROR: operator with no subpackets, how did that happen")
		return versionSum, 0
	}
	switch typeId {
	case 0: // sum values
		for _, v := range values {
			value += v
		}
	case 1: // product
		value = 1
		for _, v := range values {
			value *= v
		}
	case 2: // min
		value, _ = minMax(values)
	case 3: // max
		_, value = minMax(values)
	case 4: // should never see a value packet here
		panic("value packet in operator packet code- how did that happen")
	case 5: // >
		if values[0] > values[1] {
			value = 1
		}
	case 6: // <
		if values[0] < values[1] {
			value = 1
		}
	case 7: // ==
		if values[0] == values[1] {
			value = 1
		}
	}
	return versionSum, value
}

func parseLiteralValue(buf *bytes.Buffer) int {
	// chunks of 5 bytes PNNNN
	// P=1 means there are more chunks, 0= last chunk
	var num []byte
	for p, _ := buf.ReadByte(); p == ONE; p, _ = buf.ReadByte() {
		num = append(num, buf.Next(4)...)
	}
	// and finally pull the last bit that starts with 0
	num = append(num, buf.Next(4)...)
	return binaryStringInBytesToInt(num)
}

func binaryStringInBytesToInt(bin []byte) int {
	i, err := strconv.ParseInt(string(bin), 2, 0)
	if err != nil {
		panic("broke parsing bytes to int" + string(bin))
	}

	return int(i)
}

func decodeHex(inputText string) []byte {
	var b bytes.Buffer
	for _, r := range inputText {
		// r is a hex number, convert it to a nybble
		bits, _ := hexToBin(string(r))
		b.WriteString(bits)
	}

	return b.Bytes()
}

func hexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 0)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%04b", ui), nil
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
