package client

import (
	"sync"

	"github.com/toffer/sendit/models"
)

var (
	requests  []*models.VerifiableRequest
	responses []*models.VerifiableResponse

	requestsLock  sync.Mutex = sync.Mutex{}
	responsesLock sync.Mutex = sync.Mutex{}
)

func init() {
	requests = make([]*models.VerifiableRequest, 0)
	responses = make([]*models.VerifiableResponse, 0)
}

func addReq(r *models.VerifiableRequest) {
	requestsLock.Lock()
	defer requestsLock.Unlock()

	requests = append(requests, r)
}

func addResp(r *models.VerifiableResponse) {
	responsesLock.Lock()
	defer responsesLock.Unlock()

	responses = append(responses, r)
}

func removeReq() *models.VerifiableRequest {
	requestsLock.Lock()
	defer requestsLock.Unlock()

	if len(requests) > 1 {
		r := requests[len(requests)-1]
		requests = requests[:len(requests)-1]
		return r
	} else if len(requests) == 1 {
		r := requests[0]
		requests = make([]*models.VerifiableRequest, 0)
		return r
	} else {
		return nil
	}
}

func removeResp() *models.VerifiableResponse {
	responsesLock.Lock()
	defer responsesLock.Unlock()

	if len(responses) > 1 {
		r := responses[len(responses)-1]
		responses = responses[:len(responses)-1]
		return r
	} else if len(responses) == 1 {
		r := responses[0]
		requests = make([]*models.VerifiableRequest, 0)
		return r
	} else {
		return nil
	}
}

func clearReqs() {
	for sizeReqs() > 0 {
		_ = removeReq()
	}
}

func clearResps() {
	for sizeResps() > 0 {
		_ = removeResp()
	}
}

func sizeReqs() int {
	return len(requests)
}

func sizeResps() int {
	return len(responses)
}
