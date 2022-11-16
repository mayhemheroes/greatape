package core

import (
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

func (api *api) SystemCall(request *SystemCallRequest) (*SystemCallResult, error) {
	result, err := api.call(SYSTEM_CALL_REQUEST, request)

	if err != nil {
		return nil, err
	} else {
		return result.(*SystemCallResult), nil
	}
}

func (api *api) Echo(request *EchoRequest) (*EchoResult, error) {
	result, err := api.call(ECHO_REQUEST, request)

	if err != nil {
		return nil, err
	} else {
		return result.(*EchoResult), nil
	}
}

func (api *api) Signup(request *SignupRequest) (*SignupResult, error) {
	result, err := api.call(SIGNUP_REQUEST, request)

	if err != nil {
		return nil, err
	} else {
		return result.(*SignupResult), nil
	}
}

func (api *api) Verify(request *VerifyRequest) (*VerifyResult, error) {
	result, err := api.call(VERIFY_REQUEST, request)

	if err != nil {
		return nil, err
	} else {
		return result.(*VerifyResult), nil
	}
}

func (api *api) Login(request *LoginRequest) (*LoginResult, error) {
	result, err := api.call(LOGIN_REQUEST, request)

	if err != nil {
		return nil, err
	} else {
		return result.(*LoginResult), nil
	}
}

func (api *api) GetProfileByUser(request *GetProfileByUserRequest) (*GetProfileByUserResult, error) {
	result, err := api.call(GET_PROFILE_BY_USER_REQUEST, request)

	if err != nil {
		return nil, err
	} else {
		return result.(*GetProfileByUserResult), nil
	}
}

func (api *api) UpdateProfileByUser(request *UpdateProfileByUserRequest) (*UpdateProfileByUserResult, error) {
	result, err := api.call(UPDATE_PROFILE_BY_USER_REQUEST, request)

	if err != nil {
		return nil, err
	} else {
		return result.(*UpdateProfileByUserResult), nil
	}
}

func init() {
	API_RESULT[SYSTEM_CALL_RESULT] = SystemCallResult{}
	API_RESULT[ECHO_RESULT] = EchoResult{}
	API_RESULT[SIGNUP_RESULT] = SignupResult{}
	API_RESULT[VERIFY_RESULT] = VerifyResult{}
	API_RESULT[LOGIN_RESULT] = LoginResult{}
	API_RESULT[GET_PROFILE_BY_USER_RESULT] = GetProfileByUserResult{}
	API_RESULT[UPDATE_PROFILE_BY_USER_RESULT] = UpdateProfileByUserResult{}
}
