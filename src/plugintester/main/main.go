package main

import (
	"log"
	"os"
	"path/filepath"
	// "log"
	// "plugin"
)

func main() {

	exePath, err := os.Executable()
	checkError(err)

	pluginPath := filepath.Join(filepath.Dir(exePath), "plugins")
	log.Println("Using path: " + pluginPath)

	app := NewPluginTester(pluginPath)
	err = app.Run()
	checkError(err)
	// plug, err := plugin.Open("plugin01.so")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}
