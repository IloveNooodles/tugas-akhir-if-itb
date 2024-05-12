package main

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/cmd"
	_ "github.com/lib/pq"
)

//TODO get active deployments pada device
//TODO banyak lah anjing
//TODO 

func main() {
	cli := cmd.New()
	_ = cli.Execute()
}
