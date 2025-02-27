delete 
end











































































	file_google_cloud_billing_budgets_v1beta1_budget_service_proto_rawDescOnce sync.Once
	file_google_cloud_billing_budgets_v1beta1_budget_service_proto_rawDescData = file_google_cloud_billing_budgets_v1beta1_budget_service_proto_rawDesc
)

func file_google_cloud_billing_budgets_v1beta1_budget_service_proto_rawDescGZIP() []byte {
	file_google_cloud_billing_budgets_v1beta1_budget_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_billing_budgets_v1beta1_budget_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_billing_budgets_v1beta1_budget_service_proto_rawDescData)
	})
	return file_google_cloud_billing_budgets_v1beta1_budget_service_proto_rawDescData
}

var file_google_cloud_billing_budgets_v1beta1_budget_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_billing_budgets_v1beta1_budget_service_proto_goTypes = []any{
	(*CreateBudgetRequest)(nil),   // 0: google.cloud.billing.budgets.v1beta1.CreateBudgetRequest
	(*UpdateBudgetRequest)(nil),   // 1: google.cloud.billing.budgets.v1beta1.UpdateBudgetRequest
	(*GetBudgetRequest)(nil),      // 2: google.cloud.billing.budgets.v1beta1.GetBudgetRequest
	(*ListBudgetsRequest)(nil),    // 3: google.cloud.billing.budgets.v1beta1.ListBudgetsRequest
	(*ListBudgetsResponse)(nil),   // 4: google.cloud.billing.budgets.v1beta1.ListBudgetsResponse
	(*DeleteBudgetRequest)(nil),   // 5: google.cloud.billing.budgets.v1beta1.DeleteBudgetRequest
	(*Budget)(nil),                // 6: google.cloud.billing.budgets.v1beta1.Budget
	(*fieldmaskpb.FieldMask)(nil), // 7: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_google_cloud_billing_budgets_v1beta1_budget_service_proto_depIdxs = []int32{
	6, // 0: google.cloud.billing.budgets.v1beta1.CreateBudgetRequest.budget:type_name -> google.cloud.billing.budgets.v1beta1.Budget
	6, // 1: google.cloud.billing.budgets.v1beta1.UpdateBudgetRequest.budget:type_name -> google.cloud.billing.budgets.v1beta1.Budget
	7, // 2: google.cloud.billing.budgets.v1beta1.UpdateBudgetRequest.update_mask:type_name -> google.protobuf.FieldMask
	6, // 3: google.cloud.billing.budgets.v1beta1.ListBudgetsResponse.budgets:type_name -> google.cloud.billing.budgets.v1beta1.Budget
	0, // 4: google.cloud.billing.budgets.v1beta1.BudgetService.CreateBudget:input_type -> google.cloud.billing.budgets.v1beta1.CreateBudgetRequest
	1, // 5: google.cloud.billing.budgets.v1beta1.BudgetService.UpdateBudget:input_type -> google.cloud.billing.budgets.v1beta1.UpdateBudgetRequest
	2, // 6: google.cloud.billing.budgets.v1beta1.BudgetService.GetBudget:input_type -> google.cloud.billing.budgets.v1beta1.GetBudgetRequest
	3, // 7: google.cloud.billing.budgets.v1beta1.BudgetService.ListBudgets:input_type -> google.cloud.billing.budgets.v1beta1.ListBudgetsRequest
	5, // 8: google.cloud.billing.budgets.v1beta1.BudgetService.DeleteBudget:input_type -> google.cloud.billing.budgets.v1beta1.DeleteBudgetRequest
	6, // 9: google.cloud.billing.budgets.v1beta1.BudgetService.CreateBudget:output_type -> google.cloud.billing.budgets.v1beta1.Budget
	6, // 10: google.cloud.billing.budgets.v1beta1.BudgetService.UpdateBudget:output_type -> google.cloud.billing.budgets.v1beta1.Budget
	6, // 11: google.cloud.billing.budgets.v1beta1.BudgetService.GetBudget:output_type -> google.cloud.billing.budgets.v1beta1.Budget
	4, // 12: google.cloud.billing.budgets.v1beta1.BudgetService.ListBudgets:output_type -> google.cloud.billing.budgets.v1beta1.ListBudgetsResponse
	8, // 13: google.cloud.billing.budgets.v1beta1.BudgetService.DeleteBudget:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_billing_budgets_v1beta1_budget_service_proto_init() }
func file_google_cloud_billing_budgets_v1beta1_budget_service_proto_init() {
	if File_google_cloud_billing_budgets_v1beta1_budget_service_proto != nil {
		return
	}
	file_google_cloud_billing_budgets_v1beta1_budget_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_billing_budgets_v1beta1_budget_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_billing_budgets_v1beta1_budget_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_billing_budgets_v1beta1_budget_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_billing_budgets_v1beta1_budget_service_proto_msgTypes,
	}.Build()
	File_google_cloud_billing_budgets_v1beta1_budget_service_proto = out.File
	file_google_cloud_billing_budgets_v1beta1_budget_service_proto_rawDesc = nil
	file_google_cloud_billing_budgets_v1beta1_budget_service_proto_goTypes = nil
	file_google_cloud_billing_budgets_v1beta1_budget_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BudgetServiceClient is the client API for BudgetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BudgetServiceClient interface {
	// Creates a new budget. See
	// [Quotas and limits](https://cloud.google.com/billing/quotas)
	// for more information on the limits of the number of budgets you can create.
	CreateBudget(ctx context.Context, in *CreateBudgetRequest, opts ...grpc.CallOption) (*Budget, error)
	// Updates a budget and returns the updated budget.
	//
	// WARNING: There are some fields exposed on the Google Cloud Console that
	// aren't available on this API. Budget fields that are not exposed in
	// this API will not be changed by this method.
	UpdateBudget(ctx context.Context, in *UpdateBudgetRequest, opts ...grpc.CallOption) (*Budget, error)
	// Returns a budget.
	//
	// WARNING: There are some fields exposed on the Google Cloud Console that
	// aren't available on this API. When reading from the API, you will not
	// see these fields in the return value, though they may have been set
	// in the Cloud Console.
	GetBudget(ctx context.Context, in *GetBudgetRequest, opts ...grpc.CallOption) (*Budget, error)
	// Returns a list of budgets for a billing account.
	//
	// WARNING: There are some fields exposed on the Google Cloud Console that
	// aren't available on this API. When reading from the API, you will not
	// see these fields in the return value, though they may have been set
	// in the Cloud Console.
	ListBudgets(ctx context.Context, in *ListBudgetsRequest, opts ...grpc.CallOption) (*ListBudgetsResponse, error)
	// Deletes a budget. Returns successfully if already deleted.
	DeleteBudget(ctx context.Context, in *DeleteBudgetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type budgetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBudgetServiceClient(cc grpc.ClientConnInterface) BudgetServiceClient {
	return &budgetServiceClient{cc}
}

func (c *budgetServiceClient) CreateBudget(ctx context.Context, in *CreateBudgetRequest, opts ...grpc.CallOption) (*Budget, error) {
	out := new(Budget)
	err := c.cc.Invoke(ctx, "/google.cloud.billing.budgets.v1beta1.BudgetService/CreateBudget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) UpdateBudget(ctx context.Context, in *UpdateBudgetRequest, opts ...grpc.CallOption) (*Budget, error) {
	out := new(Budget)
	err := c.cc.Invoke(ctx, "/google.cloud.billing.budgets.v1beta1.BudgetService/UpdateBudget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) GetBudget(ctx context.Context, in *GetBudgetRequest, opts ...grpc.CallOption) (*Budget, error) {
	out := new(Budget)
	err := c.cc.Invoke(ctx, "/google.cloud.billing.budgets.v1beta1.BudgetService/GetBudget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) ListBudgets(ctx context.Context, in *ListBudgetsRequest, opts ...grpc.CallOption) (*ListBudgetsResponse, error) {
	out := new(ListBudgetsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.billing.budgets.v1beta1.BudgetService/ListBudgets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) DeleteBudget(ctx context.Context, in *DeleteBudgetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.billing.budgets.v1beta1.BudgetService/DeleteBudget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BudgetServiceServer is the server API for BudgetService service.
type BudgetServiceServer interface {
	// Creates a new budget. See
	// [Quotas and limits](https://cloud.google.com/billing/quotas)
	// for more information on the limits of the number of budgets you can create.
	CreateBudget(context.Context, *CreateBudgetRequest) (*Budget, error)
	// Updates a budget and returns the updated budget.
	//
	// WARNING: There are some fields exposed on the Google Cloud Console that
	// aren't available on this API. Budget fields that are not exposed in
	// this API will not be changed by this method.
	UpdateBudget(context.Context, *UpdateBudgetRequest) (*Budget, error)
	// Returns a budget.
	//
	// WARNING: There are some fields exposed on the Google Cloud Console that
	// aren't available on this API. When reading from the API, you will not
	// see these fields in the return value, though they may have been set
	// in the Cloud Console.
	GetBudget(context.Context, *GetBudgetRequest) (*Budget, error)
	// Returns a list of budgets for a billing account.
	//
	// WARNING: There are some fields exposed on the Google Cloud Console that
	// aren't available on this API. When reading from the API, you will not
	// see these fields in the return value, though they may have been set
	// in the Cloud Console.
	ListBudgets(context.Context, *ListBudgetsRequest) (*ListBudgetsResponse, error)
	// Deletes a budget. Returns successfully if already deleted.
	DeleteBudget(context.Context, *DeleteBudgetRequest) (*emptypb.Empty, error)
}

// UnimplementedBudgetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBudgetServiceServer struct {
}

func (*UnimplementedBudgetServiceServer) CreateBudget(context.Context, *CreateBudgetRequest) (*Budget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBudget not implemented")
}
func (*UnimplementedBudgetServiceServer) UpdateBudget(context.Context, *UpdateBudgetRequest) (*Budget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBudget not implemented")
}
func (*UnimplementedBudgetServiceServer) GetBudget(context.Context, *GetBudgetRequest) (*Budget, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBudget not implemented")
}
func (*UnimplementedBudgetServiceServer) ListBudgets(context.Context, *ListBudgetsRequest) (*ListBudgetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBudgets not implemented")
}
func (*UnimplementedBudgetServiceServer) DeleteBudget(context.Context, *DeleteBudgetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBudget not implemented")
}

func RegisterBudgetServiceServer(s *grpc.Server, srv BudgetServiceServer) {
	s.RegisterService(&_BudgetService_serviceDesc, srv)
}

func _BudgetService_CreateBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).CreateBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.billing.budgets.v1beta1.BudgetService/CreateBudget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).CreateBudget(ctx, req.(*CreateBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_UpdateBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).UpdateBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.billing.budgets.v1beta1.BudgetService/UpdateBudget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).UpdateBudget(ctx, req.(*UpdateBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_GetBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).GetBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.billing.budgets.v1beta1.BudgetService/GetBudget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).GetBudget(ctx, req.(*GetBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_ListBudgets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBudgetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).ListBudgets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.billing.budgets.v1beta1.BudgetService/ListBudgets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).ListBudgets(ctx, req.(*ListBudgetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_DeleteBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).DeleteBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.billing.budgets.v1beta1.BudgetService/DeleteBudget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).DeleteBudget(ctx, req.(*DeleteBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BudgetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.billing.budgets.v1beta1.BudgetService",
	HandlerType: (*BudgetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBudget",
			Handler:    _BudgetService_CreateBudget_Handler,
		},
		{
			MethodName: "UpdateBudget",
			Handler:    _BudgetService_UpdateBudget_Handler,
		},
		{
			MethodName: "GetBudget",
			Handler:    _BudgetService_GetBudget_Handler,
		},
		{
			MethodName: "ListBudgets",
			Handler:    _BudgetService_ListBudgets_Handler,
		},
		{
			MethodName: "DeleteBudget",
			Handler:    _BudgetService_DeleteBudget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/billing/budgets/v1beta1/budget_service.proto",
}
