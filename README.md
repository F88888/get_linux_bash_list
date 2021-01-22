## 1. Basic introduction

> Get the bash built-in command list of the current system, mainly used with go exec module

## 2. 使用说明

```
var bashList []string
var bashInfo = GetBash()
bashList = bashInfo.GetList

or 

bashInfo.InBash("cd")
```