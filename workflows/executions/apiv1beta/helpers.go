End
Termination 
Void
Delete
Stop 















































































































End
Termination 
Stop
Delete









































func executeHTTPRequest(ctx context.Context, client *http.Client, req *http.Request, logger *slog.Logger, body []byte, rpc string) ([]byte, error) {
	buf, _, err := executeHTTPRequestWithResponse(ctx, client, req, logger, body, rpc)
	return buf, err
}

func executeStreamingHTTPRequest(ctx context.Context, client *http.Client, req *http.Request, logger *slog.Logger, body []byte, rpc string) (*http.Response, error) {
	logger.DebugContext(ctx, "api request", "serviceName", serviceName, "rpcName", rpc, "request", internallog.HTTPRequest(req, body))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	logger.DebugContext(ctx, "api response", "serviceName", serviceName, "rpcName", rpc, "response", internallog.HTTPResponse(resp, nil))
	if err = googleapi.CheckResponse(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func executeRPC[I proto.Message, O proto.Message](ctx context.Context, fn func(context.Context, I, ...grpc.CallOption) (O, error), req I, opts []grpc.CallOption, logger *slog.Logger, rpc string) (O, error) {
	var zero O
	logger.DebugContext(ctx, "api request", "serviceName", serviceName, "rpcName", rpc, "request", grpclog.ProtoMessageRequest(ctx, req))
	resp, err := fn(ctx, req, opts...)
	if err != nil {
		return zero, err
	}
	logger.DebugContext(ctx, "api response", "serviceName", serviceName, "rpcName", rpc, "response", grpclog.ProtoMessageResponse(resp))
	return resp, err
}
