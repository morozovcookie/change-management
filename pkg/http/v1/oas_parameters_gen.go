// Code generated by ogen, DO NOT EDIT.

package v1

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// GetCRQByIdParams is parameters of getCRQById operation.
type GetCRQByIdParams struct {
	// Change request unique identifier.
	CrqId ID
}

func unpackGetCRQByIdParams(packed middleware.Parameters) (params GetCRQByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "crqId",
			In:   "path",
		}
		params.CrqId = packed[key].(ID)
	}
	return params
}

func decodeGetCRQByIdParams(args [1]string, r *http.Request) (params GetCRQByIdParams, _ error) {
	// Decode path: crqId.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "crqId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotCrqIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCrqIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CrqId = ID(paramsDotCrqIdVal)
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := params.CrqId.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "crqId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetIncidentByIdParams is parameters of getIncidentById operation.
type GetIncidentByIdParams struct {
	// Incident unique identifier.
	IncidentId ID
}

func unpackGetIncidentByIdParams(packed middleware.Parameters) (params GetIncidentByIdParams) {
	{
		key := middleware.ParameterKey{
			Name: "incidentId",
			In:   "path",
		}
		params.IncidentId = packed[key].(ID)
	}
	return params
}

func decodeGetIncidentByIdParams(args [1]string, r *http.Request) (params GetIncidentByIdParams, _ error) {
	// Decode path: incidentId.
	if err := func() error {
		param, err := url.PathUnescape(args[0])
		if err != nil {
			return errors.Wrap(err, "unescape path")
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "incidentId",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				var paramsDotIncidentIdVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotIncidentIdVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.IncidentId = ID(paramsDotIncidentIdVal)
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := params.IncidentId.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "incidentId",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ListCRQParams is parameters of listCRQ operation.
type ListCRQParams struct {
	// The count of records that should be skipped.
	Start OptStartType
	// The maximum records that should be returned.
	Limit OptLimitType
}

func unpackListCRQParams(packed middleware.Parameters) (params ListCRQParams) {
	{
		key := middleware.ParameterKey{
			Name: "start",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Start = v.(OptStartType)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptLimitType)
		}
	}
	return params
}

func decodeListCRQParams(args [0]string, r *http.Request) (params ListCRQParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: start.
	{
		val := StartType(0)
		params.Start.SetTo(val)
	}
	// Decode query: start.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "start",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStartVal StartType
				if err := func() error {
					var paramsDotStartValVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						paramsDotStartValVal = c
						return nil
					}(); err != nil {
						return err
					}
					paramsDotStartVal = StartType(paramsDotStartValVal)
					return nil
				}(); err != nil {
					return err
				}
				params.Start.SetTo(paramsDotStartVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Start.Set {
					if err := func() error {
						if err := params.Start.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "start",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: limit.
	{
		val := LimitType(20)
		params.Limit.SetTo(val)
	}
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal LimitType
				if err := func() error {
					var paramsDotLimitValVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						paramsDotLimitValVal = c
						return nil
					}(); err != nil {
						return err
					}
					paramsDotLimitVal = LimitType(paramsDotLimitValVal)
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Limit.Set {
					if err := func() error {
						if err := params.Limit.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListIncidentsParams is parameters of listIncidents operation.
type ListIncidentsParams struct {
	// The count of records that should be skipped.
	Start OptStartType
	// The maximum records that should be returned.
	Limit OptLimitType
}

func unpackListIncidentsParams(packed middleware.Parameters) (params ListIncidentsParams) {
	{
		key := middleware.ParameterKey{
			Name: "start",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Start = v.(OptStartType)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptLimitType)
		}
	}
	return params
}

func decodeListIncidentsParams(args [0]string, r *http.Request) (params ListIncidentsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Set default value for query: start.
	{
		val := StartType(0)
		params.Start.SetTo(val)
	}
	// Decode query: start.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "start",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStartVal StartType
				if err := func() error {
					var paramsDotStartValVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						paramsDotStartValVal = c
						return nil
					}(); err != nil {
						return err
					}
					paramsDotStartVal = StartType(paramsDotStartValVal)
					return nil
				}(); err != nil {
					return err
				}
				params.Start.SetTo(paramsDotStartVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Start.Set {
					if err := func() error {
						if err := params.Start.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "start",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: limit.
	{
		val := LimitType(20)
		params.Limit.SetTo(val)
	}
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal LimitType
				if err := func() error {
					var paramsDotLimitValVal int
					if err := func() error {
						val, err := d.DecodeValue()
						if err != nil {
							return err
						}

						c, err := conv.ToInt(val)
						if err != nil {
							return err
						}

						paramsDotLimitValVal = c
						return nil
					}(); err != nil {
						return err
					}
					paramsDotLimitVal = LimitType(paramsDotLimitValVal)
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Limit.Set {
					if err := func() error {
						if err := params.Limit.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}