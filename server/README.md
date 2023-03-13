# go-lib/server
A http/https server module using Fiber

## Usage
1. Get a new server object pointer. Default is a http server at port 80. Override defaults using the ServerOptions.
2. Add routes and handlers.
3. Start the server

### ServerOptions
| Option | Description |
| --- | --- |
| IsHttps | boolean, default is false |
| CertFilePath | string, path to the .crt file |
| KeyFilePath | string, path to the .key file |
| Port | string, e.g. ":4400". Defaults to 80 for http and 443 for https |

## Example of http server with default port
```
package main
import (
	"github.com/leenasavkur/go-lib/server"

	"github.com/gofiber/fiber/v2"
)

func main() {
	myserver := server.New(nil)
	myserver.AddRoute("GET", "/", handler)
	myserver.Start()
}

func handler(c *fiber.Ctx) error {
	c.SendString("I'm a GET request!")
	return nil
}
```

## Example of https server at non default port
```
package main

import (
	"github.com/leenasavkur/go-lib/server"

	"github.com/gofiber/fiber/v2"
)

func main() {
	options := server.ServerOptions{
		Port:         ":4400",
		CertFilePath: "/mypath/abc.crt",
		KeyFilePath:  "/mypath/abc.decrypted.key",
		IsHttps:      true,
	}

	myserver := server.New(&options)
	myserver.AddRoute("GET", "/", handler)
	myserver.Start()
}

func handler(c *fiber.Ctx) error {
	c.SendString("I'm a GET request!")
	return nil
}

```
