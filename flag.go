package main

import (
	"flag"
	f "fmt"
	"os"
)

func main() {

	flag.Usage = func() {
		usageTxt := `
		Usage example [option]
		An example of customizing usage output

		-s --s STRING argument, default: String help message
		-i --i INTEGER argument, default: Int help message
		-b --b BOOLEAN argument, default: Bool help message
		`

		f.Fprintf(os.Stderr, "%s\n", usageTxt)
	}

	/* フラグの定義 */
	strCmd := flag.String("s", "Shuto Nakano", "String help message")
	intCmd := flag.Int("i", 23, "Int help message")
	boolCmd := flag.Bool("b", false, "Bool help message")

	flag.Parse() //コマンドラインを解析し, 定義されたフラグにセット

	f.Println("Name: ", *strCmd)
	f.Println("Age: ", *intCmd)
	f.Println("Bool: ", *boolCmd)

}
