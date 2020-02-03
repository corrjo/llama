# llama

A tool for packaging and shipping serverless functions.

## The scaffold

`llama init my_project --lang go --cloud aws` will create basic scaffolding for a project

```
my_project/
    README.md
    config.yml
    src/code.go
    src/entrypoint.go -- potentially useful for switching between cloud providers?
    deployment/function.tf
    deployment/provider.tf
```

- README.md: a basic readme to be filled out with details about your project
- config.yml: a place to store project configuration, lang type, cloud type, etc
- src/code.(go|py|js): The source code for your serverless app
- deployment/(function.tf, provider.tf): the deployment code, this will be different depending on which cloud is specified
