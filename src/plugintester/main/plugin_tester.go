package main

import (
	"errors"
	"log"
	"path/filepath"
	"plugin"
)

type PluginTester struct {
	PluginsPath string
}

type PluginGreeter interface {
	Greet()
}

func NewPluginTester(pluginsPath string) *PluginTester {
	return &PluginTester{
		PluginsPath: pluginsPath,
	}
}

func (pluginTester *PluginTester) Run() error {

	var err error

	if err = pluginTester.loadPlugin("plugin01"); err != nil {
		return err
	}

	if err = pluginTester.loadPlugin("plugin02"); err != nil {
		return err
	}

	return nil
}

func (pluginTester *PluginTester) loadPlugin(name string) error {

	pluginPath := filepath.Join(pluginTester.PluginsPath, name+".so")

	log.Println("Loading plugin: " + pluginPath)
	plug, err := plugin.Open(pluginPath)
	if err != nil {
		return err
	}

	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		return err
	}

	var greeter PluginGreeter
	greeter, ok := symGreeter.(PluginGreeter)
	if !ok {
		return errors.New("Unexpected type from module symbol")
	}

	greeter.Greet()

	return nil

}
