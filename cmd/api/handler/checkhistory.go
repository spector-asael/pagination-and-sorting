// Filename: cmd/api/checkhistory.go
package handler

import (
	"net/http"

	"github.com/spector-asael/banking/internal/data"
	"github.com/spector-asael/banking/internal/validator"
)

func (a *ApplicationDependencies) checkHistoryHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	var input struct {
		UserID int64 `json:"user_id"`
	}

	// Decode JSON
	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	// Validate input using the model's validator
	v := validator.New()
	data.ValidateHistory(v, input.UserID)
	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Get ledger history
	history, err := a.Models.History.GetByUserID(input.UserID)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	// Return JSON response
	err = a.writeJSON(
		w,
		http.StatusOK,
		envelope{"history": history},
		nil,
	)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}