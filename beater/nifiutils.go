package beater

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// JSONStruct is the structure of the NIFI response
type JSONStruct struct {
	ProcessGroups []struct {
		Revision struct {
			ClientID string
			Version  int
		}
		ID       string
		URI      string
		Position struct {
			X float32
			Y float32
		}
		Permissions struct {
			CanRead  bool
			CanWrite bool
		}
		Bulletins []struct {
		}
		Component struct {
			ID            string
			ParentGroupID string
			Position      struct {
				X float32
				Y float32
			}
			Name      string
			Comments  string
			Variables struct {
			}
			RunningCount                 int
			StoppedCount                 int
			InvalidCount                 int
			DisabledCount                int
			ActiveRemotePortCount        int
			InactiveRemotePortCount      int
			UpToDateCount                int
			LocallyModifiedCount         int
			StaleCount                   int
			LocallyModifiedAndStaleCount int
			SyncFailureCount             int
			InputPortCount               int
			OutputPortCount              int
		}
		Status struct {
			ID                 string
			Name               string
			StatsLastRefreshed string
			AggregateSnapshot  struct {
				ID                   string
				Name                 string
				FlowFilesIn          int
				BytesIn              int
				Input                string
				FlowFilesQueued      int
				BytesQueued          int
				Queued               string
				QueuedCount          int
				QueuedSize           string
				BytesRead            int
				Read                 string
				BytesWritten         int
				Written              string
				FlowFilesOut         int
				BytesOut             int
				Output               string
				FlowFilesTransferred int
				BytesTransferred     int
				Transferred          string
				BytesReceived        int
				FlowFilesReceived    int
				Received             string
				BytesSent            int
				FlowFilesSent        int
				Sent                 string
				ActiveThreadCount    int
			}
		}
		RunningCount                 int
		StoppedCount                 int
		InvalidCount                 int
		DisabledCount                int
		ActiveRemotePortCount        int
		InactiveRemotePortCount      int
		UpToDateCount                int
		LocallyModifiedCount         int
		StaleCount                   int
		LocallyModifiedAndStaleCount int
		SyncFailureCount             int
		InputPortCount               int
		OutputPortCount              int
	}
}

// JSONConvert text
func JSONConvert(body []byte) JSONStruct {
	res := JSONStruct{}
	json.Unmarshal(body, &res)
	return res
}

// RequestNifi text
func RequestNifi(url string) []byte {
	res, _ := http.Get(url)
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return body
}
