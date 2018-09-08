package main

import (
	"fmt"
	"shared/parameters"
	"shared/conf"
	"executionenvironment/executionenvironment"
	"os"
	"shared/shared"
)

func main(){

	// Read OS arguments
	shared.ProcessOSArguments(os.Args[1:])

	// start configuration
	EE := executionenvironment.ExecutionEnvironment{}
	EE.Exec(conf.GenerateConf(parameters.DIR_CONF + "/MiddlewareNamingServer.conf"),parameters.IS_ADAPTIVE)

	fmt.Println("Naming service started!!")
	fmt.Scanln()
}
