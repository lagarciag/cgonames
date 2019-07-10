package cgonames

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"github.com/etgryphon/stringUp"


)

//ConvertCformat2GoFormatName is self explanatory
func ConvertCformat2GoFormatName(in string) (out string) {

	inSlice := strings.Split(in, "_")
	for _, word := range inSlice {
		newWord := strings.Title(word)

		if len(word) == 3 && strings.ContainsAny(string(word[1]), "0123456789") {
			out = out + strings.ToUpper(word)
		} else if word == "v6e" {
			out = out + "V6E"
		} else if word == "l2a" {
			out = out + "L2A"
		} else if word == "l2b" {
			out = out + "L2B"
		} else if word == "l3a" {
			out = out + "L3A"
		} else if word == "l3b" {
			out = out + "L3B"
		} else if word == "128b" {
			out = out + "128B"
		} else {
			out = out + newWord
		}
	}
	return out
}

//ConvertCformat2GoFormatName is self explanatory
func ConvertCformat2GoFormatNameAndPipelineID(in string) (out string) {
	inSlicePre := strings.Split(in, "_")

	hasPid := true
	pipelineID, err := strconv.Atoi(inSlicePre[len(inSlicePre)-1])

	if err != nil {
		hasPid = false
	}

	preIDExists := false

	if len(inSlicePre) >= 2 {
		_, err = strconv.Atoi(inSlicePre[len(inSlicePre)-2])

		if err == nil {
			preIDExists = true
		}

	}

	inSlice := inSlicePre
	if hasPid {
		inSlice = inSlicePre[0 : len(inSlicePre)-1]
	}

	for _, word := range inSlice {
		newWord := strings.Title(word)
		if word == "v6e" {
			out = out + "V6E"
		} else if word == "l2a" {
			out = out + "L2A"
		} else if word == "l2b" {
			out = out + "L2B"
		} else if word == "l3a" {
			out = out + "L3A"
		} else if word == "l3b" {
			out = out + "L3B"
		} else if word == "128b" {
			out = out + "128B"
		} else {
			out = out + newWord
		}
	}

	if preIDExists && hasPid {
		out = fmt.Sprintf("%sY%d", out, pipelineID)
	} else if !preIDExists && hasPid {
		out = fmt.Sprintf("%s%d", out, pipelineID)
	}

	return out
}

//
func ConvertWidthToCtype(width int, name string) (aType string, bsize int, bool bool) {
	if width > 64 {
		size := width / 8
		mod := width % 8
		if mod != 0 {
			size++
		}
		aType := "char"
		return aType, size, true

	} else if width > 32 {
		return "uint64_t", 0, false
	} else if width > 16 {
		return "uint32_t", 0, false
	} else if width > 8 {
		return "uint16_t", 0, false
	} else {
		return "uint8_t", 0, false
	}
	return "", 0, false
}

func ConvertWidthToGotype(width int, name string) (aType string, bool bool) {
	if width > 64 {
		return "[]byte", true
	} else if width > 32 {
		return "uint64", false
	} else if width > 16 {
		return "uint32", false
	} else if width > 8 {
		return "uint16", false
	} else {
		return "uint8", false
	}
	return aType, false
}

func StringToLittleEndianBytes(in string) []byte {
	var x big.Int
	x.SetString(in, 10)
	myBytes := x.Bytes()
	//reverse bytes to little endian
	for i, j := 0, len(myBytes)-1; i < j; i, j = i+1, j-1 {
		myBytes[i], myBytes[j] = myBytes[j], myBytes[i]
	}
	return myBytes
}

func CNameToGoCamelCase(name string) string {
	lowName := strings.ToLower(name)
	camelName := stringUp.CamelCase(lowName)
	titledCamel := strings.Title(camelName)
	if titledCamel == "Type" {
		titledCamel = "AType"
	}
	return titledCamel

}

func CamelCaseEnum(enum string) string {

	enum = strings.TrimSuffix(enum, "_enum")
	lowEnum := strings.ToLower(enum)
	camelEnum := stringUp.CamelCase(lowEnum)
	titledCamelEnum := strings.Title(camelEnum)
	if titledCamelEnum == "Type" {
		titledCamelEnum = "AType"
	}
	return titledCamelEnum

}

