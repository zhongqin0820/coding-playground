package problem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.
// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

// Note:
// Although the above answer is in lexicographical order, your answer could be in any order you want.

func letterCombinations(digits string) []string {
	var m = map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	if digits == "" {
		return nil
	}

	ret := []string{""}
	for i := 0; i < len(digits); i++ { // 让digits中所有的数字都有机会进行迭代。
		temp := []string{}
		for j := 0; j < len(ret); j++ { // 让ret中的每个元素都能添加新的字母。
			for k := 0; k < len(m[digits[i]]); k++ { // 把digits[i]所对应的字母，多次添加到ret[j]的末尾
				temp = append(temp, ret[j]+m[digits[i]][k])
			}
		}
		ret = temp
	}

	return ret
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		digits string
		output []string
	}{
		{
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"34567",
			[]string{"dgjmp", "dgjmq", "dgjmr", "dgjms", "dgjnp", "dgjnq", "dgjnr", "dgjns", "dgjop", "dgjoq", "dgjor", "dgjos", "dgkmp", "dgkmq", "dgkmr", "dgkms", "dgknp", "dgknq", "dgknr", "dgkns", "dgkop", "dgkoq", "dgkor", "dgkos", "dglmp", "dglmq", "dglmr", "dglms", "dglnp", "dglnq", "dglnr", "dglns", "dglop", "dgloq", "dglor", "dglos", "dhjmp", "dhjmq", "dhjmr", "dhjms", "dhjnp", "dhjnq", "dhjnr", "dhjns", "dhjop", "dhjoq", "dhjor", "dhjos", "dhkmp", "dhkmq", "dhkmr", "dhkms", "dhknp", "dhknq", "dhknr", "dhkns", "dhkop", "dhkoq", "dhkor", "dhkos", "dhlmp", "dhlmq", "dhlmr", "dhlms", "dhlnp", "dhlnq", "dhlnr", "dhlns", "dhlop", "dhloq", "dhlor", "dhlos", "dijmp", "dijmq", "dijmr", "dijms", "dijnp", "dijnq", "dijnr", "dijns", "dijop", "dijoq", "dijor", "dijos", "dikmp", "dikmq", "dikmr", "dikms", "diknp", "diknq", "diknr", "dikns", "dikop", "dikoq", "dikor", "dikos", "dilmp", "dilmq", "dilmr", "dilms", "dilnp", "dilnq", "dilnr", "dilns", "dilop", "diloq", "dilor", "dilos", "egjmp", "egjmq", "egjmr", "egjms", "egjnp", "egjnq", "egjnr", "egjns", "egjop", "egjoq", "egjor", "egjos", "egkmp", "egkmq", "egkmr", "egkms", "egknp", "egknq", "egknr", "egkns", "egkop", "egkoq", "egkor", "egkos", "eglmp", "eglmq", "eglmr", "eglms", "eglnp", "eglnq", "eglnr", "eglns", "eglop", "egloq", "eglor", "eglos", "ehjmp", "ehjmq", "ehjmr", "ehjms", "ehjnp", "ehjnq", "ehjnr", "ehjns", "ehjop", "ehjoq", "ehjor", "ehjos", "ehkmp", "ehkmq", "ehkmr", "ehkms", "ehknp", "ehknq", "ehknr", "ehkns", "ehkop", "ehkoq", "ehkor", "ehkos", "ehlmp", "ehlmq", "ehlmr", "ehlms", "ehlnp", "ehlnq", "ehlnr", "ehlns", "ehlop", "ehloq", "ehlor", "ehlos", "eijmp", "eijmq", "eijmr", "eijms", "eijnp", "eijnq", "eijnr", "eijns", "eijop", "eijoq", "eijor", "eijos", "eikmp", "eikmq", "eikmr", "eikms", "eiknp", "eiknq", "eiknr", "eikns", "eikop", "eikoq", "eikor", "eikos", "eilmp", "eilmq", "eilmr", "eilms", "eilnp", "eilnq", "eilnr", "eilns", "eilop", "eiloq", "eilor", "eilos", "fgjmp", "fgjmq", "fgjmr", "fgjms", "fgjnp", "fgjnq", "fgjnr", "fgjns", "fgjop", "fgjoq", "fgjor", "fgjos", "fgkmp", "fgkmq", "fgkmr", "fgkms", "fgknp", "fgknq", "fgknr", "fgkns", "fgkop", "fgkoq", "fgkor", "fgkos", "fglmp", "fglmq", "fglmr", "fglms", "fglnp", "fglnq", "fglnr", "fglns", "fglop", "fgloq", "fglor", "fglos", "fhjmp", "fhjmq", "fhjmr", "fhjms", "fhjnp", "fhjnq", "fhjnr", "fhjns", "fhjop", "fhjoq", "fhjor", "fhjos", "fhkmp", "fhkmq", "fhkmr", "fhkms", "fhknp", "fhknq", "fhknr", "fhkns", "fhkop", "fhkoq", "fhkor", "fhkos", "fhlmp", "fhlmq", "fhlmr", "fhlms", "fhlnp", "fhlnq", "fhlnr", "fhlns", "fhlop", "fhloq", "fhlor", "fhlos", "fijmp", "fijmq", "fijmr", "fijms", "fijnp", "fijnq", "fijnr", "fijns", "fijop", "fijoq", "fijor", "fijos", "fikmp", "fikmq", "fikmr", "fikms", "fiknp", "fiknq", "fiknr", "fikns", "fikop", "fikoq", "fikor", "fikos", "filmp", "filmq", "filmr", "films", "filnp", "filnq", "filnr", "filns", "filop", "filoq", "filor", "filos"},
		},
	}
	for i, ts := range tests {
		t.Run(fmt.Sprintf("Example %d", i+1), func(t *testing.T) {
			ast := assert.New(t)
			ast.Equal(ts.output, letterCombinations(ts.digits))
		})
	}
}
