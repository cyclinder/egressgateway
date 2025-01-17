// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/spidernet-io/egressgateway/pkg/k8s/apis/egressgateway.spidernet.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EgressGatewayLister helps list EgressGateways.
// All objects returned here must be treated as read-only.
type EgressGatewayLister interface {
	// List lists all EgressGateways in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EgressGateway, err error)
	// Get retrieves the EgressGateway from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.EgressGateway, error)
	EgressGatewayListerExpansion
}

// egressGatewayLister implements the EgressGatewayLister interface.
type egressGatewayLister struct {
	indexer cache.Indexer
}

// NewEgressGatewayLister returns a new EgressGatewayLister.
func NewEgressGatewayLister(indexer cache.Indexer) EgressGatewayLister {
	return &egressGatewayLister{indexer: indexer}
}

// List lists all EgressGateways in the indexer.
func (s *egressGatewayLister) List(selector labels.Selector) (ret []*v1.EgressGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EgressGateway))
	})
	return ret, err
}

// Get retrieves the EgressGateway from the index for a given name.
func (s *egressGatewayLister) Get(name string) (*v1.EgressGateway, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("egressgateway"), name)
	}
	return obj.(*v1.EgressGateway), nil
}
