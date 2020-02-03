package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"llama/cmd"
	//"os"
)

var projectName []string

func iNeedToCreateMy_project(name string) error {
	projectName = []string{name}
	return nil
}

func iInvokeLlamaInitMy_project() error {
	cmd.CreateProject(projectName)
	return nil
}

func aDirectoryWithAllNeededFilesWillBeCreated() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I need to create (\w+)$`, iNeedToCreateMy_project)
	s.Step(`^I invoke llama init (\w+)$`, iInvokeLlamaInitMy_project)
	s.Step(`^a directory with all needed files will be created$`, aDirectoryWithAllNeededFilesWillBeCreated)

	s.AfterScenario(func(interface{}, error) {
		//current_path, err := os.Getwd()
		//if err != nil {
		//fmt.Println(err)
		//}
		fmt.Println(projectName)
		//err = os.RemoveAll(current_path + "/" + projectName[0])
		//if err != nil {
		//fmt.Println(err)
		//}
	})

}
