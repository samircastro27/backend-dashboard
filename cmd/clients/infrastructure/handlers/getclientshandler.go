package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/samircastro27/backend-dashboard/cmd/clients/svc"
	"github.com/samircastro27/backend-dashboard/pkg/logger"
)

func GetUsersHandler(svc *svc.ServiceContext) http.HandlerFunc {
	useCaseHandler := svc.UseCase
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// params, err := types.GetParams(req.URL.Query())
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		clients, err := useCaseHandler.Execute()
		if err != nil {
			logger.LogErr("error is %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := &APIResponse{
			Success: true,
			Data:    clients,
			Error:   nil,
			Message: "Users found successfully",
		}

		if err := json.NewEncoder(w).Encode(data); err != nil {
			logger.LogErr()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
