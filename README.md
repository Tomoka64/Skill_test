# go-skilltest
This repository presents a command tool for searching a word from go-package. This provides three usages; it can be used to search a selected word from a selected go-package in terminal, to utilize server-mode and search through http request and to see all the search history in terminal. The result will include the selected filename, the selected keyword, the line number in which the searched word is located in the file, and the details. 


## Getting Started


### Installing


```
go get github.com/Tomoka64/go-skilltest
```
### Usage

Usage1: 'package-name' 'word' 

e.g.) fmt TODO
  
```
fmt TODO
```

Usage2: localhost

```
localhost
```

Usage3: history

```
history
```

### Basic Model

```
type Result struct {
	Filename string `json:"filename"`
	Keyword  string `json:"keyword"`
	Line     int    `json:"line"`
	Detail   string `json:"detail"`
}
```

## Author

Tomoka Yokomizo

