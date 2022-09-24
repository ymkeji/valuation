package server

import (
	"context"
	"net/http"

	"valuation/api/valuation/v1"
	"valuation/internal/conf"
	"valuation/internal/service"
	"valuation/pkg/aop"
	"valuation/pkg/errorx"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/pprof"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, good *service.GoodService, user *service.UserService, logger log.Logger) *kratosHttp.Server {
	var opts = []kratosHttp.ServerOption{
		kratosHttp.ResponseEncoder(encodeResponseFunc),
		kratosHttp.ErrorEncoder(errorEncoder),
		kratosHttp.Middleware(
			aop.Recovery(),
			//selector.Server(
			//	aop.JWTAuthMiddleware(),
			//).Match(NewWhiteListMatcher()).Build(),
			aop.Validator(),
		),

		kratosHttp.Filter(
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"}),
			)),
	}
	if c.Http.Network != "" {
		opts = append(opts, kratosHttp.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, kratosHttp.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, kratosHttp.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := kratosHttp.NewServer(opts...)
	srv.Handle("", pprof.NewHandler())
	route := srv.Route("/",
		aop.FilterRecovery,
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"}),
		))
	goods := service.GoodService{}
	route.POST("/good/upload", goods.GoodUpload)
	valuation.RegisterGoodHTTPServer(srv, good)
	valuation.RegisterUserHTTPServer(srv, user)
	return srv
}

func encodeResponseFunc(w http.ResponseWriter, r *http.Request, v interface{}) error {
	type response struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}
	reply := &response{
		Code: 200,
		Data: v,
		Msg:  "ok",
	}
	codec := encoding.GetCodec("json")
	data, err := codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func errorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errorx.FromError(err)
	codec, _ := kratosHttp.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(se.Code)
	_, _ = w.Write(body)
}

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	//whiteList["/api.valuation.v1.User/UserLogin"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
