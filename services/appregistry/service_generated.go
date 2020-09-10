/*
 * Copyright © 2020 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * App Registry
 *
 * With the App Registry service in Splunk Cloud Services, you can create, update, and manage your apps.
 *
 * API version: v1beta2.0 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package appregistry

import (
	"net/http"

	"github.com/splunk/splunk-cloud-sdk-go/services"
	"github.com/splunk/splunk-cloud-sdk-go/util"
)

const serviceCluster = "api"

type Service services.BaseService

// NewService creates a new appregistry service client from the given Config
func NewService(iClient services.IClient) *Service {
	return &Service{Client: iClient}
}

/*
	CreateApp - appregistry service endpoint
	Creates an app.
	Parameters:
		createAppRequest: Creates a new app.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateApp(createAppRequest CreateAppRequest, resp ...*http.Response) (*AppResponseCreateUpdate, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/app-registry/v1beta2/apps`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: createAppRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb AppResponseCreateUpdate
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	CreateSubscription - appregistry service endpoint
	Creates a subscription.
	Parameters:
		appName: Creates a subscription between a tenant and an app.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateSubscription(appName AppName, resp ...*http.Response) error {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/app-registry/v1beta2/subscriptions`, nil)
	if err != nil {
		return err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: appName})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	DeleteApp - appregistry service endpoint
	Removes an app.
	Parameters:
		appName: App name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteApp(appName string, resp ...*http.Response) error {
	pp := struct {
		AppName string
	}{
		AppName: appName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/app-registry/v1beta2/apps/{{.AppName}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	DeleteSubscription - appregistry service endpoint
	Removes a subscription.
	Parameters:
		appName: App name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteSubscription(appName string, resp ...*http.Response) error {
	pp := struct {
		AppName string
	}{
		AppName: appName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/app-registry/v1beta2/subscriptions/{{.AppName}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
	GetApp - appregistry service endpoint
	Returns the metadata of an app.
	Parameters:
		appName: App name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetApp(appName string, resp ...*http.Response) (*AppResponseGetList, error) {
	pp := struct {
		AppName string
	}{
		AppName: appName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/app-registry/v1beta2/apps/{{.AppName}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb AppResponseGetList
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetKeys - appregistry service endpoint
	Returns a list of the public keys used for verifying signed webhook requests.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetKeys(resp ...*http.Response) ([]Key, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/system/app-registry/v1beta2/keys`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb []Key
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	GetSubscription - appregistry service endpoint
	Returns or validates a subscription.
	Parameters:
		appName: App name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetSubscription(appName string, resp ...*http.Response) (*Subscription, error) {
	pp := struct {
		AppName string
	}{
		AppName: appName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/app-registry/v1beta2/subscriptions/{{.AppName}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb Subscription
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListAppSubscriptions - appregistry service endpoint
	Returns the collection of subscriptions to an app.
	Parameters:
		appName: App name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListAppSubscriptions(appName string, resp ...*http.Response) ([]Subscription, error) {
	pp := struct {
		AppName string
	}{
		AppName: appName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/app-registry/v1beta2/apps/{{.AppName}}/subscriptions`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb []Subscription
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListApps - appregistry service endpoint
	Returns a list of apps.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListApps(resp ...*http.Response) ([]AppResponseGetList, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/app-registry/v1beta2/apps`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb []AppResponseGetList
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	ListSubscriptions - appregistry service endpoint
	Returns the tenant subscriptions.
	Parameters:
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListSubscriptions(query *ListSubscriptionsQueryParams, resp ...*http.Response) ([]Subscription, error) {
	values := util.ParseURLParams(query)
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/app-registry/v1beta2/subscriptions`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb []Subscription
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	RotateSecret - appregistry service endpoint
	Rotates the client secret for an app.
	Parameters:
		appName: App name.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) RotateSecret(appName string, resp ...*http.Response) (*AppResponseCreateUpdate, error) {
	pp := struct {
		AppName string
	}{
		AppName: appName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/app-registry/v1beta2/apps/{{.AppName}}/rotate-secret`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb AppResponseCreateUpdate
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	UpdateApp - appregistry service endpoint
	Updates an app.
	Parameters:
		appName: App name.
		updateAppRequest: Updates app contents.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UpdateApp(appName string, updateAppRequest UpdateAppRequest, resp ...*http.Response) (*AppResponseCreateUpdate, error) {
	pp := struct {
		AppName string
	}{
		AppName: appName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/app-registry/v1beta2/apps/{{.AppName}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Put(services.RequestParams{URL: u, Body: updateAppRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb AppResponseCreateUpdate
	err = util.ParseResponse(&rb, response)
	return &rb, err
}
