package beater

import (
	"fmt"
	"strconv"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/sigrist/nifibeat/config"
)

// Nifibeat structure
type Nifibeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Nifibeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Nifibeat) request(url string, beatname string) bool {
	logp.Info("Calling: " + url)

	response := RequestNifi(url)

	result := JSONConvert(response)

	logp.Info("Process Groups: " + strconv.Itoa(len(result.ProcessGroups)))
	if len(result.ProcessGroups) != 0 {

		for process := range result.ProcessGroups {
			event := beat.Event{
				Timestamp: time.Now(),
				Fields: common.MapStr{
					"type": beatname,
					"URL":  bt.config.URL,
					"Nifi": result.ProcessGroups[process],
				},
			}
			bt.client.Publish(event)

			logp.Info("Evento Enviado ao Kibana")

			urlRecursiva := fmt.Sprintf(bt.config.URL+"/nifi-api/process-groups/%s/process-groups", result.ProcessGroups[process].ID)

			bt.request(urlRecursiva, beatname)
		}
	}

	return true
}

// Run Execute
func (bt *Nifibeat) Run(b *beat.Beat) error {
	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	var url string
	var urlGroups string

	url = bt.config.URL
	urlGroups = url + "/nifi-api/process-groups/root/process-groups"

	bt.request(urlGroups, b.Info.Name)
	return nil
}

// Stop stop
func (bt *Nifibeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
