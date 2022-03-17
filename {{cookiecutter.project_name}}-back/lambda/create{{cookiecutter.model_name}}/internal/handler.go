package internal

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/divilo/{{cookiecutter.project_name}}-back/internal/service"
	"github.com/divilo/utils-go/interfaces"
	"github.com/divilo/utils-go/service/eventmapper"
	"go.uber.org/zap"
	"net/http"
)

type handler struct {
	lgr            *zap.SugaredLogger
	corsMiddleware interfaces.APIGatewayProxyMiddleware
	eventMapper    eventmapper.ServiceEventMapper
	{{cookiecutter.model_name}}ervice  service.{{cookiecutter.model_name}}ervice
}

// New returns a Handler instance
func New(lgr *zap.SugaredLogger, corsMiddleware interfaces.APIGatewayProxyMiddleware, eventMapper eventmapper.ServiceEventMapper, {{cookiecutter.model_name}}ervice service.{{cookiecutter.model_name}}ervice) interfaces.APIGatewayProxyHandler {
	return &handler{lgr, corsMiddleware, eventMapper, {{cookiecutter.model_name}}ervice}
}

// HandleProxy implements API Gateway proxy events handling
func (h handler) HandleProxy() interfaces.APIGatewayProxyHandlerFunc {
	return h.corsMiddleware.ProxyMiddleware(
		func(ctx context.Context, event *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
			defer h.lgr.Sync()
			lgr := h.lgr.With("APIGatewayProxyRequest", event)
			lgr.Debug("create device request")
			var req = &handlerRequest{}
			// Map request and validate
			err := h.eventMapper.FromProxyRequest(event, req)
			if err != nil {
				lgr.Warnw("the request is not valid", zap.Error(err))
				return h.eventMapper.ToProxyResponse(http.StatusBadRequest, "")
			}

			// Do
			dvc := toModelDevice(&req.CreateDeviceRequest)
			_, err = h.{{cookiecutter.model_name}}ervice.Create(ctx, *dvc)
			if err != nil {
				lgr.Errorw("Unexpected error", zap.Error(err))
				return h.eventMapper.ToProxyResponse(http.StatusInternalServerError, "")
			}
			return h.eventMapper.ToProxyResponse(http.StatusNoContent, "")
		},
	)
}
