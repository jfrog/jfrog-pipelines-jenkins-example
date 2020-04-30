package main

import (
	"github.com/jefferyfry/funclog"
	"goci-example/api"
)

var (
	LogE = funclog.NewErrorLogger("ERROR: ")
	LogI = funclog.NewErrorLogger("INFO: ")
)

func main() {
	LogI.Print("goci-example started!")
	LogE.Fatal(api.StartApiService())
}