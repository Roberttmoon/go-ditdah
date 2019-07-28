package code

func InternationalMorse() *CharSet {
	cs := &CharSet{
		Chars: map[string]*Char{
			"a": {
				ASCII:           "a",
				MorseCodeString: "._",
				MorseCode: []MorseChar{
					DitChar, DahChar,
				},
			},
			"b": {
				ASCII:           "b",
				MorseCodeString: "-...",
				MorseCode: []MorseChar{
					DahChar, DitChar, DitChar, DitChar,
				},
			},
			"c": {
				ASCII:           "c",
				MorseCodeString: "-.-.",
				MorseCode: []MorseChar{
					DahChar, DitChar, DahChar, DitChar,
				},
			},
			"d": {
				ASCII:           "d",
				MorseCodeString: "-..",
				MorseCode: []MorseChar{
					DahChar, DitChar, DitChar,
				},
			},
			"e": {
				ASCII:           "e",
				MorseCodeString: ".",
				MorseCode: []MorseChar{
					DitChar,
				},
			},
			"f": {
				ASCII:           "f",
				MorseCodeString: "..-.",
				MorseCode: []MorseChar{
					DitChar, DitChar, DahChar, DitChar,
				},
			},
			"g": {
				ASCII:           "g",
				MorseCodeString: "--.",
				MorseCode: []MorseChar{
					DahChar, DahChar, DitChar,
				},
			},
			"h": {
				ASCII:           "h",
				MorseCodeString: "....",
				MorseCode: []MorseChar{
					DitChar, DitChar, DitChar, DitChar,
				},
			},
			"i": {
				ASCII:           "i",
				MorseCodeString: "..",
				MorseCode: []MorseChar{
					DitChar, DitChar,
				},
			},
			"j": {
				ASCII:           "j",
				MorseCodeString: ".---",
				MorseCode: []MorseChar{
					DitChar, DahChar, DahChar, DahChar,
				},
			},
			"k": {
				ASCII:           "k",
				MorseCodeString: "-.-",
				MorseCode: []MorseChar{
					DahChar, DitChar, DahChar,
				},
			},
			"l": {
				ASCII:           "l",
				MorseCodeString: ".-..",
				MorseCode: []MorseChar{
					DitChar, DahChar, DitChar, DitChar,
				},
			},
			"m": {
				ASCII:           "m",
				MorseCodeString: "--",
				MorseCode: []MorseChar{
					DahChar, DahChar,
				},
			},
			"n": {
				ASCII:           "n",
				MorseCodeString: "-.",
				MorseCode: []MorseChar{
					DahChar, DitChar,
				},
			},
			"o": {
				ASCII:           "o",
				MorseCodeString: "---",
				MorseCode: []MorseChar{
					DahChar, DahChar, DahChar,
				},
			},
			"p": {
				ASCII:           "p",
				MorseCodeString: ".--.",
				MorseCode: []MorseChar{
					DitChar, DahChar, DahChar, DitChar,
				},
			},
			"q": {
				ASCII:           "q",
				MorseCodeString: "--.-",
				MorseCode: []MorseChar{
					DahChar, DahChar, DitChar, DahChar,
				},
			},
			"r": {
				ASCII:           "r",
				MorseCodeString: ".-.",
				MorseCode: []MorseChar{
					DitChar, DahChar, DitChar,
				},
			},
			"s": {
				ASCII:           "s",
				MorseCodeString: "...",
				MorseCode: []MorseChar{
					DitChar, DitChar, DitChar,
				},
			},
			"t": {
				ASCII:           "t",
				MorseCodeString: "-",
				MorseCode: []MorseChar{
					DahChar,
				},
			},
			"u": {
				ASCII:           "u",
				MorseCodeString: "..-",
				MorseCode: []MorseChar{
					DitChar, DitChar, DahChar,
				},
			},
			"v": {
				ASCII:           "v",
				MorseCodeString: "...-",
				MorseCode: []MorseChar{
					DitChar, DitChar, DitChar, DahChar,
				},
			},
			"w": {
				ASCII:           "w",
				MorseCodeString: ".--",
				MorseCode: []MorseChar{
					DitChar, DahChar, DahChar,
				},
			},
			"x": {
				ASCII:           "x",
				MorseCodeString: "-..-",
				MorseCode: []MorseChar{
					DahChar, DitChar, DitChar, DahChar,
				},
			},
			"y": {
				ASCII:           "y",
				MorseCodeString: "-.--",
				MorseCode: []MorseChar{
					DahChar, DitChar, DahChar, DahChar,
				},
			},
			"z": {
				ASCII:           "z",
				MorseCodeString: "--..",
				MorseCode: []MorseChar{
					DahChar, DahChar, DitChar, DitChar,
				},
			},
			"1": {
				ASCII:           "1",
				MorseCodeString: ".----",
				MorseCode: []MorseChar{
					DitChar, DahChar, DahChar, DahChar, DahChar,
				},
			},
			"2": {
				ASCII:           "2",
				MorseCodeString: "..---",
				MorseCode: []MorseChar{
					DitChar, DitChar, DahChar, DahChar, DahChar,
				},
			},
			"3": {
				ASCII:           "3",
				MorseCodeString: "...--",
				MorseCode: []MorseChar{
					DitChar, DitChar, DitChar, DahChar, DahChar,
				},
			},
			"4": {
				ASCII:           "4",
				MorseCodeString: "....-",
				MorseCode: []MorseChar{
					DitChar, DitChar, DitChar, DitChar, DahChar,
				},
			},
			"5": {
				ASCII:           "5",
				MorseCodeString: ".....",
				MorseCode: []MorseChar{
					DitChar, DitChar, DitChar, DitChar, DitChar,
				},
			},
			"6": {
				ASCII:           "6",
				MorseCodeString: "-....",
				MorseCode: []MorseChar{
					DahChar, DitChar, DitChar, DitChar, DitChar,
				},
			},
			"7": {
				ASCII:           "7",
				MorseCodeString: "--...",
				MorseCode: []MorseChar{
					DahChar, DahChar, DitChar, DitChar, DitChar,
				},
			},
			"8": {
				ASCII:           "8",
				MorseCodeString: "---..",
				MorseCode: []MorseChar{
					DahChar, DahChar, DahChar, DitChar, DitChar,
				},
			},
			"9": {
				ASCII:           "9",
				MorseCodeString: "----.",
				MorseCode: []MorseChar{
					DahChar, DahChar, DahChar, DahChar, DitChar,
				},
			},
			"0": {
				ASCII:           "0",
				MorseCodeString: "-----",
				MorseCode: []MorseChar{
					DahChar, DahChar, DahChar, DahChar, DahChar,
				},
			},
			".": {
				ASCII:           ".",
				MorseCodeString: ".-.-.-",
				MorseCode: []MorseChar{
					DitChar, DahChar, DitChar, DahChar, DitChar, DahChar,
				},
			},
			",": {
				ASCII:           ",",
				MorseCodeString: "--..--",
				MorseCode: []MorseChar{
					DahChar, DahChar, DitChar, DitChar, DahChar, DahChar,
				},
			},
			"?": {
				ASCII:           "?",
				MorseCodeString: "..--..",
				MorseCode: []MorseChar{
					DitChar, DitChar, DahChar, DahChar, DitChar, DitChar,
				},
			},
			"!": {
				ASCII:           "!",
				MorseCodeString: "-.-.--",
				MorseCode: []MorseChar{
					DahChar, DitChar, DahChar, DitChar, DahChar, DahChar,
				},
			},
			" ": {
				ASCII:           " ",
				MorseCodeString: " ",
				MorseCode: []MorseChar{
					SpaceChar,
				},
			},
		},
	}

	return cs
}
