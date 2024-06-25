package asciiart

import "fmt"

// AsciiArtPrint prints ASCII art to the console.
func AsciiArtPrint() {
	asciiArt := `
                                                        
                                                        
  .g8"""bgd             .M"""bgd                        
.dP'     `M            ,MI    "Y                        
dM'       `   ,pW"Wq.  `MMb.      .gP"Ya   ,p6"bo `7MM  `7MM  `7Mb,od8 .gP"Ya   
MM           6W'   `Wb   `YMMNq. ,M'   Yb 6M'  OO   MM    MM    MM' "' ,M'   Yb 
MM.    `7MMF'8M     M8 .     `MM 8M"""""" 8M        MM    MM    MM    8M"""""" 
`Mb.     MM  YA.   ,A9 Mb     dM YM.    , YM.    ,  MM    MM    MM    YM.    , 
  `"bmmmdPY   `Ybmd9'  P"Ybmmd"   `Mbmmd'  YMbmd'   `Mbod"YML..JMML.   `Mbmmd' 
`
	fmt.Println(asciiArt)
}
