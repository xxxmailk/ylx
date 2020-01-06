
#### 1. 转化首字母为大写字母：
```go
// Capitalize: change first character to upper
func Capitalize(str string) string 
```


### △ 时间处理
```go
// return start timestamp in this day
// 返回本天第一个时间节点的时间戳
func ThisDay(t time.Time) (start int64)

// return start timestamp in this month
// 返回本月第一个时间节点的时间戳
func ThisMonth(t time.Time) (start int64) 


// return start timestamp in this quarte
// 返回本季度第一个时间节点的时间戳
func ThisQuarte(t time.Time) (start int64) 

// return start timestamp in this year
// 返回本年第一个时间节点的时间戳
func ThisYear(t time.Time) (start int64)
```


