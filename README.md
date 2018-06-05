[![Go Report Card](https://goreportcard.com/badge/redorb/owl)](https://goreportcard.com/report/redorb/owl)
### GO client for Overwatch League REST API

### New Client

```go
import "github.com/redorb/owl"

// Create a client instance not using Chinese API endpoint
c, err := owl.NewClient(false)
```