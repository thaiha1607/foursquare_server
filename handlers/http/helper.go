package http

import (
	"net/http"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/models"
)

func handleError(err error) *models.FailureResponse {
	if err_rsp := handleEntError(err); err_rsp != nil {
		return err_rsp
	}
	return &models.FailureResponse{
		Message:        err.Error(),
		HttpStatusCode: http.StatusInternalServerError,
	}
}

func handleEntError(err error) *models.FailureResponse {
	err_rsp := &models.FailureResponse{}
	err_rsp.Message = err.Error()
	if ent.IsNotFound(err) {
		err_rsp.HttpStatusCode = http.StatusNotFound
	}
	if ent.IsValidationError(err) {
		err_rsp.HttpStatusCode = http.StatusBadRequest
	}
	if ent.IsConstraintError(err) {
		err_rsp.HttpStatusCode = http.StatusForbidden
	}
	if err_rsp.HttpStatusCode != 0 {
		return err_rsp
	}
	return nil
}
