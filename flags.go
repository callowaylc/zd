package zd

import (
	"flag"
	"os"
	"fmt"
)

type Flags struct{

}

func InitFlags() {
	listcmd := flag.NewFlagSet("list", flag.ExitOnError)
	listcmd.Bool("recent", true, "list recently added provideres")

	showcmd := flag.NewFlagSet("show", flag.ExitOnError)
	showcmd.Int("id", 0, "show provider details with given id")

	if len(os.Args) == 1 {
		fmt.Println("usage: provider <command> [<args>]")
		fmt.Println("The most commonly used commands are:")
		fmt.Println("  list  List providers")
		fmt.Println("  show  Show provider details")

		os.Exit(0)
	}

	switch os.Args[1] {
	case "list":
		listcmd.Parse(os.Args[2:])
	case "show":
		showcmd.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
