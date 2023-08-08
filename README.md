How to run app:
1. run main.go from server folder

```shell
 go run main.go
```


2. To see all quiz questions go in quiz folder and run:

```shell
go run quiz.go  get-questions
```
Obs. Server from step1 must be running in order to run quiz.go


3.To start quiz go in quiz folder and run:

```shell
go run quiz.go  start-quiz 
```
Obs. Server from step1 must be running in order to run quiz.go

--Local storage will be up and store all quiz scores as long as main.go is running. 

