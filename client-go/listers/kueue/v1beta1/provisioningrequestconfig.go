/*
Copyright 2022 The Kubernetes Authors.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "sigs.k8s.io/kueue/apis/kueue/v1beta1"
)

// ProvisioningRequestConfigLister helps list ProvisioningRequestConfigs.
// All objects returned here must be treated as read-only.
type ProvisioningRequestConfigLister interface {
	// List lists all ProvisioningRequestConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ProvisioningRequestConfig, err error)
	// Get retrieves the ProvisioningRequestConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.ProvisioningRequestConfig, error)
	ProvisioningRequestConfigListerExpansion
}

// provisioningRequestConfigLister implements the ProvisioningRequestConfigLister interface.
type provisioningRequestConfigLister struct {
	indexer cache.Indexer
}

// NewProvisioningRequestConfigLister returns a new ProvisioningRequestConfigLister.
func NewProvisioningRequestConfigLister(indexer cache.Indexer) ProvisioningRequestConfigLister {
	return &provisioningRequestConfigLister{indexer: indexer}
}

// List lists all ProvisioningRequestConfigs in the indexer.
func (s *provisioningRequestConfigLister) List(selector labels.Selector) (ret []*v1beta1.ProvisioningRequestConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ProvisioningRequestConfig))
	})
	return ret, err
}

// Get retrieves the ProvisioningRequestConfig from the index for a given name.
func (s *provisioningRequestConfigLister) Get(name string) (*v1beta1.ProvisioningRequestConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("provisioningrequestconfig"), name)
	}
	return obj.(*v1beta1.ProvisioningRequestConfig), nil
}