# {{cookiecutter.model_name.capitalize()}}s management back

Manages {{cookiecutter.model_name}}s that use Divilo APP.

## Dev tools

* Setup dev: `make setup`
* See available commands: `make help`

## Data model

```plantuml
@startuml
skinparam padding 15
component {{cookiecutter.project_name}} {
  rectangle {{cookiecutter.model_name.capitalize()}}
  rectangle {{cookiecutter.model_name.capitalize()}}Error
}

{{cookiecutter.model_name.capitalize()}} "1" - "0..n" {{cookiecutter.model_name.capitalize()}}Error : has >
@enduml
```

### DynamoDB table: {{cookiecutter.project_name}}-{{cookiecutter.model_name}}s

#### Key definition

* Partition key (HASH): `{{cookiecutter.model_name}}id`

#### {{cookiecutter.model_name.capitalize()}} records

| Attribute | Description             | Example                                       |
|-----------|-------------------------|-----------------------------------------------|
| {{cookiecutter.model_name}}id  | {{cookiecutter.model_name.capitalize()}} Identification   | `81a0aabc-7fe1-4b42-a387-d9f685a212e3`        |
| meta      | Map with {{cookiecutter.model_name}} data    |                                               |
| createdat | Record creation date    | 1646644427                                    |
| updatedat | Record last update date | 1646644427                                    |
