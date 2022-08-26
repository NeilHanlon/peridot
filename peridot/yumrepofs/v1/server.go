// Copyright (c) All respective contributors to the Peridot Project. All rights reserved.
// Copyright (c) 2021-2022 Rocky Enterprise Software Foundation, Inc. All rights reserved.
// Copyright (c) 2021-2022 Ctrl IQ, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
// may be used to endorse or promote products derived from this software without
// specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package yumrepofsv1

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	commonpb "peridot.resf.org/common"
	peridotdb "peridot.resf.org/peridot/db"
	"peridot.resf.org/peridot/lookaside"
	"peridot.resf.org/peridot/lookaside/s3"
	yumrepofspb "peridot.resf.org/peridot/yumrepofs/pb"
	"peridot.resf.org/utils"
)

type Server struct {
	yumrepofspb.UnimplementedYumrepofsServiceServer

	log     *logrus.Logger
	db      peridotdb.Access
	storage lookaside.Storage
	s3      *awss3.S3
}

func NewServer(db peridotdb.Access, session *session.Session) (*Server, error) {
	storage, err := s3.New(osfs.New("/"))
	if err != nil {
		return nil, err
	}

	cfg := &aws.Config{}
	if assumeRole := viper.GetString("s3-assume-role"); assumeRole != "" {
		cfg.Credentials = stscreds.NewCredentials(session, assumeRole)
	}

	return &Server{
		log:     logrus.New(),
		db:      db,
		storage: storage,
		s3:      awss3.New(session, cfg),
	}, nil
}

func (s *Server) interceptor(ctx context.Context, req interface{}, usi *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	n := utils.EndInterceptor

	return n(ctx, req, usi, handler)
}

func (s *Server) Run() {
	res := utils.NewGRPCServer(
		&utils.GRPCOptions{
			Interceptor: s.interceptor,
		},
		func(r *utils.Register) {
			endpoints := []utils.GrpcEndpointRegister{
				commonpb.RegisterHealthCheckServiceHandlerFromEndpoint,
				yumrepofspb.RegisterYumrepofsServiceHandlerFromEndpoint,
			}

			for _, endpoint := range endpoints {
				err := endpoint(r.Context, r.Mux, r.Endpoint, r.Options)
				if err != nil {
					s.log.Fatalf("could not register handler - %v", err)
				}
			}
		},
		func(r *utils.RegisterServer) {
			commonpb.RegisterHealthCheckServiceServer(r.Server, &utils.HealthServer{})

			yumrepofspb.RegisterYumrepofsServiceServer(r.Server, s)
		},
	)

	defer res.Cancel()
	res.WaitGroup.Wait()
}

func (s *Server) GetPublicUrl(_ context.Context, _ *yumrepofspb.GetPublicUrlRequest) (*yumrepofspb.GetPublicUrlResponse, error) {
	publicUrl := os.Getenv("YUMREPOFS_HTTP_PUBLIC_URL")
	if publicUrl == "" {
		return nil, status.Error(codes.FailedPrecondition, "Private instance or missing YUMREPOFS_HTTP_PUBLIC_URL")
	}

	return &yumrepofspb.GetPublicUrlResponse{
		PublicUrl: publicUrl,
	}, nil
}
