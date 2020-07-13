package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/asaskevich/govalidator"
	log "github.com/sirupsen/logrus"

	"github.com/xdhuxc/go-web-generator/src/generate"
)

var (
	project = flag.String("project", "", "the project name")
	port    = flag.Int("port", 8080, "the web project port")
	model   = flag.String("model", "", "the model name")
	// http://127.0.0.1:8080/user/api/v1/ 中的 user 即为前缀
	prefix = flag.String("prefix", "", "prefix of request URL")
	// /Users/wanghuan/GolandProjects/GoPath/src/  gitlab.ushareit.me/sgt/devops  /scmp-cmdb-cloud-provider
	dir = flag.String("dir", "", "gitlab path of the project, which generally is between the GOPATH and project folder")

	format = flag.String("format", "json", "the log format of logrus")
	level  = flag.String("level", "INFO", "the log level of logrus")
)

func main() {
	flag.Parse()

	// 初始化 log
	switch *format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{})
	}
	level, err := log.ParseLevel(*level)
	if err == nil {
		log.SetLevel(level)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if r := govalidator.Trim(*project, ""); govalidator.IsNull(r) {
		log.Fatalln("please specify the project name")
	}
	if r := govalidator.Trim(*prefix, ""); govalidator.IsNull(r) {
		log.Fatalln("please specify the URL prefix")
	}
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		log.Fatalln("the environment variable GOPATH is not set")
		return
	}
	log.Printf("the default port is %d", *port)

	dataModel := generate.Model{
		ModelName:  generate.ConvertModel(*model),
		LowerName:  generate.Convert2Camel(*model),
		SimplyName: generate.GenerateSimplyName(*model),
		GitLabPath: *dir,
		Project:    *project,
		Port:       *port,
		Prefix:     *prefix,
	}
	p := generate.ProjectProvider{
		Port:        *port,
		Project:     *project,
		ProjectPath: fmt.Sprintf(goPath+"/src/%s/%s", *dir, *project),
		Model:       dataModel,
	}

	err = p.Generate()
	if err != nil {
		log.Fatal(err)
	}
}
