End
Cancel
Delete
Stop
Void 
End




































































































End 
Delete
Stop
Cancel 
End
Delete













































		ListInputs: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		CreateEvent: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListEvents: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetEvent: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteEvent: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListClips:       []gax.CallOption{},
		GetClip:         []gax.CallOption{},
		CreateClip:      []gax.CallOption{},
		DeleteClip:      []gax.CallOption{},
		CreateAsset:     []gax.CallOption{},
		DeleteAsset:     []gax.CallOption{},
		GetAsset:        []gax.CallOption{},
		ListAssets:      []gax.CallOption{},
		GetPool:         []gax.CallOption{},
		UpdatePool:      []gax.CallOption{},
		GetLocation:     []gax.CallOption{},
		ListLocations:   []gax.CallOption{},
		CancelOperation: []gax.CallOption{},
		DeleteOperation: []gax.CallOption{},
		GetOperation:    []gax.CallOption{},
		ListOperations:  []gax.CallOption{},
	}
}

func defaultRESTCallOptions() *CallOptions {
	return &CallOptions{
		CreateChannel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListChannels: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetChannel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		DeleteChannel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateChannel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		StartChannel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		StopChannel: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		CreateInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListInputs: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		DeleteInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateInput: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		CreateEvent: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListEvents: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetEvent: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		DeleteEvent: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListClips:       []gax.CallOption{},
		GetClip:         []gax.CallOption{},
		CreateClip:      []gax.CallOption{},
		DeleteClip:      []gax.CallOption{},
		CreateAsset:     []gax.CallOption{},
		DeleteAsset:     []gax.CallOption{},
		GetAsset:        []gax.CallOption{},
		ListAssets:      []gax.CallOption{},
		GetPool:         []gax.CallOption{},
		UpdatePool:      []gax.CallOption{},
		GetLocation:     []gax.CallOption{},
		ListLocations:   []gax.CallOption{},
		CancelOperation: []gax.CallOption{},
		DeleteOperation: []gax.CallOption{},
		GetOperation:    []gax.CallOption{},
		ListOperations:  []gax.CallOption{},
	}
}

// internalClient is an interface that defines the methods available from Live Stream API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateChannel(context.Context, *livestreampb.CreateChannelRequest, ...gax.CallOption) (*CreateChannelOperation, error)
	CreateChannelOperation(name string) *CreateChannelOperation
	ListChannels(context.Context, *livestreampb.ListChannelsRequest, ...gax.CallOption) *ChannelIterator
	GetChannel(context.Context, *livestreampb.GetChannelRequest, ...gax.CallOption) (*livestreampb.Channel, error)
	DeleteChannel(context.Context, *livestreampb.DeleteChannelRequest, ...gax.CallOption) (*DeleteChannelOperation, error)
	DeleteChannelOperation(name string) *DeleteChannelOperation
	UpdateChannel(context.Context, *livestreampb.UpdateChannelRequest, ...gax.CallOption) (*UpdateChannelOperation, error)
	UpdateChannelOperation(name string) *UpdateChannelOperation
	StartChannel(context.Context, *livestreampb.StartChannelRequest, ...gax.CallOption) (*StartChannelOperation, error)
	StartChannelOperation(name string) *StartChannelOperation
	StopChannel(context.Context, *livestreampb.StopChannelRequest, ...gax.CallOption) (*StopChannelOperation, error)
	StopChannelOperation(name string) *StopChannelOperation
	CreateInput(context.Context, *livestreampb.CreateInputRequest, ...gax.CallOption) (*CreateInputOperation, error)
	CreateInputOperation(name string) *CreateInputOperation
	ListInputs(context.Context, *livestreampb.ListInputsRequest, ...gax.CallOption) *InputIterator
	GetInput(context.Context, *livestreampb.GetInputRequest, ...gax.CallOption) (*livestreampb.Input, error)
	DeleteInput(context.Context, *livestreampb.DeleteInputRequest, ...gax.CallOption) (*DeleteInputOperation, error)
	DeleteInputOperation(name string) *DeleteInputOperation
	UpdateInput(context.Context, *livestreampb.UpdateInputRequest, ...gax.CallOption) (*UpdateInputOperation, error)
	UpdateInputOperation(name string) *UpdateInputOperation
	CreateEvent(context.Context, *livestreampb.CreateEventRequest, ...gax.CallOption) (*livestreampb.Event, error)
	ListEvents(context.Context, *livestreampb.ListEventsRequest, ...gax.CallOption) *EventIterator
	GetEvent(context.Context, *livestreampb.GetEventRequest, ...gax.CallOption) (*livestreampb.Event, error)
	DeleteEvent(context.Context, *livestreampb.DeleteEventRequest, ...gax.CallOption) error
	ListClips(context.Context, *livestreampb.ListClipsRequest, ...gax.CallOption) *ClipIterator
	GetClip(context.Context, *livestreampb.GetClipRequest, ...gax.CallOption) (*livestreampb.Clip, error)
	CreateClip(context.Context, *livestreampb.CreateClipRequest, ...gax.CallOption) (*CreateClipOperation, error)
	CreateClipOperation(name string) *CreateClipOperation
	DeleteClip(context.Context, *livestreampb.DeleteClipRequest, ...gax.CallOption) (*DeleteClipOperation, error)
	DeleteClipOperation(name string) *DeleteClipOperation
	CreateAsset(context.Context, *livestreampb.CreateAssetRequest, ...gax.CallOption) (*CreateAssetOperation, error)
	CreateAssetOperation(name string) *CreateAssetOperation
	DeleteAsset(context.Context, *livestreampb.DeleteAssetRequest, ...gax.CallOption) (*DeleteAssetOperation, error)
	DeleteAssetOperation(name string) *DeleteAssetOperation
	GetAsset(context.Context, *livestreampb.GetAssetRequest, ...gax.CallOption) (*livestreampb.Asset, error)
	ListAssets(context.Context, *livestreampb.ListAssetsRequest, ...gax.CallOption) *AssetIterator
	GetPool(context.Context, *livestreampb.GetPoolRequest, ...gax.CallOption) (*livestreampb.Pool, error)
	UpdatePool(context.Context, *livestreampb.UpdatePoolRequest, ...gax.CallOption) (*UpdatePoolOperation, error)
	UpdatePoolOperation(name string) *UpdatePoolOperation
	GetLocation(context.Context, *locationpb.GetLocationRequest, ...gax.CallOption) (*locationpb.Location, error)
	ListLocations(context.Context, *locationpb.ListLocationsRequest, ...gax.CallOption) *LocationIterator
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest, ...gax.CallOption) error
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// Client is a client for interacting with Live Stream API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Using Live Stream API, you can generate live streams in the various
// renditions and streaming formats. The streaming format include HTTP Live
// Streaming (HLS) and Dynamic Adaptive Streaming over HTTP (DASH). You can send
// a source stream in the various ways, including Real-Time Messaging
// Protocol (RTMP) and Secure Reliable Transport (SRT).
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateChannel creates a channel with the provided unique ID in the specified
// region.
func (c *Client) CreateChannel(ctx context.Context, req *livestreampb.CreateChannelRequest, opts ...gax.CallOption) (*CreateChannelOperation, error) {
	return c.internalClient.CreateChannel(ctx, req, opts...)
}

// CreateChannelOperation returns a new CreateChannelOperation from a given name.
// The name must be that of a previously created CreateChannelOperation, possibly from a different process.
func (c *Client) CreateChannelOperation(name string) *CreateChannelOperation {
	return c.internalClient.CreateChannelOperation(name)
}

// ListChannels returns a list of all channels in the specified region.
func (c *Client) ListChannels(ctx context.Context, req *livestreampb.ListChannelsRequest, opts ...gax.CallOption) *ChannelIterator {
	return c.internalClient.ListChannels(ctx, req, opts...)
}

// GetChannel returns the specified channel.
func (c *Client) GetChannel(ctx context.Context, req *livestreampb.GetChannelRequest, opts ...gax.CallOption) (*livestreampb.Channel, error) {
	return c.internalClient.GetChannel(ctx, req, opts...)
}

// DeleteChannel deletes the specified channel.
func (c *Client) DeleteChannel(ctx context.Context, req *livestreampb.DeleteChannelRequest, opts ...gax.CallOption) (*DeleteChannelOperation, error) {
	return c.internalClient.DeleteChannel(ctx, req, opts...)
}

// DeleteChannelOperation returns a new DeleteChannelOperation from a given name.
// The name must be that of a previously created DeleteChannelOperation, possibly from a different process.
func (c *Client) DeleteChannelOperation(name string) *DeleteChannelOperation {
	return c.internalClient.DeleteChannelOperation(name)
}

// UpdateChannel updates the specified channel.
func (c *Client) UpdateChannel(ctx context.Context, req *livestreampb.UpdateChannelRequest, opts ...gax.CallOption) (*UpdateChannelOperation, error) {
	return c.internalClient.UpdateChannel(ctx, req, opts...)
}

// UpdateChannelOperation returns a new UpdateChannelOperation from a given name.
// The name must be that of a previously created UpdateChannelOperation, possibly from a different process.
func (c *Client) UpdateChannelOperation(name string) *UpdateChannelOperation {
	return c.internalClient.UpdateChannelOperation(name)
}

// StartChannel starts the specified channel. Part of the video pipeline will be created
// only when the StartChannel request is received by the server.
func (c *Client) StartChannel(ctx context.Context, req *livestreampb.StartChannelRequest, opts ...gax.CallOption) (*StartChannelOperation, error) {
	return c.internalClient.StartChannel(ctx, req, opts...)
}

// StartChannelOperation returns a new StartChannelOperation from a given name.
// The name must be that of a previously created StartChannelOperation, possibly from a different process.
func (c *Client) StartChannelOperation(name string) *StartChannelOperation {
	return c.internalClient.StartChannelOperation(name)
}

// StopChannel stops the specified channel. Part of the video pipeline will be released
// when the StopChannel request is received by the server.
func (c *Client) StopChannel(ctx context.Context, req *livestreampb.StopChannelRequest, opts ...gax.CallOption) (*StopChannelOperation, error) {
	return c.internalClient.StopChannel(ctx, req, opts...)
}

// StopChannelOperation returns a new StopChannelOperation from a given name.
// The name must be that of a previously created StopChannelOperation, possibly from a different process.
func (c *Client) StopChannelOperation(name string) *StopChannelOperation {
	return c.internalClient.StopChannelOperation(name)
}

// CreateInput creates an input with the provided unique ID in the specified region.
func (c *Client) CreateInput(ctx context.Context, req *livestreampb.CreateInputRequest, opts ...gax.CallOption) (*CreateInputOperation, error) {
	return c.internalClient.CreateInput(ctx, req, opts...)
}

// CreateInputOperation returns a new CreateInputOperation from a given name.
// The name must be that of a previously created CreateInputOperation, possibly from a different process.
func (c *Client) CreateInputOperation(name string) *CreateInputOperation {
	return c.internalClient.CreateInputOperation(name)
}

// ListInputs returns a list of all inputs in the specified region.
func (c *Client) ListInputs(ctx context.Context, req *livestreampb.ListInputsRequest, opts ...gax.CallOption) *InputIterator {
	return c.internalClient.ListInputs(ctx, req, opts...)
}

// GetInput returns the specified input.
func (c *Client) GetInput(ctx context.Context, req *livestreampb.GetInputRequest, opts ...gax.CallOption) (*livestreampb.Input, error) {
	return c.internalClient.GetInput(ctx, req, opts...)
}

// DeleteInput deletes the specified input.
func (c *Client) DeleteInput(ctx context.Context, req *livestreampb.DeleteInputRequest, opts ...gax.CallOption) (*DeleteInputOperation, error) {
	return c.internalClient.DeleteInput(ctx, req, opts...)
}

// DeleteInputOperation returns a new DeleteInputOperation from a given name.
// The name must be that of a previously created DeleteInputOperation, possibly from a different process.
func (c *Client) DeleteInputOperation(name string) *DeleteInputOperation {
	return c.internalClient.DeleteInputOperation(name)
}

// UpdateInput updates the specified input.
func (c *Client) UpdateInput(ctx context.Context, req *livestreampb.UpdateInputRequest, opts ...gax.CallOption) (*UpdateInputOperation, error) {
	return c.internalClient.UpdateInput(ctx, req, opts...)
}

// UpdateInputOperation returns a new UpdateInputOperation from a given name.
// The name must be that of a previously created UpdateInputOperation, possibly from a different process.
func (c *Client) UpdateInputOperation(name string) *UpdateInputOperation {
	return c.internalClient.UpdateInputOperation(name)
}

// CreateEvent creates an event with the provided unique ID in the specified channel.
func (c *Client) CreateEvent(ctx context.Context, req *livestreampb.CreateEventRequest, opts ...gax.CallOption) (*livestreampb.Event, error) {
	return c.internalClient.CreateEvent(ctx, req, opts...)
}

// ListEvents returns a list of all events in the specified channel.
func (c *Client) ListEvents(ctx context.Context, req *livestreampb.ListEventsRequest, opts ...gax.CallOption) *EventIterator {
	return c.internalClient.ListEvents(ctx, req, opts...)
}

// GetEvent returns the specified event.
func (c *Client) GetEvent(ctx context.Context, req *livestreampb.GetEventRequest, opts ...gax.CallOption) (*livestreampb.Event, error) {
	return c.internalClient.GetEvent(ctx, req, opts...)
}

// DeleteEvent deletes the specified event.
func (c *Client) DeleteEvent(ctx context.Context, req *livestreampb.DeleteEventRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteEvent(ctx, req, opts...)
}

// ListClips returns a list of all clips in the specified channel.
func (c *Client) ListClips(ctx context.Context, req *livestreampb.ListClipsRequest, opts ...gax.CallOption) *ClipIterator {
	return c.internalClient.ListClips(ctx, req, opts...)
}

// GetClip returns the specified clip.
func (c *Client) GetClip(ctx context.Context, req *livestreampb.GetClipRequest, opts ...gax.CallOption) (*livestreampb.Clip, error) {
	return c.internalClient.GetClip(ctx, req, opts...)
}

// CreateClip creates a clip with the provided clip ID in the specified channel.
func (c *Client) CreateClip(ctx context.Context, req *livestreampb.CreateClipRequest, opts ...gax.CallOption) (*CreateClipOperation, error) {
	return c.internalClient.CreateClip(ctx, req, opts...)
}

// CreateClipOperation returns a new CreateClipOperation from a given name.
// The name must be that of a previously created CreateClipOperation, possibly from a different process.
func (c *Client) CreateClipOperation(name string) *CreateClipOperation {
	return c.internalClient.CreateClipOperation(name)
}

// DeleteClip deletes the specified clip job resource. This method only deletes the clip
// job and does not delete the VOD clip stored in the GCS.
func (c *Client) DeleteClip(ctx context.Context, req *livestreampb.DeleteClipRequest, opts ...gax.CallOption) (*DeleteClipOperation, error) {
	return c.internalClient.DeleteClip(ctx, req, opts...)
}

// DeleteClipOperation returns a new DeleteClipOperation from a given name.
// The name must be that of a previously created DeleteClipOperation, possibly from a different process.
func (c *Client) DeleteClipOperation(name string) *DeleteClipOperation {
	return c.internalClient.DeleteClipOperation(name)
}

// CreateAsset creates a Asset with the provided unique ID in the specified
// region.
func (c *Client) CreateAsset(ctx context.Context, req *livestreampb.CreateAssetRequest, opts ...gax.CallOption) (*CreateAssetOperation, error) {
	return c.internalClient.CreateAsset(ctx, req, opts...)
}

// CreateAssetOperation returns a new CreateAssetOperation from a given name.
// The name must be that of a previously created CreateAssetOperation, possibly from a different process.
func (c *Client) CreateAssetOperation(name string) *CreateAssetOperation {
	return c.internalClient.CreateAssetOperation(name)
}

// DeleteAsset deletes the specified asset if it is not used.
func (c *Client) DeleteAsset(ctx context.Context, req *livestreampb.DeleteAssetRequest, opts ...gax.CallOption) (*DeleteAssetOperation, error) {
	return c.internalClient.DeleteAsset(ctx, req, opts...)
}

// DeleteAssetOperation returns a new DeleteAssetOperation from a given name.
// The name must be that of a previously created DeleteAssetOperation, possibly from a different process.
func (c *Client) DeleteAssetOperation(name string) *DeleteAssetOperation {
	return c.internalClient.DeleteAssetOperation(name)
}

// GetAsset returns the specified asset.
func (c *Client) GetAsset(ctx context.Context, req *livestreampb.GetAssetRequest, opts ...gax.CallOption) (*livestreampb.Asset, error) {
	return c.internalClient.GetAsset(ctx, req, opts...)
}

// ListAssets returns a list of all assets in the specified region.
func (c *Client) ListAssets(ctx context.Context, req *livestreampb.ListAssetsRequest, opts ...gax.CallOption) *AssetIterator {
	return c.internalClient.ListAssets(ctx, req, opts...)
}

// GetPool returns the specified pool.
func (c *Client) GetPool(ctx context.Context, req *livestreampb.GetPoolRequest, opts ...gax.CallOption) (*livestreampb.Pool, error) {
	return c.internalClient.GetPool(ctx, req, opts...)
}

// UpdatePool updates the specified pool.
func (c *Client) UpdatePool(ctx context.Context, req *livestreampb.UpdatePoolRequest, opts ...gax.CallOption) (*UpdatePoolOperation, error) {
	return c.internalClient.UpdatePool(ctx, req, opts...)
}

// UpdatePoolOperation returns a new UpdatePoolOperation from a given name.
// The name must be that of a previously created UpdatePoolOperation, possibly from a different process.
func (c *Client) UpdatePoolOperation(name string) *UpdatePoolOperation {
	return c.internalClient.UpdatePoolOperation(name)
}

// GetLocation gets information about a location.
func (c *Client) GetLocation(ctx context.Context, req *locationpb.GetLocationRequest, opts ...gax.CallOption) (*locationpb.Location, error) {
	return c.internalClient.GetLocation(ctx, req, opts...)
}

// ListLocations lists information about the supported locations for this service.
func (c *Client) ListLocations(ctx context.Context, req *locationpb.ListLocationsRequest, opts ...gax.CallOption) *LocationIterator {
	return c.internalClient.ListLocations(ctx, req, opts...)
}

// CancelOperation is a utility method from google.longrunning.Operations.
func (c *Client) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.CancelOperation(ctx, req, opts...)
}

// DeleteOperation is a utility method from google.longrunning.Operations.
func (c *Client) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteOperation(ctx, req, opts...)
}

// GetOperation 

// NewClient creates a new livestream service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Using Live Stream API, you can generate live 
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		client:           livestreampb.NewLivestreamServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		logger:           internaloption.GetLogger(opts),
		operationsClient: longrunningpb.NewOperationsClient(connPool),
		locationsClient:  locationpb.NewLocationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type restClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	logger *slog.Logger
}

// NewRESTClient creates a new livestream service rest client.
//
// Using Live Stream API, you can generate live streams in the various
// renditions and streaming formats. The streaming format include HTTP Live
// Streaming (HLS) and Dynamic Adaptive Streaming over HTTP (DASH). You can send
// a source stream in the various ways, including Real-Time Messaging
// Protocol (RTMP) and Secure Reliable Transport (SRT).
func NewRESTClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := append(defaultRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRESTCallOptions()
	c := &restClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://livestream.googleapis.com"),
		internaloption.WithDefaultEndpointTemplate("https://livestream.UNIVERSE_DOMAIN"),
		internaloption.WithDefaultMTLSEndpoint("https://livestream.mtls.googleapis.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://livestream.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *restClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *restClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *restClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *gRPCClient) CreateChannel(ctx context.Context, req *livestreampb.CreateChannelRequest, opts ...gax.CallOption) (*CreateChannelOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateChannel[0:len((*c.CallOptions).CreateChannel):len((*c.CallOptions).CreateChannel)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.CreateChannel, req, settings.GRPC, c.logger, "CreateChannel")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateChannelOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) ListChannels(ctx context.Context, req *livestreampb.ListChannelsRequest, opts ...gax.CallOption) *ChannelIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListChannels[0:len((*c.CallOptions).ListChannels):len((*c.CallOptions).ListChannels)], opts...)
	it := &ChannelIterator{}
	req = proto.Clone(req).(*livestreampb.ListChannelsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*livestreampb.Channel, string, error) {
		resp := &livestreampb.ListChannelsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.client.ListChannels, req, settings.GRPC, c.logger, "ListChannels")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetChannels(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) GetChannel(ctx context.Context, req *livestreampb.GetChannelRequest, opts ...gax.CallOption) (*livestreampb.Channel, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetChannel[0:len((*c.CallOptions).GetChannel):len((*c.CallOptions).GetChannel)], opts...)
	var resp *livestreampb.Channel
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.GetChannel, req, settings.GRPC, c.logger, "GetChannel")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteChannel(ctx context.Context, req *livestreampb.DeleteChannelRequest, opts ...gax.CallOption) (*DeleteChannelOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteChannel[0:len((*c.CallOptions).DeleteChannel):len((*c.CallOptions).DeleteChannel)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.DeleteChannel, req, settings.GRPC, c.logger, "DeleteChannel")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteChannelOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) UpdateChannel(ctx context.Context, req *livestreampb.UpdateChannelRequest, opts ...gax.CallOption) (*UpdateChannelOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "channel.name", url.QueryEscape(req.GetChannel().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateChannel[0:len((*c.CallOptions).UpdateChannel):len((*c.CallOptions).UpdateChannel)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.UpdateChannel, req, settings.GRPC, c.logger, "UpdateChannel")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateChannelOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) StartChannel(ctx context.Context, req *livestreampb.StartChannelRequest, opts ...gax.CallOption) (*StartChannelOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StartChannel[0:len((*c.CallOptions).StartChannel):len((*c.CallOptions).StartChannel)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.StartChannel, req, settings.GRPC, c.logger, "StartChannel")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &StartChannelOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) StopChannel(ctx context.Context, req *livestreampb.StopChannelRequest, opts ...gax.CallOption) (*StopChannelOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StopChannel[0:len((*c.CallOptions).StopChannel):len((*c.CallOptions).StopChannel)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.StopChannel, req, settings.GRPC, c.logger, "StopChannel")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &StopChannelOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) CreateInput(ctx context.Context, req *livestreampb.CreateInputRequest, opts ...gax.CallOption) (*CreateInputOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.append((*c.CallOptions).CreateEvent[0:len((*c.CallOptions).CreateEvent):len((*c.CallOptions).CreateEvent)], opts...)
	var resp *livestreampb.Event
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.CreateEvent, req, settings.GRPC, c.logger, "CreateEvent")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListEvents(ctx context.Context, req *livestreampb.ListEventsRequest, opts ...gax.CallOption) *EventIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%
		resp, err = executeRPC(ctx, c.client.GetEvent, req, settings.GRPC, c.logger, "GetEvent")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteEvent(ctx context.Context, req *livestreampb.DeleteEventRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteEvent[0:len((*c.CallOptions).DeleteEvent):len((*c.CallOptions).DeleteEvent)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.client.DeleteEvent, req, settings.GRPC, c.logger, "DeleteEvent")
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) ListClips(ctx context.Context, req *livestreampb.ListClipsRequest, opts ...gax.CallOption) *ClipIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListClips[0:len((*c.CallOptions).ListClips):len((*c.CallOptions).ListClips)], opts...)
	it := &ClipIterator{}
	req = proto.Clone(req).(*livestreampb.ListClipsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*livestreampb.Clip, string, error) {
		resp := &livestreampb.ListClipsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.client.ListClips, req, settings.GRPC, c.logger, "ListClips")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetClips(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) GetClip(ctx context.Context, req *livestreampb.GetClipRequest, opts ...gax.CallOption) (*livestreampb.Clip, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetClip[0:len((*c.CallOptions).GetClip):len((*c.CallOptions).GetClip)], opts...)
	var resp *livestreampb.Clip
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.GetClip, req, settings.GRPC, c.logger, "GetClip")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateClip(ctx context.Context, req *livestreampb.CreateClipRequest, opts ...gax.CallOption) (*CreateClipOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateClip[0:len((*c.CallOptions).CreateClip):len((*c.CallOptions).CreateClip)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.CreateClip, req, settings.GRPC, c.logger, "CreateClip")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateClipOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) DeleteClip(ctx context.Context, req *livestreampb.DeleteClipRequest, opts ...gax.CallOption) (*DeleteClipOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteClip[0:len((*c.CallOptions).DeleteClip):len((*c.CallOptions).DeleteClip)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.DeleteClip, req, settings.GRPC, c.logger, "DeleteClip")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteClipOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) CreateAsset(ctx context.Context, req *livestreampb.CreateAssetRequest, opts ...gax.CallOption) (*CreateAssetOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateAsset[0:len((*c.CallOptions).CreateAsset):len((*c.CallOptions).CreateAsset)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.CreateAsset, req, settings.GRPC, c.logger, "CreateAsset")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateAssetOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) DeleteAsset(ctx context.Context, req *livestreampb.DeleteAssetRequest, opts ...gax.CallOption) (*DeleteAssetOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteAsset[0:len((*c.CallOptions).DeleteAsset):len((*c.CallOptions).DeleteAsset)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.DeleteAsset, req, settings.GRPC, c.logger, "DeleteAsset")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteAssetOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *gRPCClient) GetAsset(ctx context.Context, req *livestreampb.GetAssetRequest, opts ...gax.CallOption) (*livestreampb.Asset, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetAsset[0:len((*c.CallOptions).GetAsset):len((*c.CallOptions).GetAsset)], opts...)
	var resp *livestreampb.Asset
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.GetAsset, req, settings.GRPC, c.logger, "GetAsset")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListAssets(ctx context.Context, req *livestreampb.ListAssetsRequest, opts ...gax.CallOption) *AssetIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListAssets[0:len((*c.CallOptions).ListAssets):len((*c.CallOptions).ListAssets)], opts...)
	it := &AssetIterator{}
	req = proto.Clone(req).(*livestreampb.ListAssetsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*livestreampb.Asset, string, error) {
		resp := &livestreampb.ListAssetsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.client.ListAssets, req, settings.GRPC, c.logger, "ListAssets")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAssets(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) GetPool(ctx context.Context, req *livestreampb.GetPoolRequest, opts ...gax.CallOption) (*livestreampb.Pool, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetPool[0:len((*c.CallOptions).GetPool):len((*c.CallOptions).GetPool)], opts...)
	var resp *livestreampb.Pool
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.client.GetPool, req, settings.GRPC, c.logger, "GetPool")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdatePool(ctx context.Context, req *livestreampb.UpdatePoolRequest, opts ...gax.CallOption) (*UpdatePoolOperation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "pool.name", url.QueryEscape(req.GetPool().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(locationpb.Location, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetLocation[0:len((*c.CallOptions).GetLocation):len((*c.CallOptions).GetLocation)], opts...)
	var resp *locationpb.Location
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.locationsClient.GetLocation, req, settings.GRPC, c.logger, "GetLocation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListLocations[0:len((*c.CallOptions).ListLocations):len((*c.CallOptions).ListLocations)], opts...)
	it := &LocationIterator{}
	req = proto.Clone(req).(*locationpb.ListLocationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*locationpb.Location, string, error) {
		resp := &locationpb.ListLocationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.locationsClient.ListLocations, CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.operationsClient.CancelOperation, req, settings.GRPC, c.logger, "CancelOperation")
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.operationsClient.DeleteOperation, req, settings.GRPC, c.logger, "DeleteOperation")
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.operationsClient.GetOperation, req, settings.GRPC, c.logger, "GetOperation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.operationsClient.ListOperations, req, settings.GRPC, c.logger, "ListOperations")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// CreateChannel creates a channel with the provided unique ID in the specified
// region.
func (c *restClient) CreateChannel(ctx context.Context, req *livestreampb.CreateChannelRequest, opts ...gax.CallOption) (*CreateChannelOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetChannel()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/channels", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	params.Add("channelId", fmt.Sprintf("%v", req.GetChannelId()))
	if req.GetRequestId() != "" {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "CreateChannel")
		if err != nil {
			return err
		}
		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &CreateChannelOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// ListChannels returns a list of all channels in the specified region.
func (c *restClient) ListChannels(ctx context.Context, req *livestreampb.ListChannelsRequest, opts ...gax.CallOption) *ChannelIterator {
	it := &ChannelIterator{}
	req = proto.Clone(req).(*livestreampb.ListChannelsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*livestreampb.Channel, string, error) {
		resp := &livestreampb.ListChannelsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1/%v/channels", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetOrderBy() != "" {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListChannels")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetChannels(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// GetChannel returns the specified channel.
func (c *restClient) GetChannel(ctx context.Context, req *livestreampb.GetChannelRequest, opts ...gax.CallOption) (*livestreampb.Channel, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetChannel[0:len((*c.CallOptions).GetChannel):len((*c.CallOptions).GetChannel)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &livestreampb.Channel{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetChannel")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// DeleteChannel deletes the specified channel.
func (c *restClient) DeleteChannel(ctx context.Context, req *livestreampb.DeleteChannelRequest, opts ...gax.CallOption) (*DeleteChannelOperation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetForce() {
		params.Add("force", fmt.Sprintf("%v", req.GetForce()))
	}
	if req.GetRequestId() != "" {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "DeleteChannel")
		if err != nil {
			return err
		}
		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &DeleteChannelOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// UpdateChannel updates the specified channel.
func (c *restClient) UpdateChannel(ctx context.Context, req *livestreampb.UpdateChannelRequest, opts ...gax.CallOption) (*UpdateChannelOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetChannel()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetChannel().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetRequestId() != "" {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}
	if req.GetUpdateMask() != nil {
		field, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(field[1:len(field)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "channel.name", url.QueryEscape(req.GetChannel().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateChannel")
		if err != nil {
			return err
		}
		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &UpdateChannelOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// StartChannel starts the specified channel. Part of the video pipeline will be created
// only when the StartChannel request is received by the server.
func (c *restClient) StartChannel(ctx context.Context, req *livestreampb.StartChannelRequest, opts ...gax.CallOption) (*StartChannelOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:start", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "StartChannel")
		if err != nil {
			return err
		}
		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &StartChannelOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// StopChannel stops the specified channel. Part of the video pipeline will be released
// when the StopChannel request is received by the server.
func (c *restClient) StopChannel(ctx context.Context, req *livestreampb.StopChannelRequest, opts ...gax.CallOption) (*StopChannelOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v:stop", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "StopChannel")
		if err != nil {
			return err
		}
		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &StopChannelOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// CreateInput creates an input with the provided unique ID in the specified region.
func (c *restClient) CreateInput(ctx context.Context, req *livestreampb.CreateInputRequest, opts ...gax.CallOption) (*CreateInputOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetInput()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/inputs", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	params.Add("inputId", fmt.Sprintf("%v", req.GetInputId()))
	if req.GetRequestId() != "" {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "CreateInput")
		if err != nil {
			return err
		}
		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &CreateInputOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// ListInputs returns a list of all inputs in the specified region.
func (c *restClient) ListInputs(ctx context.Context, req *livestreampb.ListInputsRequest, opts ...gax.CallOption) *InputIterator {
	it := &InputIterator{}
	req = proto.Clone(req).(*livestreampb.ListInputsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*livestreampb.Input, string, error) {
		resp := &livestreampb.ListInputsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1/%v/inputs", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetOrderBy() != "" {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListInputs")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetInputs(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// GetInput returns the specified input.
func (c *restClient) GetInput(ctx context.Context, req *livestreampb.GetInputRequest, opts ...gax.CallOption) (*livestreampb.Input, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetInput[0:len((*c.CallOptions).GetInput):len((*c.CallOptions).GetInput)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &livestreampb.Input{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetInput")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// DeleteInput deletes the specified input.
func (c *restClient) DeleteInput(ctx context.Context, req *livestreampb.DeleteInputRequest, opts ...gax.CallOption) (*DeleteInputOperation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetRequestId() != "" {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "DeleteInput")
		if err != nil {
			return err
		}
		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &DeleteInputOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// UpdateInput updates the specified input.
func (c *restClient) UpdateInput(ctx context.Context, req *livestreampb.UpdateInputRequest, opts ...gax.CallOption) (*UpdateInputOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetInput()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetInput().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetRequestId() != "" {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}
	if req.GetUpdateMask() != nil {
		field, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(field[1:len(field)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "input.name", url.QueryEscape(req.GetInput().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings 
