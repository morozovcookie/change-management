/*
Change Management

Testing CRQApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_CRQApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CRQApiService CreateCRQ", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CRQApi.CreateCRQ(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CRQApiService GetCRQById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var crqId string

		resp, httpRes, err := apiClient.CRQApi.GetCRQById(context.Background(), crqId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CRQApiService ListCRQ", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CRQApi.ListCRQ(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
