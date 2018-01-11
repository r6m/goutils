# gotuils

Common functions


### Installing

```bash
go get -u github.com/rezam90/goutils
```

## Usage

```go
import "github.com/rezam90/goutils"
``` 

#### functions

```go
// Parse
ToInt(interface{}) int
ToInt64(interface{}) int64
FloatToString(float64, bool) string
ToBool(interface{}) bool // "true", "false", "1", "0", "1.0", true

// Go Patch
MapExists(map[string]interface{}, []string) bool
MapExists(interface{}, interface{}) bool // slice could be any type of slice

// Random
RandInt(min, max int) int
RandInt64() int64
RandString(length int) string

// Time
TimeIsToday(time.Time) bool
TimeSameDay(t1, t2 time.Time) bool
TimeSameMonth(t1, t2 time.Time) bool
TimeSameYear(t1, t2 time.Time) bool
TimeBeginningOfDay(time.Time, string) time.Time
TimeEndOfDay(time.Time, string) time.Time
```

