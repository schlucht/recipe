package fhx

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

func ReadBlock(startString string, lines []string) ([][]string, error) {

	results := [][]string{}
	if strings.Trim(startString, "") == "" {
		return results, errors.New("kein suchparameter vorhanden")
	}
	if len(lines) == 0 {
		return results, errors.New("kein text übergeben")
	}

	regParam, _ := regexp.Compile(startString)

	var start = false
	var curlybreak = 0
	var blockLines = []string{}

	for _, l := range lines {

		if start {
			blockLines = append(blockLines, l)
			if strings.Contains(l, "{") {
				sl := strings.Trim(l, " ")
				if len(sl) <= 2 {
					curlybreak++
				}
			}
			if strings.Contains(l, "}") {
				sl := strings.Trim(l, " ")
				if len(sl) <= 2 {
					curlybreak--
					if curlybreak == 0 {
						start = false
						results = append(results, blockLines)
						blockLines = []string{}
					}
				}
			}

		} else {
			start = regParam.MatchString(l)
			if start {
				blockLines = append(blockLines, l)
			}
		}
	}
	return results, nil
}

func ReadRegex(regex map[string]string, txt []string) map[string]interface{} {
	res := make(map[string]interface{})
	for _, l := range txt {
		for key, r := range regex {
			rCompile := regexp.MustCompile(r)
			matches := rCompile.FindStringSubmatch(l)

			if len(matches) > 0 {
				if rCompile.SubexpIndex("s") > -1 {
					res[key] = matches[rCompile.SubexpIndex("s")]
				}
				if rCompile.SubexpIndex("i") > -1 {
					i, err := strconv.ParseInt(matches[rCompile.SubexpIndex("i")], 10, 32)
					if err != nil {
						log.Printf("%v\n", err)
					}
					res[key] = i
				}
				if rCompile.SubexpIndex("b") > -1 {
					b := matches[rCompile.SubexpIndex("b")]
					res[key] = b == "T"
				}
			}
		}
	}
	return res
}

// Einlesen der FHX Datei
func ReadUTF16(file string) ([]string, error) {
	res := []string{}
	if IsFhxFile(file) {
		return res, errors.New("keine fhx datei")
	}

	f, err := os.Open(file)
	var str = ""
	if err != nil {
		return res, err
	}
	defer f.Close()

	buf := make([]byte, 1024)

	i := 0
	for {
		n, err := f.Read(buf)
		i++
		if err == io.EOF {
			break
		}
		if err != nil {
			return res, err
		}
		if n > 0 {
			if !utf8.Valid(buf[:n]) {
				s, err := DecodeUtf16(buf[:n], binary.LittleEndian)
				if err != nil {
					return res, err
				}
				str = str + s
			} else {
				str = str + string(buf[:n])
			}
		}
	}
	return splitFhx(str), nil
}

func ReadClassLines(lines []string, start string, end string) []string {
	l := []string{}
	b := false
	for _, v := range lines {
		if strings.Contains(v, start) {
			b = true
			l = append(l, v)
			continue
		}
		if strings.Contains(v, end) {
			return l
		}
		if b {
			l = append(l, v)
		}
	}
	return l
}

func DecodeUtf16(b []byte, order binary.ByteOrder) (string, error) {
	ints := make([]uint16, len(b)/2)
	if err := binary.Read(bytes.NewReader(b), order, &ints); err != nil {
		return "", err
	}
	return string(utf16.Decode(ints)), nil
}

func IsFhxFile(pathStr string) bool {
	ext := path.Ext(pathStr)
	return strings.ToUpper(ext) != ".FHX"
}

func splitFhx(txt string) []string {
	return SplitLines(txt)
}

func SplitLines(text string) []string {
	return strings.Split(text, "\n")
}
