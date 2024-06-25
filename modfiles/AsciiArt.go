package modfiles

import "fmt"

// AsciiArtPrint prints ASCII art to the console.
func AsciiArtPrint() {
	asciiArt := `
  ____      ____                           
 / ___| ___/ ___|  ___  ___ _   _ _ __ ___ 
| |  _ / _ \___ \ / _ \/ __| | | | '__/ _ \
| |_| | (_) |__) |  __/ (__| |_| | | |  __/
 \____|\___/____/ \___|\___|\__,_|_|  \___|                                                        
`
	fmt.Println("\t\t", asciiArt)
}
