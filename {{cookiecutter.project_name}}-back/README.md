# {{cookiecutter.model_name}} management back

Manages {{cookiecutter.model_name}} that use Divilo APP.

## Dev tools

* Setup dev: `make setup`
* See available commands: `make help`

## Data model

```plantuml
@startuml
skinparam padding 15
component {{cookiecutter.project_name}} {
  rectangle Device
  rectangle DeviceError
}

Device "1" - "0..n" DeviceError : has >
@enduml
```

### DynamoDB table: {{cookiecutter.project_name}}-{{cookiecutter.model_name}}

#### Key definition

* Partition key (HASH): `deviceid`

#### Device records

| Attribute | Description             | Example                                       |
|-----------|-------------------------|-----------------------------------------------|
| deviceid  | Device Identification   | `81a0aabc-7fe1-4b42-a387-d9f685a212e3`        |
| meta      | Map with device data    |                                               |
| createdat | Record creation date    | 1646644427                                    |
| updatedat | Record last update date | 1646644427                                    |
