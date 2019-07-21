package function

import (
	faasflow "github.com/s8sg/faas-flow"
	consulStateStore "github.com/s8sg/faas-flow-consul-statestore"
	minioDataStore "github.com/s8sg/faas-flow-minio-datastore"
	"os"
)

// Define provide definiton of the workflow
func Define(flow *faasflow.Workflow, context *faasflow.Context) (err error) {

	dag := flow.Dag()
	dag.Node("n1").Modify(func(data []byte) ([]byte, error) {
		data = []byte(string(data) + "modifier")
		return data, nil
	})
	conditionalDags := dag.ConditionalBranch("C",
		[]string{"c1", "c2"}, // possible conditions
		func(response []byte) []string {
			return []string{"c1", "c2"}
		},
		faasflow.Aggregator(func(data map[string][]byte) ([]byte, error) {
			result := ""
			for key, value := range data {
				result = result + "+ (" + key + "=" + string(value) + ")"
			}
			return []byte(result), nil
		}),
	)

	conditionalDags["c2"].Node("n1").Apply("func1").Modify(func(data []byte) ([]byte, error) {
		data = []byte(string(data) + "modifier")
		return data, nil
	})
	c1ConditionalDags := conditionalDags["c1"].ConditionalBranch("X",
		[]string{"x1", "x2", "x3"}, // possible conditions
		func(data []byte) []string {
			return []string{"x2", "x3"}
		},
		faasflow.Aggregator(func(data map[string][]byte) ([]byte, error) {
			result := ""
			for key, value := range data {
				result = result + "+ (" + key + "=" + string(value) + ")"
			}
			return []byte(result), nil
		}),
	)

	c1ConditionalDags["x1"].Node("n1").Modify(func(data []byte) ([]byte, error) {
		data = []byte(string(data) + "modifier")
		return data, nil
	})

	c1ConditionalDags["x2"].Node("n1").Apply("func1").Modify(func(data []byte) ([]byte, error) {
		data = []byte(string(data) + "modifier")
		return data, nil
	})
	c1ConditionalDags["x2"].Node("n2").Apply("func2")
	c1ConditionalDags["x2"].Edge("n1", "n2")

	c1ConditionalDags["x3"].Node("n1").Apply("func1").Apply("func2")

	dag.Node("n2").Callback("http://gateway:8080/function/fake-storage")
	dag.Edge("n1", "C")
	dag.Edge("C", "n2")
	return
}

// DefineStateStore provides the override of the default StateStore
func DefineStateStore() (faasflow.StateStore, error) {
	consulss, err := consulStateStore.GetConsulStateStore(
		os.Getenv("consul_url"),
		os.Getenv("consul_dc"),
	)
	if err != nil {
		return nil, err
	}
	return consulss, nil
}

// DefineDataStore provides the override of the default DataStore
func DefineDataStore() (faasflow.DataStore, error) {
	miniods, err := minioDataStore.InitFromEnv()
	if err != nil {
		return nil, err
	}
	return miniods, nil
}
