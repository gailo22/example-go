package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {

	str := "aHR0cHM6Ly90ZC1hbGFybS1hcHNlMi5vYnMuYXAtc291dGhlYXN0LTIubXlodWF3ZWljbG91ZC5jb20vMjAyNDEwMTkvMzk4NTlmMjY2OTVjOGJmM2ZmZmNmOWVmMTYyYmUzMjYvMjAyNDEwMTAuaS4xNzI4NTU0NzEyLjk0NmFjNWMzLWMwMTUtNGFjYy03MmIwLWU4ODA1YjAzOTNhNC5qcGc_ZXhwaXJlcz0xNzI5MTU5NTEyJnN0b3I9b2Jz"
	// str := "aHR0cHM6Ly90ZC1hbGFybS1hcHNlNy5vc3MtYXAtc291dGhlYXN0LTcuYWxpeXVuY3MuY29tLzIwMjQxMDIwL2M4MGNhMGIzNzc1MWY0MTc3MGQ0OTc1MWE1MGUzODgxLzIwMjQxMDExLmkuMTcyODYxMTQxNC4wYTY1ODllZC0wZjgzLTQ4NTktYzdkMi03MTgxNmQ1MzZkNjMuanBnP2V4cGlyZXM9MTcyOTIxNjIxNCZzdG9yPW9zcw"

	// Check if string starts with "https://"
	if !strings.HasPrefix(str, "https://") {
		// Decode the base64 string using StdEncoding (handles padding automatically)
		d, err := base64.RawStdEncoding.DecodeString(str)
		if err != nil {
			d, _ = base64.URLEncoding.DecodeString(str)
		}
		fmt.Println(string(d))
	} else {
		// If the string is already a URL, just print it
		fmt.Println(str)
	}

}
