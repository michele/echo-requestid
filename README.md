[![Build Status](https://travis-ci.org/michele/echo-requestid.svg?branch=master)](https://travis-ci.org/michele/echo-requestid) [![Go Report Card](https://goreportcard.com/badge/github.com/michele/echo-requestid)](https://goreportcard.com/report/github.com/michele/echo-requestid)

## Echo X-Request-Id middleware

### Usage

Include it in your project and then add the middleware like this:

```
import (
  requestid "github.com/michele/echo-requestid"
)

...

r := echo.New()
r.Use(requestid.RequestIdMiddleware)
```