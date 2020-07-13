package generate

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

var p *ProjectProvider

func init() {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		log.Fatalln("the environment variable GOPATH is not set")
		return
	}

	model := "Pipeline"
	dir := "github.com/xdhuxc"
	project := "wang-test"
	port := 1234

	dataModel := Model{
		ModelName:  ConvertModel(model),
		LowerName:  Convert2Camel(model),
		GitLabPath: dir,
		Project:    project,
		SimplyName: GenerateSimplyName(model),
	}

	p = &ProjectProvider{
		Port:        port,
		Project:     project,
		ProjectPath: fmt.Sprintf(goPath+"/src/%s/%s", dir, project),
		Model:       dataModel,
		GitLabPath:  dir,
	}
}

func TestProjectProvider_GenerateService(t *testing.T) {
	err := p.generateService()
	require.NoError(t, err)
}

func TestProjectProvider_GenerateController(t *testing.T) {
	err := p.generateController()
	require.NoError(t, err)
}
