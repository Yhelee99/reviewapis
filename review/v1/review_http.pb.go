// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.0--rc3
// source: api/review/v1/review.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationReviewAuditAppeal = "/api.review.v1.Review/AuditAppeal"
const OperationReviewAuditReview = "/api.review.v1.Review/AuditReview"
const OperationReviewCreateReview = "/api.review.v1.Review/CreateReview"
const OperationReviewGetReview = "/api.review.v1.Review/GetReview"
const OperationReviewListReviewByUserID = "/api.review.v1.Review/ListReviewByUserID"

type ReviewHTTPServer interface {
	// AuditAppeal 审核申诉
	AuditAppeal(context.Context, *AuditAppealRequest) (*AuditAppealReply, error)
	// AuditReview O端
	// 审核评价
	AuditReview(context.Context, *AuditReviewRequest) (*AuditReviewReply, error)
	// CreateReview C端
	// CreateReview 创建评价
	CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewReply, error)
	// GetReview GetReview 获取评价详情
	GetReview(context.Context, *GetReviewRequest) (*GetReviewReply, error)
	// ListReviewByUserID ListReviewByUserID 查看userID下所有评价
	ListReviewByUserID(context.Context, *ListReviewByUserIDRequest) (*ListReviewByUserIDReply, error)
}

func RegisterReviewHTTPServer(s *http.Server, srv ReviewHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/review/create", _Review_CreateReview0_HTTP_Handler(srv))
	r.GET("/v1/review/getReview/{reviewID}", _Review_GetReview0_HTTP_Handler(srv))
	r.GET("/v1/review/getReviewList", _Review_ListReviewByUserID0_HTTP_Handler(srv))
	r.POST("/v1/audit/review", _Review_AuditReview0_HTTP_Handler(srv))
	r.POST("/v1/audit/appeal", _Review_AuditAppeal0_HTTP_Handler(srv))
}

func _Review_CreateReview0_HTTP_Handler(srv ReviewHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateReviewRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReviewCreateReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateReview(ctx, req.(*CreateReviewRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateReviewReply)
		return ctx.Result(200, reply)
	}
}

func _Review_GetReview0_HTTP_Handler(srv ReviewHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetReviewRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReviewGetReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetReview(ctx, req.(*GetReviewRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetReviewReply)
		return ctx.Result(200, reply)
	}
}

func _Review_ListReviewByUserID0_HTTP_Handler(srv ReviewHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListReviewByUserIDRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReviewListReviewByUserID)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListReviewByUserID(ctx, req.(*ListReviewByUserIDRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListReviewByUserIDReply)
		return ctx.Result(200, reply)
	}
}

func _Review_AuditReview0_HTTP_Handler(srv ReviewHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuditReviewRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReviewAuditReview)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuditReview(ctx, req.(*AuditReviewRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuditReviewReply)
		return ctx.Result(200, reply)
	}
}

func _Review_AuditAppeal0_HTTP_Handler(srv ReviewHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuditAppealRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReviewAuditAppeal)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuditAppeal(ctx, req.(*AuditAppealRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuditAppealReply)
		return ctx.Result(200, reply)
	}
}

type ReviewHTTPClient interface {
	AuditAppeal(ctx context.Context, req *AuditAppealRequest, opts ...http.CallOption) (rsp *AuditAppealReply, err error)
	AuditReview(ctx context.Context, req *AuditReviewRequest, opts ...http.CallOption) (rsp *AuditReviewReply, err error)
	CreateReview(ctx context.Context, req *CreateReviewRequest, opts ...http.CallOption) (rsp *CreateReviewReply, err error)
	GetReview(ctx context.Context, req *GetReviewRequest, opts ...http.CallOption) (rsp *GetReviewReply, err error)
	ListReviewByUserID(ctx context.Context, req *ListReviewByUserIDRequest, opts ...http.CallOption) (rsp *ListReviewByUserIDReply, err error)
}

type ReviewHTTPClientImpl struct {
	cc *http.Client
}

func NewReviewHTTPClient(client *http.Client) ReviewHTTPClient {
	return &ReviewHTTPClientImpl{client}
}

func (c *ReviewHTTPClientImpl) AuditAppeal(ctx context.Context, in *AuditAppealRequest, opts ...http.CallOption) (*AuditAppealReply, error) {
	var out AuditAppealReply
	pattern := "/v1/audit/appeal"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationReviewAuditAppeal))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ReviewHTTPClientImpl) AuditReview(ctx context.Context, in *AuditReviewRequest, opts ...http.CallOption) (*AuditReviewReply, error) {
	var out AuditReviewReply
	pattern := "/v1/audit/review"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationReviewAuditReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ReviewHTTPClientImpl) CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...http.CallOption) (*CreateReviewReply, error) {
	var out CreateReviewReply
	pattern := "/v1/review/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationReviewCreateReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ReviewHTTPClientImpl) GetReview(ctx context.Context, in *GetReviewRequest, opts ...http.CallOption) (*GetReviewReply, error) {
	var out GetReviewReply
	pattern := "/v1/review/getReview/{reviewID}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationReviewGetReview))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *ReviewHTTPClientImpl) ListReviewByUserID(ctx context.Context, in *ListReviewByUserIDRequest, opts ...http.CallOption) (*ListReviewByUserIDReply, error) {
	var out ListReviewByUserIDReply
	pattern := "/v1/review/getReviewList"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationReviewListReviewByUserID))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
