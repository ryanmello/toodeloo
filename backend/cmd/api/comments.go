package main

import (
	"net/http"
	"strconv"
	"toodeloo/internal/store"

	"github.com/go-chi/chi/v5"
)

type CreateCommentPayload struct {
	Content string `json:"content"`
	UserId  int64  `json:"user_id"`
}

func (app *application) CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	postIdParam := chi.URLParam(r, "postId")
	
	postId, err := strconv.ParseInt(postIdParam, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	var payload CreateCommentPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	comment := &store.Comment{
		PostId:  postId,
		UserId:  payload.UserId,
		Content: payload.Content,
	}

	ctx := r.Context()

	if err := app.store.Comments.Create(ctx, comment); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, comment); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
} 
