/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-csi/execution-hook/pkg/apis/executionhook/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ExecutionHookLister helps list ExecutionHooks.
type ExecutionHookLister interface {
	// List lists all ExecutionHooks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ExecutionHook, err error)
	// ExecutionHooks returns an object that can list and get ExecutionHooks.
	ExecutionHooks(namespace string) ExecutionHookNamespaceLister
	ExecutionHookListerExpansion
}

// executionHookLister implements the ExecutionHookLister interface.
type executionHookLister struct {
	indexer cache.Indexer
}

// NewExecutionHookLister returns a new ExecutionHookLister.
func NewExecutionHookLister(indexer cache.Indexer) ExecutionHookLister {
	return &executionHookLister{indexer: indexer}
}

// List lists all ExecutionHooks in the indexer.
func (s *executionHookLister) List(selector labels.Selector) (ret []*v1alpha1.ExecutionHook, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExecutionHook))
	})
	return ret, err
}

// ExecutionHooks returns an object that can list and get ExecutionHooks.
func (s *executionHookLister) ExecutionHooks(namespace string) ExecutionHookNamespaceLister {
	return executionHookNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ExecutionHookNamespaceLister helps list and get ExecutionHooks.
type ExecutionHookNamespaceLister interface {
	// List lists all ExecutionHooks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ExecutionHook, err error)
	// Get retrieves the ExecutionHook from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ExecutionHook, error)
	ExecutionHookNamespaceListerExpansion
}

// executionHookNamespaceLister implements the ExecutionHookNamespaceLister
// interface.
type executionHookNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ExecutionHooks in the indexer for a given namespace.
func (s executionHookNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ExecutionHook, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExecutionHook))
	})
	return ret, err
}

// Get retrieves the ExecutionHook from the indexer for a given namespace and name.
func (s executionHookNamespaceLister) Get(name string) (*v1alpha1.ExecutionHook, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("executionhook"), name)
	}
	return obj.(*v1alpha1.ExecutionHook), nil
}
