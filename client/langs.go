package client

// Langs mapping of programming language IDs to programming languages.
var Langs = map[string]string{
	"89": "GNU G++20 13.2 (64 bit, winlibs)",
}

// ExtToLangID mapping of file extensions to programming language IDs.
var ExtToLangID = map[string]string{
	".cpp": "89",
	".cc":  "89",
	".cxx": "89",
	".c++": "89",
}
