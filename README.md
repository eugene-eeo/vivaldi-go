# vivaldi-go

[Vivaldi Coordinates](https://en.wikipedia.org/wiki/Vivaldi_coordinates)
library for Go. Vivaldi Coordinates, like other internet coordinates aims
to predict the communication latency between two nodes without direct
measurement. The distance between two coordinates is the prediction.

```go
import "github.com/eugene-eeo/vivaldi-go"

func main() {
    local := vivaldi.NewContext()
    remote := vivaldi.NewHVector(
        23.0, // x
        45.0, // y
        10.0, // height
    )
    rtt := 5.0
    local.update(rtt, vivaldi.NewContextFromValues(
        remote,
        5.0, // error estimate
    ))
}
```
