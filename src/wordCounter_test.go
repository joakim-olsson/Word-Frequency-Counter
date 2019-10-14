package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

// Test that the word count on loremipsum.txt is correct
func TestWordCountOnLoremIpsum(t *testing.T) {
	data, _ := ioutil.ReadFile(DataFile)

	expected := map[string]int{"a": 466, "ac": 612, "accumsan": 201,
		"adipiscing": 479, "aenean": 237, "aliqua": 1, "aliquam": 574,
		"aliquet": 462, "amet": 1111, "ante": 76, "arcu": 584, "at": 786,
		"auctor": 214, "augue": 264, "bibendum": 305, "blandit": 286, "commodo": 336,
		"condimentum": 199, "congue": 183, "consectetur": 324, "consequat": 278,
		"convallis": 225, "cras": 358, "cum": 52, "curabitur": 85, "cursus": 439,
		"dapibus": 41, "diam": 662, "dictum": 228, "dictumst": 95, "dignissim": 293,
		"dis": 64, "do": 1, "dolor": 332, "dolore": 1, "donec": 329,
		"dui": 321, "duis": 282, "egestas": 629, "eget": 893, "eiusmod": 1,
		"eleifend": 148, "elementum": 427, "elit": 398, "enim": 761, "erat": 187, "eros": 108,
		"est": 401, "et": 804, "etiam": 221, "eu": 624, "euismod": 228,
		"facilisi": 175, "facilisis": 283, "fames": 137, "faucibus": 541, "felis": 266,
		"fermentum": 296, "feugiat": 343, "fringilla": 197, "fusce": 86,
		"gravida": 362, "habitant": 122, "habitasse": 103, "hac": 100, "hendrerit": 124,
		"iaculis": 174, "id": 950, "imperdiet": 226, "in": 1145, "incididunt": 1,
		"integer": 252, "interdum": 242, "ipsum": 364, "justo": 188, "labore": 1,
		"lacinia": 91, "lacus": 392, "laoreet": 177, "lectus": 442, "leo": 351,
		"libero": 229, "ligula": 36, "lobortis": 194, "lorem": 292, "luctus": 117,
		"maecenas": 257, "magna": 283, "magnis": 61, "malesuada": 302, "massa": 572,
		"mattis": 398, "mauris": 597, "metus": 135, "mi": 386, "molestie": 187,
		"mollis": 104, "montes": 54, "morbi": 492, "mus": 50, "nam": 125,
		"nascetur": 50, "natoque": 51, "nec": 364, "neque": 449, "netus": 130,
		"nibh": 474, "nisi": 288, "nisl": 456, "non": 585, "nulla": 558,
		"nullam": 205, "nunc": 818, "odio": 421, "orci": 371, "ornare": 338,
		"parturient": 60, "pellentesque": 704, "penatibus": 56, "pharetra": 399,
		"phasellus": 147, "placerat": 193, "platea": 105, "porta": 153,
		"porttitor": 232, "posuere": 242, "potenti": 45, "praesent": 141,
		"pretium": 310, "proin": 270, "pulvinar": 312, "purus": 465, "quam": 446,
		"quis": 646, "quisque": 195, "rhoncus": 258, "ridiculus": 50,
		"risus": 566, "rutrum": 114, "sagittis": 328, "sapien": 211,
		"scelerisque": 440, "sed": 1364, "sem": 257, "semper": 281, "senectus": 133,
		"sit": 1106, "sociis": 50, "sodales": 123, "sollicitudin": 182,
		"suscipit": 97, "suspendisse": 277, "tellus": 614, "tempor": 222,
		"tempus": 211, "tincidunt": 540, "tortor": 510,
		"tristique": 338, "turpis": 505, "ullamcorper": 335, "ultrices": 419,
		"ultricies": 261, "urna": 472, "ut": 923, "varius": 241, "vehicula": 43,
		"vel": 498, "velit": 378, "venenatis": 250, "vestibulum": 294, "vitae": 817,
		"vivamus": 92, "viverra": 646, "volutpat": 474, "vulputate": 343}

	if actual := WordCount(string(data)); !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected: %v\nactual: %v", expected, actual)
	}

}
