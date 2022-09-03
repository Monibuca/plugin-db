package influxdb

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go.uber.org/zap"
	. "m7s.live/engine/v4"
)

type InfluxdbConfig struct {
	Server string
	Token string
	Org string
	Bucket string
}

func (conf *InfluxdbConfig) OnEvent(event any){
	switch event.(type) {
	case FirstConfig:
		client := influxdb2.NewClient(conf.Server, conf.Token)
		writeAPI := client.WriteAPI(conf.Org, conf.Bucket)
		errorsCh := writeAPI.Errors()
    // Create go proc for reading and logging errors
    go func() {
        for err := range errorsCh {
					InfluxdbPlugin.Error("write error",zap.Error(err))
        }
    }()
	}
}


var InfluxdbPlugin = InstallPlugin(&InfluxdbConfig{
	Server: "http://localhost:8086",
	Org:"m7s", Bucket: "test",
})