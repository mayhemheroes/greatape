package handlers

import (
	"net/http"

	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

type getProfileByUserHandler struct {
}

func GetProfileByUserHandler() IHttpHandler {
	return &getProfileByUserHandler{}
}

func (handler *getProfileByUserHandler) Method() string {
	return http.MethodGet
}

func (handler *getProfileByUserHandler) Path() string {
	return "/api/v1/profile"
}

func (handler *getProfileByUserHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &GetProfileByUserRequest{}
		result := &GetProfileByUserResult{}

		onRequestUnmarshalled := func(request *GetProfileByUserRequest) {
		}

		return pipeline.Handle(x,
			"get_profile_by_user",
			GET_PROFILE_BY_USER_REQUEST,
			GET_PROFILE_BY_USER_RESULT,
			request, result,
			onRequestUnmarshalled,
			false,
		)
	}
}
