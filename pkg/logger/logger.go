package logger

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	Silent  bool
	NoColor bool
)

func Banner() {
	banner := `
   __                  ____                       
  / /  ___ ________ __/ _  \__ ________ ___ _    
 / /__/ _ /_ / // / _/ ___/ _ / __/ _ '/  ' \   
/____/\_,_/__/\_, / /_/   \_,_/_/  \_,_/_/_/_/   
             /___/                               
      v0.1.0 | By: mahyar
      Email: mohammadmahyarkermani@gmail.com
      Website: mahyar.sbs`

	if NoColor {
		fmt.Println(banner)
	} else {
		color.Yellow(banner)
	}
	fmt.Println()
}


func Info(format string, a ...interface{}) {
	if Silent { return }
	if NoColor { fmt.Printf("[INFO] "+format+"\n", a...) } else { color.Cyan("[INFO] "+format, a...) }
}
func Success(format string, a ...interface{}) {
	if Silent { return }
	if NoColor { fmt.Printf("[SUCCESS] "+format+"\n", a...) } else { color.Green("[SUCCESS] "+format, a...) }
}
func Error(format string, a ...interface{}) {
	if Silent { return }
	if NoColor { fmt.Printf("[ERROR] "+format+"\n", a...) } else { color.Red("[ERROR] "+format, a...) }
}
