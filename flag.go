package main

import (
	"flag"
	f "fmt"
	"os"
	"strconv"
)

func flagUsage() {
	usageTxt := `example is an example cli tool.

	Usage:
	example command [arguments]
	The commands are:
	convhex    convert Number to Hex
	convbinary convert Number to binary
	Use "Example [command] --help" for more infomation about a command`

	f.Fprintf(os.Stderr, "%s\n\n", usageTxt)
}

func main() {

	flag.Usage = flagUsage
	convHexCmd := flag.NewFlagSet("convhex", flag.ExitOnError)
	convBinaryCmd := flag.NewFlagSet("convbinary", flag.ExitOnError)

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "convhex":
		i := convHexCmd.Int64("i", 0, "Convert number to hex")
		convHexCmd.Parse(os.Args[2:])
		f.Println(strconv.FormatInt(*i, 16))
	case "convbinary":
		i := convBinaryCmd.Uint64("i", 0, "convert number to binary")
		convBinaryCmd.Parse(os.Args[2:])
		f.Println(strconv.FormatUint(*i, 2))
	default:
		flag.Usage()
	}

}
