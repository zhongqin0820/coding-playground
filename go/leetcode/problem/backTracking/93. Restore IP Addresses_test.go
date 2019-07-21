package problem

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/restore-ip-addresses/

// Given a string containing only digits, restore it by returning all possible valid IP address combinations.

func restoreIpAddresses(s string) []string {
	res := []string{}
	helper93(s, 4, "", &res)
	return res
}

func helper93(s string, part int, prefix string, list *[]string) {
	if part == 0 {
		if len(s) == 0 {
			//remve last .
			*list = append(*list, prefix[:len(prefix)-1])
		}
		return
	}
	if len(s) < part {
		return
	}
	// default 3
	round := 3
	// exmple, len(s) = 4, part = 3, can only cut on 1 or 2
	if len(s)-part+1 < 3 {
		round = len(s) - part + 1
	}
	for i := 1; i <= round; i++ {
		ip := s[0:i]
		if helper93isValid(ip) {
			helper93(s[i:], part-1, prefix+ip+".", list)
		}
	}
}

func helper93isValid(s string) bool {
	if s == "0" {
		return true
	}
	if strings.HasPrefix(s, "0") {
		return false
	}
	num, _ := strconv.Atoi(s)
	if num > 255 {
		return false
	}
	return true
}

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		input  string
		output []string
	}{
		{"", []string{}},
		{"26526511135", []string{}},
		{"000", []string{}},
		{"10111", []string{"1.0.1.11", "1.0.11.1", "10.1.1.1"}},
		{"111111", []string{"1.1.1.111", "1.1.11.11", "1.1.111.1", "1.11.1.11", "1.11.11.1", "1.111.1.1", "11.1.1.11", "11.1.11.1", "11.11.1.1", "111.1.1.1"}},
		{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(restoreIpAddresses(ts.input), ts.output)
		})
	}
}
