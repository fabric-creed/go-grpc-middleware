// Copyright 2017 David Ackroyd. All Rights Reserved.
// See LICENSE for licensing terms.

package grpc_recovery_test

import (
	"github.com/fabric-creed/go-grpc-middleware"
	"github.com/fabric-creed/go-grpc-middleware/recovery"
	"github.com/fabric-creed/grpc"
	"github.com/fabric-creed/grpc/codes"
	"github.com/fabric-creed/grpc/status"
)

var (
	customFunc grpc_recovery.RecoveryHandlerFunc
)

// Initialization shows an initialization sequence with a custom recovery handler func.
func Example_initialization() {
	// Define customfunc to handle panic
	customFunc = func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(customFunc),
	}
	// Create a server. Recovery handlers should typically be last in the chain so that other middleware
	// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
	_ = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(opts...),
		),
	)
}
