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

package gateway

import (
	"context"
	"fmt"

	"github.com/argoproj/argo-events/common"
	"github.com/ghodss/yaml"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

// watches configuration for gateway controller
func (c *GatewayController) watchControllerConfigMap(ctx context.Context) (cache.Controller, error) {
	c.log.Info("watching gateway-controller config map updates")
	source := c.newControllerConfigMapWatch()
	_, controller := cache.NewInformer(
		source,
		&apiv1.ConfigMap{},
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				if cm, ok := obj.(*apiv1.ConfigMap); ok {
					c.log.Info("detected EventSource update. updating the gateway-controller config.")
					err := c.updateConfig(cm)
					if err != nil {
						c.log.Error("update of config failed", "err", err)
					}
				}
			},
			UpdateFunc: func(old, new interface{}) {
				if newCm, ok := new.(*apiv1.ConfigMap); ok {
					c.log.Info("detected EventSource update. updating the gateway-controller config.")
					err := c.updateConfig(newCm)
					if err != nil {
						c.log.WithError(err).Error("update of config failed")
					}
				}
			},
		})

	go controller.Run(ctx.Done())
	return controller, nil
}

// creates a new config map watcher
func (c *GatewayController) newControllerConfigMapWatch() *cache.ListWatch {
	x := c.kubeClientset.CoreV1().RESTClient()
	resource := "configmaps"
	name := c.ConfigMap
	fieldSelector := fields.ParseSelectorOrDie(fmt.Sprintf("metadata.name=%s", name))

	listFunc := func(options metav1.ListOptions) (runtime.Object, error) {
		options.FieldSelector = fieldSelector.String()
		req := x.Get().
			Namespace(c.Namespace).
			Resource(resource).
			VersionedParams(&options, metav1.ParameterCodec)
		return req.Do().Get()
	}
	watchFunc := func(options metav1.ListOptions) (watch.Interface, error) {
		options.Watch = true
		options.FieldSelector = fieldSelector.String()
		req := x.Get().
			Namespace(c.Namespace).
			Resource(resource).
			VersionedParams(&options, metav1.ParameterCodec)
		return req.Watch()
	}
	return &cache.ListWatch{ListFunc: listFunc, WatchFunc: watchFunc}
}

// ResyncConfig reloads the gateway-controller config from the configmap
func (c *GatewayController) ResyncConfig(namespace string) error {
	cmClient := c.kubeClientset.CoreV1().ConfigMaps(namespace)
	cm, err := cmClient.Get(c.ConfigMap, metav1.GetOptions{})
	if err != nil {
		return err
	}
	return c.updateConfig(cm)
}

// updates the gateway controller configuration
func (c *GatewayController) updateConfig(cm *apiv1.ConfigMap) error {
	configStr, ok := cm.Data[common.GatewayControllerConfigMapKey]
	if !ok {
		return fmt.Errorf("configMap '%s' does not have key '%s'", c.ConfigMap, common.GatewayControllerConfigMapKey)
	}
	var config GatewayControllerConfig
	err := yaml.Unmarshal([]byte(configStr), &config)
	if err != nil {
		return err
	}
	c.Config = config
	return nil
}
