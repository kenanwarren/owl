[![Go Report Card](https://goreportcard.com/badge/redorb/owl)](https://goreportcard.com/report/redorb/owl)
[![](https://godoc.org/github.com/redorb/owl?status.svg)](http://godoc.org/github.com/redorb/owl)
[![Build Status](https://travis-ci.org/Redorb/owl.svg?branch=master)](https://travis-ci.org/Redorb/owl)
[![codecov](https://codecov.io/gh/redorb/owl/branch/master/graph/badge.svg)](https://codecov.io/gh/redorb/owl)
### Go client for Overwatch League REST API

### New Client

```go
import "github.com/redorb/owl"

// Create a client instance not using Chinese API endpoint
c, err := owl.NewClient(false)
```
