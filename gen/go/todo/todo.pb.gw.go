// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: todo/todo.proto

/*
Package todo is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package todo

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_TodoApi_Get_0(ctx context.Context, marshaler runtime.Marshaler, client TodoApiClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetRequest
	var metadata runtime.ServerMetadata

	msg, err := client.Get(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_TodoApi_Get_0(ctx context.Context, marshaler runtime.Marshaler, server TodoApiServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetRequest
	var metadata runtime.ServerMetadata

	msg, err := server.Get(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterTodoApiHandlerServer registers the http handlers for service TodoApi to "mux".
// UnaryRPC     :call TodoApiServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterTodoApiHandlerFromEndpoint instead.
func RegisterTodoApiHandlerServer(ctx context.Context, mux *runtime.ServeMux, server TodoApiServer) error {

	mux.Handle("GET", pattern_TodoApi_Get_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/todo_service.TodoApi/Get", runtime.WithHTTPPathPattern("/todo"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_TodoApi_Get_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TodoApi_Get_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterTodoApiHandlerFromEndpoint is same as RegisterTodoApiHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterTodoApiHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterTodoApiHandler(ctx, mux, conn)
}

// RegisterTodoApiHandler registers the http handlers for service TodoApi to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterTodoApiHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterTodoApiHandlerClient(ctx, mux, NewTodoApiClient(conn))
}

// RegisterTodoApiHandlerClient registers the http handlers for service TodoApi
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "TodoApiClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "TodoApiClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "TodoApiClient" to call the correct interceptors.
func RegisterTodoApiHandlerClient(ctx context.Context, mux *runtime.ServeMux, client TodoApiClient) error {

	mux.Handle("GET", pattern_TodoApi_Get_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/todo_service.TodoApi/Get", runtime.WithHTTPPathPattern("/todo"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TodoApi_Get_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TodoApi_Get_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_TodoApi_Get_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"todo"}, ""))
)

var (
	forward_TodoApi_Get_0 = runtime.ForwardResponseMessage
)
