package client

/*
Langs mapping of programming language IDs to programming languages.

Easy way to get languages list from Codeforces:
1. Open https://codeforces.com/problemset/customtest
2. Open Browser Developer Tool.
3. Paste this script into Console and Run.

let options = document.querySelectorAll("#pageContent > form > table > tbody > tr > td:nth-child(2) > div:nth-child(1) > select > option");
let res = "";

	for(let i = 0; i < options.length; ++i){
	    res += `"${options[i].value}": "${options[i].innerText}",\n`;
	}

console.log(res);
*/
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
