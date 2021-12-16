/*
 * TestKube API
 *
 * TestKube provides a Kubernetes-native framework for test definition, execution and results
 *
 * API version: 1.0.0
 * Contact: testkube@kubeshop.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package testkube

// test execution request body
type TestExecutionRequest struct {
	// script execution custom name
	Name string `json:"name,omitempty"`
	// script kubernetes namespace (\"testkube\" when not set)
	Namespace string `json:"namespace,omitempty"`
	// execution params passed to executor
	Params map[string]string `json:"params,omitempty"`
}
