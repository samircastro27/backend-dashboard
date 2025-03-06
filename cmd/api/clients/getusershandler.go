package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	response "github.com/samircastro27/backend-dashboard/cmd/api/http"
	"github.com/samircastro27/backend-dashboard/cmd/api/middlewares"
	"github.com/samircastro27/backend-dashboard/cmd/api/svc"
	"github.com/samircastro27/backend-dashboard/pkg/logger"
)

func GetUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		daprCli := svcCtx.DaprCli

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		o, err := daprCli.InvokeMethodWithContent(ctx, "clients", "/v1/clients", "get", &dapr.DataContent{
			ContentType: "application/json",
			Data:        []byte(""),
		})

		if err != nil {
			logger.LogErrWithDetails("", "", fmt.Sprintf("Error invoking method: %v", err))
			middlewares.CaptureError(fmt.Errorf("error invoking method: %v", err))
			response.JSONResponse(w, http.StatusInternalServerError, &response.APIResponse{
				Success: false,
				Data:    nil,
				Error:   fmt.Sprintf("Error invoking method: %v", err),
			})
			return
		}

		res := &response.APIResponse{}

		if err = json.Unmarshal(o, res); err != nil {
			logger.LogErrWithDetails("", "", fmt.Sprintf("Error unmarshalling response: %v", err))
			middlewares.CaptureError(fmt.Errorf("error unmarshalling response: %v", err))
			response.JSONResponse(w, http.StatusInternalServerError, &response.APIResponse{
				Success: false,
				Data:    nil,
				Error:   "Error unmarshalling response",
			})
			return
		}

		response.JSONResponse(w, http.StatusOK, res)
	}
}
