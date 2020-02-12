package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

var (
	version string = "dev"
	commit  string
	date    string
)

func main() {
	c := c2{}
	opts.New(&c).Name("c2").Complete().Parse().Run()
}

type endpointName string

// type deployName string

// Currently suggesting the same prediction type is used for both endpoints and deploy names
// There is a bug in the opts/completion library such that only the first args complete get used.
func (endpointName) Complete(s string) []string {
	// TODO get endpoints
	return []string{"aaa", "abb", "c"}
}

// func (deployName) Complete(s string) []string {
// 	// TODO get deployNames
// 	return []string{"x", "y", "z"}
// }

func (*c2) Run() {
	fmt.Printf("version: %s\n", version)
}

type c2 struct {
	Help         struct{} `opts:"mode=cmd"`
	ListReleases struct{} `opts:"mode=cmd"`
	ShowLog      struct{} `opts:"mode=cmd"`
	Version      string   `opts:"mode=flag"` // the default mode
	Status       struct {
		ShowSlaves bool
	} `opts:"mode=cmd"`
	Start struct {
		Release  string `opts:"mode=arg"`
		Asdeploy string `opts:"mode=arg"`
	} `opts:"mode=cmd"`
	Stop struct {
		Deploy string `opts:"mode=arg"`
	} `opts:"mode=cmd"`
	RestartFrontendProxy struct {
	} `opts:"mode=cmd"`
	Connect struct {
		Endpoint endpointName `opts:"mode=arg"`
		Deploy   endpointName `opts:"mode=arg"`
	} `opts:"mode=cmd"`
	Disconnect struct {
		Endpoint endpointName `opts:"mode=arg"`
	} `opts:"mode=cmd"`
	Select struct {
		Release string `opts:"mode=arg"`
	} `opts:"mode=cmd"`
	FetchContext struct {
		Retry bool
	} `opts:"mode=cmd"`
	Unpack struct {
		Release string `opts:"mode=arg"`
		Todir   string `opts:"mode=arg"`
	} `opts:"mode=cmd"`
	ExpandTemplate struct {
		TemplatePath string `opts:"mode=arg,name=templatePath"`
		DestPath     string `opts:"mode=arg,name=destPath"`
	} `opts:"mode=cmd"`
	ShowDefaultNginxConfig struct {
	} `opts:"mode=cmd"`
	AwsDockerLoginCmd struct {
	} `opts:"mode=cmd"`
	GenerateSslCertificate struct {
	} `opts:"mode=cmd"`
	SlaveFlush struct {
	} `opts:"mode=cmd"`
	SlaveUpdate struct {
		Repeat int
	} `opts:"mode=cmd"`
}
