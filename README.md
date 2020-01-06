### - 一般函数
```go
// search a string key from a slice
// 如果返回值为负数，表明该slice中不存在该该值
// 从两端往slice中间进行搜索，加快搜索速度
func SliceSearchString(s []string, key string) (rs int)
```

### -  字符串处理
```go
// Capitalize: change first character to upper
// 首字母字符串大写
func Capitalize(str string) string 
```


### -  时间处理
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


