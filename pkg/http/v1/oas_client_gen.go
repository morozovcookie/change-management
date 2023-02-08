// Code generated by ogen, DO NOT EDIT.

package v1

import (
	"context"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// CreateCRQ invokes createCRQ operation.
//
// To create a new change request you should send a POST request to
// `/api/v1/crq`.
// You should know that change request does not have any unique parameter in
// request that could be used to prevent creation twice.
// If you will send two requests with the same body than you will get two
// change requests with the different unique identifiers.
// As a result you will get relative URL to the created change request in the
// `Location` header with HTTP status code `201 Created`.
//
// POST /api/v1/crq
func (c *Client) CreateCRQ(ctx context.Context, request *CreateCRQReq) (CreateCRQRes, error) {
	res, err := c.sendCreateCRQ(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendCreateCRQ(ctx context.Context, request *CreateCRQReq) (res CreateCRQRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createCRQ"),
	}
	// Validate request before sending.
	if err := func() error {
		if err := request.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return res, errors.Wrap(err, "validate")
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateCRQ",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/api/v1/crq"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateCRQRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateCRQResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateIncident invokes createIncident operation.
//
// To create a new incident you should send a POST request to
// `/api/v1/im`.
// You should know that incident does not have any unique parameter in
// request that could be used to prevent creation twice.
// If you will send two requests with the same body than you will get two
// incidents with the different unique identifiers.
// As a result you will get relative URL to the created incident in the
// `Location` header with HTTP status code `201 Created`.
//
// POST /api/v1/im
func (c *Client) CreateIncident(ctx context.Context, request *CreateIncidentReq) (CreateIncidentRes, error) {
	res, err := c.sendCreateIncident(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendCreateIncident(ctx context.Context, request *CreateIncidentReq) (res CreateIncidentRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createIncident"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "CreateIncident",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/api/v1/im"

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateIncidentRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateIncidentResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetCRQById invokes getCRQById operation.
//
// To retrieve information about specific change request, you should send a GET
// request to `/api/v1/crq/` with specified change request unique identifier.
//
// GET /api/v1/crq/{crqId}
func (c *Client) GetCRQById(ctx context.Context, params GetCRQByIdParams) (GetCRQByIdRes, error) {
	res, err := c.sendGetCRQById(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendGetCRQById(ctx context.Context, params GetCRQByIdParams) (res GetCRQByIdRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getCRQById"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "GetCRQById",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/api/v1/crq/"
	{
		// Encode "crqId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "crqId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			if unwrapped := string(params.CrqId); true {
				return e.EncodeValue(conv.StringToString(unwrapped))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetCRQByIdResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetIncidentById invokes getIncidentById operation.
//
// To retrieve information about specific incident, you should send a GET
// request to `/api/v1/im/` with specified incident unique identifier.
//
// GET /api/v1/im/{incidentId}
func (c *Client) GetIncidentById(ctx context.Context, params GetIncidentByIdParams) (GetIncidentByIdRes, error) {
	res, err := c.sendGetIncidentById(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendGetIncidentById(ctx context.Context, params GetIncidentByIdParams) (res GetIncidentByIdRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getIncidentById"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "GetIncidentById",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/api/v1/im/"
	{
		// Encode "incidentId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "incidentId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			if unwrapped := string(params.IncidentId); true {
				return e.EncodeValue(conv.StringToString(unwrapped))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetIncidentByIdResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListCRQ invokes listCRQ operation.
//
// To list all of the change requests, send a GET request to `/api/v1/crq`.
// The result will be a JSON object with a `data` key. This will be set to an
// array of change requests, each of which will contain the change request
// attributes.
//
// GET /api/v1/crq
func (c *Client) ListCRQ(ctx context.Context, params ListCRQParams) (ListCRQRes, error) {
	res, err := c.sendListCRQ(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendListCRQ(ctx context.Context, params ListCRQParams) (res ListCRQRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listCRQ"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ListCRQ",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/api/v1/crq"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "start" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "start",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Start.Get(); ok {
				if unwrapped := int(val); true {
					return e.EncodeValue(conv.IntToString(unwrapped))
				}
				return nil
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "limit" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Limit.Get(); ok {
				if unwrapped := int(val); true {
					return e.EncodeValue(conv.IntToString(unwrapped))
				}
				return nil
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListCRQResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListIncidents invokes listIncidents operation.
//
// To list all of the incidents, send a GET request to `/api/v1/im`.
// The result will be a JSON object with a `data` key. This will be set to an
// array of incidents, each of which will contain the incident attributes.
//
// GET /api/v1/im
func (c *Client) ListIncidents(ctx context.Context, params ListIncidentsParams) (ListIncidentsRes, error) {
	res, err := c.sendListIncidents(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendListIncidents(ctx context.Context, params ListIncidentsParams) (res ListIncidentsRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listIncidents"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "ListIncidents",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	u.Path += "/api/v1/im"

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "start" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "start",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Start.Get(); ok {
				if unwrapped := int(val); true {
					return e.EncodeValue(conv.IntToString(unwrapped))
				}
				return nil
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "limit" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Limit.Get(); ok {
				if unwrapped := int(val); true {
					return e.EncodeValue(conv.IntToString(unwrapped))
				}
				return nil
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeListIncidentsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}