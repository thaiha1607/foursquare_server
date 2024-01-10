package http

import (
	"net/http"

	"github.com/thaiha1607/foursquare_server/ent"
	"github.com/thaiha1607/foursquare_server/models"
)

func handleError(err error) *models.FailureResponse {
	err_rsp := &models.FailureResponse{}
	if err_rsp = handleEntError(err); err_rsp != nil {
		return err_rsp
	}
	err_rsp.Message = err.Error()
	err_rsp.HttpStatusCode = http.StatusInternalServerError
	return err_rsp
}

func handleEntError(err error) *models.FailureResponse {
	err_rsp := &models.FailureResponse{}
	if ent.IsNotFound(err) {
		err_rsp.Message = err.Error()
		err_rsp.HttpStatusCode = http.StatusNotFound
		return err_rsp
	}
	if ent.IsValidationError(err) {
		err_rsp.Message = err.Error()
		err_rsp.HttpStatusCode = http.StatusBadRequest
		return err_rsp
	}
	if ent.IsConstraintError(err) {
		err_rsp.Message = err.Error()
		err_rsp.HttpStatusCode = http.StatusForbidden
		return err_rsp
	}
	return nil
}
