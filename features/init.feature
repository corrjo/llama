Feature: init project
  In order to start a project
  I need to be able to initialize a workspace


  Scenario: Start a new go project
    Given I need to create my_project
    When I invoke llama init my_project
    Then a directory with all needed files will be created
