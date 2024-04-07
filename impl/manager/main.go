package main

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/cmd"
	_ "github.com/lib/pq"
)

func main() {
	cli := cmd.New()
	_ = cli.Execute()
}
