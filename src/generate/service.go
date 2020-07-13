package generate

import (
	"encoding/json"
	"os"
	"path"
)

type ProjectProvider struct {
	GitLabPath  string `json:"GitLabPath"`
	ProjectPath string `json:"ProjectPath"`
	Project     string `json:"Project"`
	Port        int    `json:"Port"`
	Model       Model  `json:"Model"`
}

// 为了便于在模板中替换字符串，需要使用冗余字段
type Model struct {
	Prefix     string `json:"prefix"`
	ModelName  string `json:"ModelName"`
	LowerName  string `json:"LowerName"`
	GitLabPath string `json:"GitLabPath"`
	Project    string `json:"Project"`
	SimplyName string `json:"SimplyName"`
	Port       int    `json:"Port"`
}

func (provider *ProjectProvider) String() string {
	if dataInBytes, err := json.Marshal(&provider); err == nil {
		return string(dataInBytes)
	}

	return ""
}

func (provider *ProjectProvider) Generate() error {

	return nil
}

func (provider *ProjectProvider) CopySwagger() error {
	// https://zhijie.me/go/17.html

	return nil
}

func (provider *ProjectProvider) generateController() error {
	content, err := Render("controllerFile", TemplateControllerFilePath, provider.Model)
	if err != nil {
		return err
	}

	folderPath := provider.ProjectPath + "/src/apis/"

	return provider.generateFile(folderPath, provider.Model.LowerName+".go", content)
}

func (provider *ProjectProvider) generateDockerFile() error {
	content, err := Render("dockerFile", TemplateDockerFilePath, provider.Model)
	if err != nil {
		return err
	}

	return provider.generateFile(provider.ProjectPath, "Dockerfile", content)
}

func (provider *ProjectProvider) generateDockerIgnoreFile() error {
	content, err := Render("dockerIgnore", TemplateDockerIgnoreFilePath, provider.Model)
	if err != nil {
		return err
	}

	return provider.generateFile(provider.ProjectPath, "Dockerfile", content)
}

func (provider *ProjectProvider) generateGitIgnoreFile() error {
	content, err := Render("gitIgnore", TemplateGitIgnoreFilePath, provider.Model)
	if err != nil {
		return err
	}

	return provider.generateFile(provider.ProjectPath, ".gitignore", content)
}

func (provider *ProjectProvider) generateKubernetesFile() error {
	content, err := Render("kubernetes", TemplateKubernetesFilePath, provider.Model)
	if err != nil {
		return err
	}

	folderPath := provider.ProjectPath + "/deploy/"
	fileName := provider.Project + ".yaml"

	return provider.generateFile(folderPath, fileName, content)
}

func (provider *ProjectProvider) generateMakeFile() error {
	content, err := Render("make", TemplateMakeFilePath, provider.Model)
	if err != nil {
		return err
	}

	return provider.generateFile(provider.ProjectPath, "Makefile", content)
}

func (provider *ProjectProvider) generateModel() error {
	content, err := Render("model", TemplateModelFilePath, provider.Model)
	if err != nil {
		return err
	}

	folderPath := provider.ProjectPath + "/src/models/"

	return provider.generateFile(folderPath, provider.Model.LowerName+".go", content)
}

func (provider *ProjectProvider) generateService() error {
	content, err := Render("service", TemplateServiceFilePath, provider.Model)
	if err != nil {
		return err
	}

	folderPath := provider.ProjectPath + "/src/service/"

	return provider.generateFile(folderPath, provider.Model.LowerName+".go", content)
}

func (provider *ProjectProvider) GenerateSQL() error {
	content, err := Render("sql", TemplateSQLFilePath, provider.Model)
	if err != nil {
		return err
	}

	folderPath := provider.ProjectPath + "/sql/"
	fileName := provider.Model.Project + "_" + provider.Model.LowerName + ".ddl.sql"

	return provider.generateFile(folderPath, fileName, content)
}

func (provider *ProjectProvider) generateFile(pathName string, fileName string, content string) error {
	if !IsExist(pathName) {
		if err := os.MkdirAll(pathName, 0777); err != nil {
			return err
		}
	}

	f, err := os.Create(path.Join(pathName, fileName))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
