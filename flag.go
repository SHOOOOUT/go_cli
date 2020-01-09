package main

import (
	"flag"
	f "fmt"
)

func main() {
	/* フラグの定義 */
	strCmd := flag.String("s", "Shuto Nakano", "String help message")
	intCmd := flag.Int("i", 23, "Int help message")
	boolCmd := flag.Bool("b", false, "Bool help message")

	flag.Parse() //コマンドラインを解析し, 定義されたフラグにセット

	f.Println("Name: ", *strCmd)
	f.Println("Age: ", *intCmd)
	f.Println("Bool: ", *boolCmd)

}
