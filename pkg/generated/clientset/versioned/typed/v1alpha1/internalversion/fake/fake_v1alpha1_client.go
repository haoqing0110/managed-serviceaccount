/*
Copyright 2021.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	internalversion "open-cluster-management.io/managed-serviceaccount/pkg/generated/clientset/versioned/typed/v1alpha1/internalversion"
)

type FakeAuthentication struct {
	*testing.Fake
}

func (c *FakeAuthentication) ManagedServiceAccounts(namespace string) internalversion.ManagedServiceAccountInterface {
	return &FakeManagedServiceAccounts{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAuthentication) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
