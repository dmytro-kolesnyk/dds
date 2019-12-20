
##### overal help (command list)
```
go run main.go 
```

##### run command 
```
go run main.go upload    
               download
               delete
               list
```
##### run help for command 
```
go run main.go [command_name] -h, --help

command_name: upload
              download
              delete
              list
```
##### run upload command example
```
go run main.go upload -p ~/home -l -s sfc
```
##### run download command example
```
go run main.go download -p ~/home -u f4c8de96-4e03-4772-b83c-f8dfbe64e998 -b
```
##### run delete command example
```
go run main.go delete -u f4c8de96-4e03-4772-b83c-f8dfbe64e998
```
##### run list command example
```
go run main.go list
```



     
                
