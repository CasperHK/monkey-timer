# Monkey Timer

<br/>

## Installation
```bash
go get github.com/CasperHK/monkey-timer
```
<br/>
## Usage
```go
import (
    timer "github.com/CasperHK/monkey-timer"
)
```

```go
timer := NewTimer()
timer.SetTime(12, 30)
timer.SetAlarm(2, 45, func() {

})
timer.SetPeriodAction(0.5, timer.Second, func() {

})
```
<br/>
## License
