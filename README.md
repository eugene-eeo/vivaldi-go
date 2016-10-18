# vivaldi-go

[Vivaldi Coordinates](https://en.wikipedia.org/wiki/Vivaldi_coordinates) library
for Go. Vivaldi like other internet coordinate systems aim to predict (with some
small amount of error) the communication latency between two nodes without having
to resort to direct measurement. The distance between two coordinates is the
predicted latency. Some of Vivaldi's advantages include:

 - **Being very lightweight** – only the coordinate and relative error needs to be stored.
 - **Can piggyback on application traffic** – only the coordinates and relative errors need to be sent to other nodes.
 - **Resistant to topolgy changes** – new nodes that join the network will not affect the accuracy of old coordinates.
 - **No centralisation** – can be used in p2p networks.

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
