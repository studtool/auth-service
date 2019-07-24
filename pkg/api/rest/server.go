package rest

import (
	"github.com/studtool/go-logs/pkg/logs"
	"github.com/studtool/go-rest/pkg/rest"
)

type Server struct {
	server *rest.Server
}

func NewServer() *Server {
	lPck := rest.LogPack{
		StructLogger: logs.NewStructLogger(logs.StructLoggerParams{
			CommonLoggerParams: logs.CommonLoggerParams{
				//TODO from runtime
			},
			StructWithPkgName: "rest.Server",
		}),
		StackTraceLogger: logs.NewStackTraceLogger(logs.StackTraceLoggerParams{
			CommonLoggerParams: logs.CommonLoggerParams{
				//TODO from runtime
			},
		}),
		RequestLogger: logs.NewRequestLogger(logs.RequestLoggerParams{
			CommonLoggerParams: logs.CommonLoggerParams{
				//TODO from runtime
			},
		}),
	}

	srv := &Server {}

	rCfg := rest.RouterConfig{
		PathPrefix: "/api/v0",
		Handlers: map[string][]rest.RequestHandler{
			"/public/auth/users": {
				rest.RequestHandler{
					Method: rest.MethodPost,
					Handler: srv.addUser,
				},
			},
		},
		Middleware: []rest.Middleware{
			rest.WithLogs(rest.LogsMiddlewareParams{
				RequestTypeDetector: func(ctx *rest.Context) string {
					return rest.RequestTypeTesting //TODO
				},
				RequestACLGroupDetector: func(ctx *rest.Context) string {
					return rest.ACLGroupCommon //XXX for now
				},
			}),
		},
	}

	srv.server = rest.NewServer(rest.ServerParams{
		Address: "", //TODO from config
		RouterConfig: rCfg,
		LogPack: lPck,
	})

	return srv
}
