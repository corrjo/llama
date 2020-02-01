package main

import (
	"github.com/DATA-DOG/godog"
	"llama/cmd"
)

func iNeedToCreateMy_project() error {
	projectName := "my_project"
	return nil
}

func iInvokeLlamaInitMy_project(projectName string) error {
	cmd.CreateProject(projectName)
	return nil
}

func aDirectoryWithAllNeededFilesWillBeCreated() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I need to create my_project$`, iNeedToCreateMy_project)
	s.Step(`^I invoke llama init my_project$`, iInvokeLlamaInitMy_project)
	s.Step(`^a directory with all needed files will be created$`, aDirectoryWithAllNeededFilesWillBeCreated)
}
