package main

import (
    "os/exec"

    "github.com/gin-gonic/gin"
)

func test1(c *gin.Context) {
    a := c.Param("test")
    var b string = a

    exec.Command("bash", "-c", b)
}

func main() {
    r := gin.New()
    r.GET("/test1", test1)

    r.Run(":8888")
}

