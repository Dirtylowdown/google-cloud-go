Delete









































































delete













// The client will use your default application credentials. Clients should be reused instead of created as needed.
// The methods of Client are safe for concurrent use by multiple goroutines.
// The returned client must be Closed when it is done being used.
//
// # Using the Client
//
// The following is an example of making an API call with the newly created client, mentioned above.
//
//	req := &billingpb.CreateBillingAccountRequest{
//		// TODO: Fill request struct fields.
//		// See https://pkg.go.dev/cloud.google.com/go/billing/apiv1/billingpb#CreateBillingAccountRequest.
//	}
//	resp, err := c.CreateBillingAccount(ctx, req)
//	if err != nil {
//		// TODO: Handle error.
//	}
//	// TODO: Use resp.
//	_ = resp
//
// # Use of Context
//
// The ctx passed to NewCloudBillingClient is used for authentication requests and
// for creating the underlying connection, but is not used for subsequent calls.
// Individual methods on the client use the ctx given to them.
//
// To close the open connection, use the Close() method.
//
// [Authentication and Authorization]: https://pkg.go.dev/cloud.google.com/go#hdr-Authentication_and_Authorization
// [Timeouts and Cancellation]: https://pkg.go.dev/cloud.google.com/go#hdr-Timeouts_and_Cancellation
// [Testing against Client Libraries]: https://pkg.go.dev/cloud.google.com/go#hdr-Testing
// [Debugging Client Libraries]: https://pkg.go.dev/cloud.google.com/go#hdr-Debugging
// [Inspecting errors]: https://pkg.go.dev/cloud.google.com/go#hdr-Inspecting_errors
package billing // import "cloud.google.com/go/billing/apiv1"
