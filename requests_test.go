package singularity

import (
	"testing"
)

func TestOnDemandRequest(t *testing.T) {
	expectedID := "test-ondemand"
	expectedType := "ON_DEMAND"
	req := NewOnDemandRequest(expectedID)
	if req.ID != expectedID {
		t.Errorf("Got %s, expected %s", req.ID, expectedID)
	}
	if req.RequestType != expectedType {
		t.Errorf("Got %s, expected %s", req.RequestType, expectedType)
	}
}

func TestOnDemandRequestRetries(t *testing.T) {
	expectedID := "test-ondemand"
	var expectedRetries int64 = 10
	req := NewOnDemandRequest(expectedID)
	req.Retries(expectedRetries)
	if req.NumRetriesOnFailure != expectedRetries {
		t.Errorf("Got %v, expected %v", req.NumRetriesOnFailure, expectedRetries)
	}
}

func TestOnDemandSkipHealthChecks(t *testing.T) {
	expectedID := "test-ondemand"
	expectedSkipHealthCheck := true
	req := NewOnDemandRequest(expectedID)
	req.SetSkipHealthchecks(expectedSkipHealthCheck)
	if req.SkipHealthchecks != expectedSkipHealthCheck {
		t.Errorf("Got %v, expected %v", req.NumRetriesOnFailure, expectedSkipHealthCheck)
	}
}

func TestOnDemandSetTaskExecutionLimit(t *testing.T) {
	expectedID := "test-ondemand"
	expectedTimeLimit := 500
	req := NewOnDemandRequest(expectedID)
	req.SetTaskExecutionLimit(expectedTimeLimit)
	if req.TaskExecutionTimeLimitMillis != expectedTimeLimit {
		t.Errorf("Got %v, expected %v", req.TaskExecutionTimeLimitMillis, expectedTimeLimit)
	}
}

func TestOnDemandSetTaskPriorityLevel(t *testing.T) {
	expectedID := "test-ondemand"
	expectedPriorityLevel := 3
	req := NewOnDemandRequest(expectedID)
	req.SetTaskPriorityLevel(expectedPriorityLevel)
	if req.TaskPriorityLevel != expectedPriorityLevel {
		t.Errorf("Got %v, expected %v", req.TaskPriorityLevel, expectedPriorityLevel)
	}
}

func TestOnDemandSetBounceAfterScale(t *testing.T) {
	expectedID := "test-ondemand"
	expectedBool := true
	req := NewOnDemandRequest(expectedID)
	req.SetBounceAfterScale(expectedBool)
	if req.BounceAfterScale != expectedBool {
		t.Errorf("Got %v, expected %v", req.BounceAfterScale, expectedBool)
	}
}

func TestNewServiceRequest(t *testing.T) {
	expectedID := "test-service"
	expectedType := "SERVICE"
	var n int64 = 3
	req := NewServiceRequest(expectedID, n)
	if req.ID != expectedID {
		t.Errorf("Got %s, expected %s", req.ID, expectedID)
	}
	if req.Instances != n {
		t.Errorf("Got %v, expected %v", req.Instances, n)
	}
	if req.RequestType != expectedType {
		t.Errorf("Got %s, expected %s", req.RequestType, expectedType)
	}
}

func TestServiceRequestInstances(t *testing.T) {
	expectedID := "test-service"
	var expectedInstances int64 = 25
	req := NewServiceRequest(expectedID, expectedInstances)
	req.SetInstances(expectedInstances)
	if req.Instances != expectedInstances {
		t.Errorf("Got %v, expected %v", req.Instances, expectedInstances)
	}
}
func TestServiceSetLoadBalanced(t *testing.T) {
	expectedID := "test-service"
	expectedBool := true
	var expectedInstances int64 = 25
	req := NewServiceRequest(expectedID, expectedInstances)
	req.SetLoadBalanced(true)
	if req.LoadBalanced != expectedBool {
		t.Errorf("Got %v, expected %v", req.LoadBalanced, expectedBool)
	}
}

func TestNewScheduledRequest(t *testing.T) {
	expectedID := "test-scheduled"
	expectedType := "SCHEDULED"
	expectedCron := "*/30 * * * *"
	req, _ := NewScheduledRequest(expectedID, expectedCron)
	if req.ID != expectedID {
		t.Errorf("Got %s, expected %s", req.ID, expectedID)
	}
	if req.RequestType != expectedType {
		t.Errorf("Got %s, expected %s", req.RequestType, expectedType)
	}
	if req.Schedule != expectedCron || req.Schedule == "" {
		t.Errorf("Got %v, expected %v", req.Schedule, expectedCron)
	}

	invalidCron := "* * * * * * *"
	expectedError := "Parse * * * * * * cron schedule error Expected exactly 5 fields, found 6: * * * * * *"
	reqError, err := NewScheduledRequest(expectedID, invalidCron)

	if err == nil {
		t.Errorf("Got %v, expected %s", err, expectedError)
	}
	if reqError.Schedule != "" {
		t.Errorf("Got %v, expected %s", err, expectedError)
	}
}

func TestNewWorkerRequest(t *testing.T) {
	expectedID := "test-worker"
	expectedType := "WORKER"
	var n int64 = 5
	req := NewWorkerRequest(expectedID, n)
	if req.ID != expectedID {
		t.Errorf("Got %s, expected %s", req.ID, expectedID)
	}
	if req.Instances != n {
		t.Errorf("Got %v, expected %v", req.Instances, n)
	}
	if req.RequestType != expectedType {
		t.Errorf("Got %s, expected %s", req.RequestType, expectedType)
	}
}

func TestNewRunOnceRequet(t *testing.T) {
	expectedID := "test-runonce"
	expectedType := "RUN_ONCE"
	var n int64 = 2
	req := NewRunOnceRequest(expectedID, n)
	if req.ID != expectedID {
		t.Errorf("Got %s, expected %s", req.ID, expectedID)
	}
	if req.Instances != n {
		t.Errorf("Got %v, expected %v", req.Instances, n)
	}
	if req.RequestType != expectedType {
		t.Errorf("Got %s, expected %s", req.RequestType, expectedType)
	}
}

/*  Fix this http request test. Checkout gomega http client test
https://onsi.github.io/gomega/#ghttp-testing-http-clients

func TestClient_GetRequests(t *testing.T) {
       request := SingularityRequest{
               ID:                  "test-geostreamoffsets-launch-sqs-connector",
               requestType:         "RUN_ONCE",
               NumRetriesOnFailure: 3,
       }
       activeDeploy := ActiveDeploy{
               RequestID: "test-geostreamoffsets-launch-sqs-connector",
               DeployID:  "prodromal",
               Timestamp: 1503451301091,
       }
       deployState := SingularityDeployState{
               RequestID:    "test-geostreamoffsets-launch-sqs-connector",
               ActiveDeploy: activeDeploy,
       }
       data := Requests{
               Request{
                       SingularityRequest: request,
                       State:              "ACTIVE",
                       SingularityDeployState: deployState,
               },
       }

       config := Config{
               Host: "127.0.0.1",
       }
       c := New(config)

       httpmock.Activate()
       defer httpmock.DeactivateAndReset()
       da, _ := json.Marshal(data)
       httpmock.NewMockTransport().RegisterResponder("GET", "http://foo.com/bar", httpmock.NewStringResponder(200, string(da)))

       req, _, _ := c.SuperAgent.Get("http://foo.com/bar").End()
       //      req, _, _ := c.GetRequests()
       //      req, _ := http.NewRequest("GET", "http://foo.com/bar", nil)

       fmt.Println("val: ", req)
       //res, _ := (&http.Client{}).Do(req)
       z, _ := ioutil.ReadAll(req.Body)
       fmt.Println("val: ", string(z))

               st.Expect(t, err, nil)
               st.Expect(t, res.StatusCode, 200)

               // Verify that we don't have pending mocks
               st.Expect(t, gock.IsDone(), true)
}
*/
