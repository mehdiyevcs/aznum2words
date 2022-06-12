package helper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (

	//single digits numbers
	ZeroAsString  string = "sıfır"  // 0
	OneAsString   string = "bir"    // 1
	TwoAsString   string = "iki"    // 2
	ThreeAsString string = "üç"     // 3
	FourAsString  string = "dörd"   // 4
	FiveAsString  string = "beş"    // 5
	SixAsString   string = "altı"   // 6
	SevenAsString string = "yeddi"  // 7
	EightAsString string = "səkkiz" // 8
	NineAsString  string = "doqquz" // 9

	//two digits numbers
	TenAsString     string = "on"     // 10
	TwentyAsString  string = "iyirmi" // 20
	ThirtyAsString  string = "otuz"   // 30
	FortyAsString   string = "qırx"   // 40
	FiftyAsString   string = "əlli"   // 50
	SixtyAsString   string = "altmış" // 60
	SeventyAsString string = "yetmiş" // 70
	EightyAsString  string = "səksən" // 80
	NinetyAsString  string = "doxsan" // 90

	//three digits numbers
	HundredAsString     string = "yüz"        // 10^2
	ThousandAsString    string = "min"        // 10^3
	MillionAsString     string = "milyon"     // 10^6
	BillionAsString     string = "milyard"    // 10^9
	TrillionAsString    string = "trilyon"    // 10^12
	QuadrillionAsString string = "katrilyon"  // 10^15
	QuintillionAsString string = "kentilyon"  // 10^18
	SextillionAsString  string = "sekstilyon" // 10^21
	SeptillionAsString  string = "septilyon"  // 10^24
	OctillionAsString   string = "oktilyon"   // 10^27
	NonillionAsString   string = "nonilyon"   // 10^30
	//TODO Add remaining ones whenever have time
)

var digits = []string{
	ZeroAsString,
	OneAsString,
	TwoAsString,
	ThreeAsString,
	FourAsString,
	FiveAsString,
	SixAsString,
	SevenAsString,
	EightAsString,
	NineAsString,
}

var tens = []string{
	"",
	TenAsString,
	TwentyAsString,
	ThirtyAsString,
	FortyAsString,
	FiftyAsString,
	SixtyAsString,
	SeventyAsString,
	EightyAsString,
	NinetyAsString,
}

// Convert integer given in str to word
func ConvertIntPart(str string) string {

	//TODO 'str' may contain invalid symbols, need to sanitize it before processing
	//it contains separator like `,`

	//TODO 'str' may contain symbols like `,` to separate integers for readability

	//starting from right hand side, pick up triples and start process it

	//used to indicate level 10^3, 10^6
	pos := 1
	wordBuilder := make([]string, 0)
	for i := len(str); i >= 0; i = i - 3 {
		key, _ := getKeyword(pos)
		pos++

		var leftPointer int = i - 3
		if leftPointer < 0 {
			leftPointer = 0
		}
		var rightPointer int = i

		tripleAsWords, _ := tripleToWord(str[leftPointer:rightPointer])

		//insert triple `tripleAsWords` in the front of result
		if len(tripleAsWords) != 0 {
			if len(key) != 0 {
				wordBuilder = append([]string{tripleAsWords, key}, wordBuilder...)
			} else {
				wordBuilder = append([]string{tripleAsWords}, wordBuilder...)
			}
		}

	}

	return strings.Join(wordBuilder, " ")
}

func getKeyword(position int) (keyword string, err error) {
	switch position {
	case 1:
		//first triple: 10^3
		return "", nil
	case 2:
		return ThousandAsString, nil
	case 3:
		return MillionAsString, nil
	case 4:
		return BillionAsString, nil
	case 5:
		return TrillionAsString, nil
	case 6:
		return QuadrillionAsString, nil
	case 7:
		return QuintillionAsString, nil
	case 8:
		return SextillionAsString, nil
	case 9:
		return SeptillionAsString, nil
	case 10:
		return OctillionAsString, nil
	case 11:
		return NonillionAsString, nil
	default:
		return "", errors.New(fmt.Sprintf("Max supported number level: %s", NonillionAsString))
	}
}

func tripleToWord(triple string) (string, error) {

	if len(triple) == 0 {
		return "", nil
	}

	if len(triple) == 1 {
		if i, err := strconv.Atoi(triple); err != nil {
			return "", err
		} else {
			return digits[i], nil
		}
	}

	if len(triple) == 2 {
		var textBuilder []string
		var idxAt1 int
		if i, err := strconv.Atoi(triple[1:]); err != nil {
			return "", err
		} else {
			idxAt1 = i
		}
		var wordForIdxAt1 string
		if idxAt1 == 0 {
			wordForIdxAt1 = ""
		} else {
			wordForIdxAt1 = digits[idxAt1]
		}

		if len(wordForIdxAt1) != 0 {
			textBuilder = append([]string{wordForIdxAt1}, textBuilder...)
		}

		var idxAt0 int
		if i, err := strconv.Atoi(triple[0:1]); err != nil {
			return "", err
		} else {
			idxAt0 = i
		}

		wordForIdxAt0 := tens[idxAt0]

		if len(wordForIdxAt0) != 0 {
			textBuilder = append([]string{wordForIdxAt0}, textBuilder...)
		}

		return strings.Join(textBuilder, " "), nil
	}

	var textBuilder []string
	for i := len(triple) - 1; i >= 0; i-- {

		var word string
		switch i {
		case 2:
			var ValAtIdx2 int
			if i, err := strconv.Atoi(triple[2:]); err != nil {
				return "", err
			} else {
				ValAtIdx2 = i
			}
			if ValAtIdx2 != 0 {
				word = digits[ValAtIdx2]
			}

		case 1:
			var ValAtIdx1 int
			if i, err := strconv.Atoi(triple[1:2]); err != nil {
				return "", err
			} else {
				ValAtIdx1 = i
			}
			if ValAtIdx1 != 0 {
				word = tens[ValAtIdx1]
			} else {
				word = ""
			}
		case 0:
			var ValAtIdx0 int
			if i, err := strconv.Atoi(triple[0:1]); err != nil {
				return "", err
			} else {
				ValAtIdx0 = i
			}
			isNonZero := ValAtIdx0 != 0

			var prefix string
			if isNonZero {
				prefix = digits[ValAtIdx0]
			} else {
				prefix = ""
			}

			if len(prefix) > 0 {
				word = prefix + " " + HundredAsString
			} else {
				word = ""
			}
		default:
			return "", errors.New("")
		}

		if len(word) != 0 {
			textBuilder = append([]string{word}, textBuilder...)
		}
	}

	return strings.Join(textBuilder, " "), nil
}