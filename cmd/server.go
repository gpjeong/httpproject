package cmd

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"httpproject/internal/config"
	"httpproject/internal/datastore"
	"httpproject/internal/schedule"
	"httpproject/rest"
	"httpproject/server"
	lr "httpproject/util/logger"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	Logger *lr.Logger
)

func init() {
	config.DefaultServiceConfigFromEnv()
	Logger = lr.New(config.ServerConfig.Logger)
	logoPrint()
}

func packageName() string {
	pc, _, _, _ := runtime.Caller(1)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	pkage := ""
	if parts[pl-2][0] == '(' {
		pkage = strings.Join(parts[0:pl-2], ".")
	} else {
		pkage = strings.Join(parts[0:pl-1], ".")
	}
	packageNames := strings.Split(pkage, "/")
	config.ServerName = packageNames[0]
	return packageNames[0]
}

func logoPrint() {
	serviceLogo := `` + packageName()
	fmt.Println(figure.NewFigure(serviceLogo, "doom", true))
	lr.Log.Info().Msgf("Started at : %v", time.Now().Format(time.RFC3339))
	lr.Log.Info().Msgf("Server Logger Level : %v" + lr.LogLevelToString[config.ServerConfig.Logger.Level])
}

func ServerRun() {
	config.DefaultServiceConfigFromEnv()
	rest.NewClientInit()
	datastore.TaskManger = datastore.NewTaskManager()

	server.ApiTest()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		schedule.OjtProjectScheduler()
	}()

	wg.Wait()

}
