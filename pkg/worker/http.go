package worker

import (
	"github.com/woocoos/workflow/pkg/api"
	"github.com/woocoos/workflow/pkg/spec/bpmn"
	"net/http"
)

type HttpWorker struct {
}

func (hw *HttpWorker) ParseElement(ele *bpmn.ServiceTask, request *api.InstanceRequest) (*HttpWorker, error) {
	var (
		url, method string
		headers     = make(http.Header)
	)
	for _, header := range ele.Headers {
		switch header.Key {
		case "url":
			url = header.Value
		case "method":
			method = header.Value
		default:
			headers.Add(header.Key, header.Value)
		}
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = headers
}
