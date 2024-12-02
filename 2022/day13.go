package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func tokenizeList(s string) []string {
	s = s[1 : len(s)-1]

	var list []string
	open := 0
	tokenStart := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			open++
		} else if s[i] == ']' {
			open--
		} else if open == 0 && s[i] == ',' {
			list = append(list, s[tokenStart:i])
			tokenStart = i + 1
		}
	}

	if tokenStart < len(s) {
		list = append(list, s[tokenStart:])
	}

	return list
}

func comparePackets(p1 string, p2 string) (bool, bool) {
	p1IsList := strings.HasPrefix(p1, "[")
	p2IsList := strings.HasPrefix(p2, "[")

	if !p1IsList && !p2IsList {
		p1I := util.MustInt(p1)
		p2I := util.MustInt(p2)

		return p1I == p2I, p1I < p2I
	}

	if !p2IsList {
		return comparePackets(p1, "["+p2+"]")
	} else if !p1IsList {
		return comparePackets("["+p1+"]", p2)
	}

	p1List := tokenizeList(p1)
	p2List := tokenizeList(p2)
	p1ListLen := len(p1List)
	p2ListLen := len(p2List)

	bound := min(p1ListLen, p2ListLen)

	for i := 0; i < bound; i++ {
		if equal, order := comparePackets(p1List[i], p2List[i]); !equal {
			return false, order
		}
	}

	return p1ListLen == p2ListLen, p1ListLen < p2ListLen
}

type packetList []string

func (p packetList) Len() int {
	return len(p)
}

func (p packetList) Less(i, j int) bool {
	equal, inOrder := comparePackets(p[i], p[j])
	return !equal && inOrder
}

func (p packetList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Day13() error {
	f, err := os.Open(filepath.Join(util.Cwd(), "day13.txt"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	pairIndex := 1
	var indexSum int
	var packets packetList
	var packetPair []string

	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			continue
		}

		packetPair = append(packetPair, t)

		if len(packetPair) == 2 {
			if equal, inOrder := comparePackets(packetPair[0], packetPair[1]); !equal && inOrder {
				indexSum += pairIndex
			}

			packets = append(packets, packetPair...)
			packetPair = []string{}
			pairIndex++
		}
	}

	fmt.Println(indexSum)

	packets = append(packets, "[[2]]", "[[6]]")

	sort.Sort(packets)

	var decoderKey = 1

	for i, p := range packets {
		if p == "[[2]]" || p == "[[6]]" {
			decoderKey *= i + 1
		}
	}

	fmt.Println(decoderKey)

	return nil
}
