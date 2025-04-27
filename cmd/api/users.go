package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"toodeloo/internal/store"

	"github.com/go-chi/chi/v5"
)

type userKey string

const userCtx userKey = "user"

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type UpdateUserPayload struct {
	Username *string `json:"username" validate:"omitempty,max=100"`
	Email    *string `json:"email" validate:"omitempty,max=100"`
	Password *string `json:"password" validate:"omitempty,max=100"`
}

func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromCtx(r)

	ctx := r.Context()

	var payload UpdateUserPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Username != nil {
		user.Username = *payload.Username
	}

	if payload.Email != nil {
		user.Email = *payload.Email
	}

	if payload.Password != nil {
		user.Password = *payload.Password
	}

	if err := app.store.Users.Update(ctx, user); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) usersContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "userId")
		id, err := strconv.ParseInt(idParam, 10, 64)

		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		user, err := app.store.Users.GetById(ctx, id)

		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserFromCtx(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}

type FollowUserPayload struct {
	FollowerId int64 `json:"follower_id"`
}

// handle follow
func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "userId")
	userId, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	var payload FollowUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Follow(ctx, payload.FollowerId, userId); err != nil {
		switch err {
		case store.ErrConflict:
			app.conflictResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}

	}

	if err := app.jsonResponse(w, http.StatusNoContent, userId); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// handle unfollow
func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "userId")
	userId, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	var payload FollowUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Unfollow(ctx, payload.FollowerId, userId); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, userId); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
