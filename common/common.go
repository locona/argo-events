/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"github.com/argoproj/argo-events/pkg/apis/gateway"
)

const (

	// ErrorResponse for http request
	ErrorResponse = "Error"

	// StandardTimeFormat is time format reference for golang
	StandardTimeFormat = "2006-01-02 15:04:05"

	// StandardYYYYMMDDFormat formats date in yyyy-mm-dd format
	StandardYYYYMMDDFormat = "2006-01-02"

	// DefaultControllerNamespace is the default namespace where the sensor and gateways controllers are installed
	DefaultControllerNamespace = "argo-events"
)

// ENV VARS
const (
	// EnvVarKubeConfig is the path to the Kubernetes configuration
	EnvVarKubeConfig = "KUBE_CONFIG"

	// EnvVarDebugLog is the env var to turn on the debug mode for logging
	EnvVarDebugLog = "DEBUG_LOG"
)

// LABELS
const (
	// LabelOperation is a label for an operation in framework
	LabelOperation = "operation"

	// LabelEventSource is label for event name
	LabelEventSource = "event-source"
)

// SENSOR CONSTANTS
const (
	// SensorServiceEndpoint is the endpoint to dispatch the event to
	SensorServiceEndpoint = "/"

	// SensorName refers env var for name of sensor
	SensorName = "SENSOR_NAME"

	// SensorNamespace is used to get namespace where sensors are deployed
	SensorNamespace = "SENSOR_NAMESPACE"

	// LabelSensorName is label for sensor name
	LabelSensorName = "sensor-name"

	// EnvVarSensorControllerInstanceID is used to get sensor controller instance id
	EnvVarSensorControllerInstanceID = "SENSOR_CONTROLLER_INSTANCE_ID"
)

// GATEWAY CONSTANTS
const (
	// LabelGatewayControllerName is the default deployment name of the gateway-controller-controller
	LabelGatewayControllerName = "gateway-controller"

	// EnvVarGatewayControllerConfigMap contains name of the configmap to retrieve gateway-controller configuration from
	EnvVarGatewayControllerConfigMap = "GATEWAY_CONTROLLER_CONFIG_MAP"

	// EnvVarGatewayControllerInstanceID is used to get gateway controller instance id
	EnvVarGatewayControllerInstanceID = "GATEWAY_CONTROLLER_INSTANCE_ID"

	// EnvVarGatewayControllerName is used to get name of gateway controller
	EnvVarGatewayControllerName = "GATEWAY_CONTROLLER_NAME"

	// GatewayControllerConfigMapKey is the key in the configmap to retrieve gateway-controller configuration from.
	// Content encoding is expected to be YAML.
	GatewayControllerConfigMapKey = "config"

	// EnvVarGatewayName refers env var for name of gateway
	EnvVarGatewayName = "GATEWAY_NAME"

	// EnvVarGatewayNamespace is namespace where gateway controller is deployed
	EnvVarGatewayNamespace = "GATEWAY_NAMESPACE"

	//LabelKeyGatewayControllerInstanceID is the label which allows to separate application among multiple running gateway-controller controllers.
	LabelKeyGatewayControllerInstanceID = gateway.FullName + "/gateway-controller-instanceid"

	// LabelGatewayKeyPhase is a label applied to gateways to indicate the current phase of the gateway-controller (for filtering purposes)
	LabelGatewayKeyPhase = gateway.FullName + "/phase"

	// LabelGatewayName is the label for gateway name
	LabelGatewayName = "gateway-name"

	// LabelGatewayEventSourceName is the label for a event source in gateway
	LabelGatewayEventSourceName = "config-name"

	// LabelGatewayEventSourceID is the label for gateway configuration ID
	LabelGatewayEventSourceID = "event-source-id"

	// AnnotationGatewayResourceSpecHashName is the annotation of a gateway resource spec hash
	AnnotationGatewayResourceSpecHashName = gateway.FullName + "/resource-spec-hash"

	// Server Connection Timeout, 10 seconds
	ServerConnTimeout = 10

	LabelArgoEventsGatewayVersion = "argo-events-gateway-version"
)

// Gateway client constants
const (
	// EnvVarGatewayEventSourceConfigMap is used to get map containing event sources to run in a gateway
	EnvVarGatewayEventSourceConfigMap = "GATEWAY_EVENT_SOURCE_CONFIG_MAP"
)

// Gateway server constants
const (
	EnvVarGatewayServerPort = "GATEWAY_SERVER_PORT"
)

// CloudEvents constants
const (
	// CloudEventsVersion is the version of the CloudEvents spec targeted+
	// by this library.
	CloudEventsVersion = "0.1"
)
