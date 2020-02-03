package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"llama/cmd"
	"os"
	"path/filepath"
	"reflect"
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
	var files []string
	expected := []string{"./my_project", "my_project/deployment", "my_project/src"}

	err := filepath.Walk("./my_project",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			files = append(files, path)
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(files, expected)
	if !reflect.DeepEqual(files, expected) {
		return fmt.Errorf("files should equal %v but equals %v", expected, files)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I need to create (\w+)$`, iNeedToCreateMy_project)
	s.Step(`^I invoke llama init (\w+)$`, iInvokeLlamaInitMy_project)
	s.Step(`^a directory with all needed files will be created$`, aDirectoryWithAllNeededFilesWillBeCreated)

	s.AfterScenario(func(interface{}, error) {
		current_path, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		err = os.RemoveAll(current_path + "/" + projectName[0])
		if err != nil {
			fmt.Println(err)
		}
	})

}
