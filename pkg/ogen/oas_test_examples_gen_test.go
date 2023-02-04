// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"fmt"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"

	std "encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCRQApplicationJSONBadRequest_EncodeDecode(t *testing.T) {
	var typ CreateCRQApplicationJSONBadRequest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateCRQApplicationJSONBadRequest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateCRQApplicationJSONInternalServerError_EncodeDecode(t *testing.T) {
	var typ CreateCRQApplicationJSONInternalServerError
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateCRQApplicationJSONInternalServerError
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateCRQReq_EncodeDecode(t *testing.T) {
	var typ CreateCRQReq
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateCRQReq
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestCreateCRQReq_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "{\"description\":\"Bump \\u003cservice-name\\u003e version\\n\\nCHANGELOG:\\n  - minor fixes\\n\",\"isAutoClose\":true,\"summary\":\"Bump service \\u003cservice-name\\u003e version\",\"type\":\"autocrq\"}"},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ CreateCRQReq

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				if validateErr, ok := errors.Into[*validate.Error](err); ok {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 CreateCRQReq
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestCreateCRQReqType_EncodeDecode(t *testing.T) {
	var typ CreateCRQReqType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateCRQReqType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestCreateCRQReqType_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "\"crq\""},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ CreateCRQReqType

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				if validateErr, ok := errors.Into[*validate.Error](err); ok {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 CreateCRQReqType
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestCreateIncidentApplicationJSONBadRequest_EncodeDecode(t *testing.T) {
	var typ CreateIncidentApplicationJSONBadRequest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateIncidentApplicationJSONBadRequest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateIncidentApplicationJSONInternalServerError_EncodeDecode(t *testing.T) {
	var typ CreateIncidentApplicationJSONInternalServerError
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateIncidentApplicationJSONInternalServerError
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateIncidentReq_EncodeDecode(t *testing.T) {
	var typ CreateIncidentReq
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateIncidentReq
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestCreateIncidentReq_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "{\"description\":\"Since the release of a new Version the Data in the database is all  scrambled up and the System wont start.\",\"summary\":\"Incident on payroll system\"}"},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ CreateIncidentReq

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				if validateErr, ok := errors.Into[*validate.Error](err); ok {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 CreateIncidentReq
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestCrq_EncodeDecode(t *testing.T) {
	var typ Crq
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Crq
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestCrq_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "{\"_links\":{\"self\":\"https://example.com/api/v1/crq/zd008y7tix97kxtd6y8jddt8ab11m76f\"},\"createdAt\":1675368234,\"description\":\"Bump \\u003cservice-name\\u003e version\\n\\nCHANGELOG:\\n  - minor fixes\\n\",\"id\":\"zd008y7tix97kxtd6y8jddt8ab11m76f\",\"isAutoClose\":true,\"summary\":\"Bump service \\u003cservice-name\\u003e version\",\"type\":\"autocrq\"}"},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ Crq

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				if validateErr, ok := errors.Into[*validate.Error](err); ok {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 Crq
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestCrqType_EncodeDecode(t *testing.T) {
	var typ CrqType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CrqType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestCrqType_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "\"crq\""},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ CrqType

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				if validateErr, ok := errors.Into[*validate.Error](err); ok {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 CrqType
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestError_EncodeDecode(t *testing.T) {
	var typ Error
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Error
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestError_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "{\"code\":\"internal\",\"message\":\"an internal error has occurred\"}"},
		{Input: "{\"code\":\"invalid\",\"message\":\"The request that you did provide has an invalid values\"}"},
		{Input: "{\"code\":\"not_found\",\"message\":\"The resource with specified identifier does not exist\"}"},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ Error

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				if validateErr, ok := errors.Into[*validate.Error](err); ok {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 Error
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestErrorCode_EncodeDecode(t *testing.T) {
	var typ ErrorCode
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ErrorCode
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestErrorCode_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "\"internal\""},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ ErrorCode

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				if validateErr, ok := errors.Into[*validate.Error](err); ok {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 ErrorCode
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestGetCRQByIdApplicationJSONInternalServerError_EncodeDecode(t *testing.T) {
	var typ GetCRQByIdApplicationJSONInternalServerError
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetCRQByIdApplicationJSONInternalServerError
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetCRQByIdApplicationJSONNotFound_EncodeDecode(t *testing.T) {
	var typ GetCRQByIdApplicationJSONNotFound
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetCRQByIdApplicationJSONNotFound
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetIncidentByIdApplicationJSONBadRequest_EncodeDecode(t *testing.T) {
	var typ GetIncidentByIdApplicationJSONBadRequest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetIncidentByIdApplicationJSONBadRequest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestGetIncidentByIdApplicationJSONInternalServerError_EncodeDecode(t *testing.T) {
	var typ GetIncidentByIdApplicationJSONInternalServerError
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 GetIncidentByIdApplicationJSONInternalServerError
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestID_EncodeDecode(t *testing.T) {
	var typ ID
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ID
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestIncident_EncodeDecode(t *testing.T) {
	var typ Incident
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Incident
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}

func TestIncident_Examples(t *testing.T) {

	for i, tc := range []struct {
		Input string
	}{
		{Input: "{\"_links\":{\"self\":\"https://example.com/api/v1/incident/zd008y7tix97kxtd6y8jddt8ab11m76f\"},\"createdAt\":1675368234,\"description\":\"Since the release of a new Version the Data in the database is all  scrambled up and the System wont start.\",\"id\":\"zd008y7tix97kxtd6y8jddt8ab11m76f\",\"summary\":\"Incident on payroll system\"}"},
	} {
		tc := tc
		t.Run(fmt.Sprintf("Test%d", i+1), func(t *testing.T) {
			var typ Incident

			if err := typ.Decode(jx.DecodeStr(tc.Input)); err != nil {
				if validateErr, ok := errors.Into[*validate.Error](err); ok {
					t.Skipf("Validation error: %v", validateErr)
					return
				}
				require.NoErrorf(t, err, "Input: %s", tc.Input)
			}

			e := jx.Encoder{}
			typ.Encode(&e)
			require.True(t, std.Valid(e.Bytes()), "Encoded: %s", e.Bytes())

			var typ2 Incident
			require.NoError(t, typ2.Decode(jx.DecodeBytes(e.Bytes())))
		})
	}
}
func TestLimitType_EncodeDecode(t *testing.T) {
	var typ LimitType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 LimitType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestLink_EncodeDecode(t *testing.T) {
	var typ Link
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Link
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestLinks_EncodeDecode(t *testing.T) {
	var typ Links
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Links
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestListCRQApplicationJSONBadRequest_EncodeDecode(t *testing.T) {
	var typ ListCRQApplicationJSONBadRequest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ListCRQApplicationJSONBadRequest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestListCRQApplicationJSONInternalServerError_EncodeDecode(t *testing.T) {
	var typ ListCRQApplicationJSONInternalServerError
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ListCRQApplicationJSONInternalServerError
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestListCrq_EncodeDecode(t *testing.T) {
	var typ ListCrq
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ListCrq
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestListIncidents_EncodeDecode(t *testing.T) {
	var typ ListIncidents
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ListIncidents
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestListIncidentsApplicationJSONBadRequest_EncodeDecode(t *testing.T) {
	var typ ListIncidentsApplicationJSONBadRequest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ListIncidentsApplicationJSONBadRequest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestListIncidentsApplicationJSONInternalServerError_EncodeDecode(t *testing.T) {
	var typ ListIncidentsApplicationJSONInternalServerError
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ListIncidentsApplicationJSONInternalServerError
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSelfLink_EncodeDecode(t *testing.T) {
	var typ SelfLink
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SelfLink
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestStartType_EncodeDecode(t *testing.T) {
	var typ StartType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 StartType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestTimestamp_EncodeDecode(t *testing.T) {
	var typ Timestamp
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Timestamp
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
