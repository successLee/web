web
===
##The web is an easy free web framework to use the go source code.
##[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

#for example
    package main
    
    import (
    	"github.com/yzw/web"
    	"io"
    	"log"
    )
    
    func Hello(c *web.Controller) bool {
    	io.WriteString(c.Response, "hello")
    	return false
    }
    
    func main() {
    	// add handle
    	web.Handle.Route("/", "GET,POST", Hello)
    
    	// run server
    	webAddr := ":8000"
    	log.Printf("run server addr(%s)", webAddr)
    	web.Run(webAddr)
    }
