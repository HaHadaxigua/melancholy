# melancholy
This is a simple cloud storage system.

## How To Start
### 1.Required Environment
Install latest golang, latest MySQL.

SQL DDL in /melancholy/script/init.sql

### 2.Edit Config File
Before edit config file, there some words.
```
In dev environment, we use fixed OSS-Account to implement system.In this way, exist Account risk.
```

By the way, to start, you need add information to "/ect/application.yml",fill OSS and Cloud Module. 
 
You can input information via Web-Page,so you can input whatever you like temporary.

To connect your MySQL, you need to edit DataBase module.

### 3.Compile
In current project directory, just run the following code
```shell
go build .\cmd\melancholy.go
```
