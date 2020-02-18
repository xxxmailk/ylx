# YLX的go函数库

== ylx's golang utils ==

- [YLX的go函数库](#ylx-go---)
    + [- 一般函数](#------)
    + [- 路径处理](#------)
    + [- 字符串处理](#--------)
    + [- 时间处理](#-------)
    
### - 一般函数
```go
// search a string key from a slice
// 如果返回值为负数，表明该slice中不存在该该值
// 从两端往slice中间进行搜索，加快搜索速度
func SliceSearchString(s []string, key string) (rs int)
```

### - 路径处理
```go
// check dir is exist
// 检查目录是否存在
// 如果路径存在 exist = true
// 如果路径存在且该路径为一个目录 exist = true, isDir = true
func DirIsExist(path string) (exist bool, isDir bool)
```

```go
// check or create dir
// 如果返回值为nil，路径存在且正确
// 如果路径不存在，该函数会自动创建路径 mode = 0644
func CheckOrCreateDir(path string) error
```
### - 字符串处理
```go
// Capitalize: change first character to upper
// 首字母字符串大写
func Capitalize(str string) string 
```


### - 时间处理
```go
// return start timestamp in this day
// 返回本天第一个时间节点的时间戳
func ThisDay(t time.Time) (start int64)
```
```go
// return start timestamp in this month
// 返回本月第一个时间节点的时间戳
func ThisMonth(t time.Time) (start int64) 
````

```go
// return start timestamp in this quarte
// 返回本季度第一个时间节点的时间戳
func ThisQuarte(t time.Time) (start int64) 
```
```go
// return start timestamp in this year
// 返回本年第一个时间节点的时间戳
func ThisYear(t time.Time) (start int64)
```


