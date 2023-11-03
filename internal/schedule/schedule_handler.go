package schedule

import (
	"github.com/go-co-op/gocron"
	"httpproject/client"
	"time"
)

var OneMinuteCron = "*/1 * * * *"
var FiveMinuteCron = "*/5 * * * *"

func OjtProjectScheduler() {
	gocron := gocron.NewScheduler(time.UTC)

	gocron.Cron(OneMinuteCron).Do(func() {
		client.GetApiData("정광필")
		client.PostApiData("7", "임영웅", "100")
		client.DeleteApiData("7", "임영웅")
	})

	gocron.StartAsync()
}
