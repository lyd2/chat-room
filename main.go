package main

import (
	"fmt"
	"github.com/lyd2/live/pkg/setting"
	"github.com/lyd2/live/routers"
)

func main() {

	engine := routers.InitRouter()

	_ = engine.Run(fmt.Sprintf(":%d", setting.HTTPPort))

}