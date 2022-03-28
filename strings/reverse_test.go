package quiz

import (
	"reflect"
	"testing"
)

var testBrown string

func init() {
	testBrown = dummyString()

}
func TestMirror(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  " this is the world as it\vis\tright now ",
			output: " now right\tis\vit as world the is this ",
		},
		{
			"a",
			"a",
		},
		{
			"abc\t\v\refg",
			"efg\r\v\tabc",
		},
		{
			"",
			"",
		},
		{
			" \t \v \r \f ",
			" \f \r \v \t ",
		},
		{
			"abc   def gh",
			"gh def   abc",
		},
		{
			"   ",
			"   ",
		},
		{
			"The quick brown 狐 jumped over the lazy 犬",
			"犬 lazy the over jumped 狐 brown quick The",
		},
		{
			"The quick bròwn 狐 jumped over the lazy 犬",
			"犬 lazy the over jumped 狐 bròwn quick The",
		},
	}

	for i, tc := range tests {
		//got := reverseSentence(tc.input)
		//fmt.Println(WordPostions(tc.input, unicode.IsSpace))
		got := Mirror(tc.input)
		if got != tc.output {
			t.Errorf("%d> Mirror(%s)\ngot = %s\n want = %s", i, tc.input, got, tc.output)
			t.Errorf("%d %d", len(tc.input), len(got))
		}
	}
}

func TestReverseUTF8(t *testing.T) {
	tests := []struct {
		x string
		y string
	}{
		{
			x: "",
			y: "",
		},
		{
			x: "a",
			y: "a",
		},
		{
			x: "ab cd",
			y: "dc ba",
		},
		{
			x: "ab cde",
			y: "edc ba",
		},
		{
			x: "The quick brown 狐 jumped over the lazy 犬",
			y: "犬 yzal eht revo depmuj 狐 nworb kciuq ehT",
		},
		{
			"bròwn",
			"nwòrb",
		},

		{
			x: "The quick bròwn 狐 jumped over the lazy 犬",
			y: "犬 yzal eht revo depmuj 狐 nwòrb kciuq ehT",
		},
		{
			x: "abåd",
			y: "dåba",
		},
		{
			"abć̝d",
			"dć̝ba",
		},
	}

	for i, c := range tests {
		got := reverseUTF8Com(c.x)
		if !reflect.DeepEqual(got, c.y) {
			t.Errorf("case %d reverse(%s)\n got: %s\n want: %s", i, c.x, got, c.y)
		}
	}
}

func dummyString() string {
	input := "The quick bròwn 狐 jumped over the lazy 犬"
	cs := "狐犬ć̝狐犬ć̝狐犬ć̝狐犬ć̝"
	input += cs
	for i := 0; i < 4; i++ {
		input += input
	}
	//n := len(input)
	//fmt.Println("lenght of dummy string",n, n/8.0, n/(8.0*1024))
	//panic(n)
	return input
}

var (
	ret1 string
	ret2 string
	ret3 string
)

// func BenchmarkReverseUTF8(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		ret1 = reverseUTF8(testBrown)
// 	}
// }

func BenchmarkReverseUTF8Cox(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ret2 = reverseUTF8c(testBrown)
	}
}

// func BenchmarkReverseUTF8b(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		ret3 = reverseUTF8b(testBrown)
// 	}
// }

// func BenchmarkReverseUTF8b2(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		ret3 = reverseUTF8b2(testBrown)
// 	}
// }

func BenchmarkReverseUTF8b3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ret3 = reverseUTF8b3(testBrown)
	}
}
