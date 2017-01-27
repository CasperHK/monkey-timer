# Monkey Timer

<br/>

## Installation
```bash
go get github.com/CasperHK/mtimer
```
<br/>
## Usage
```go
import (
    timer "github.com/CasperHK/mtimer"
)
```

```go
timer := mtimer.NewTimer()
timer.SetTime(`12:15`)
timer.SetAlarm(`2:45`, func() {

})
timer.SetPeriodAction(0.5, timer.Second, func() {

})
timer.Print()
timer.Start()
timer.Stop()
```
<br/>
## License
