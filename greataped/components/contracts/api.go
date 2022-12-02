package contracts

import . "rail.town/infrastructure/components/api/protobuf"

type IApi interface {
	SetToken(string)
	SetDebugMode(bool)
	//API Methods
	SystemCall(*SystemCallRequest) (*SystemCallResult, error)
	Echo(*EchoRequest) (*EchoResult, error)
	Signup(*SignupRequest) (*SignupResult, error)
	Verify(*VerifyRequest) (*VerifyResult, error)
	Login(*LoginRequest) (*LoginResult, error)
	GetProfileByUser(*GetProfileByUserRequest) (*GetProfileByUserResult, error)
	UpdateProfileByUser(*UpdateProfileByUserRequest) (*UpdateProfileByUserResult, error)
	Logout(*LogoutRequest) (*LogoutResult, error)
	Webfinger(*WebfingerRequest) (*WebfingerResult, error)
	GetActor(*GetActorRequest) (*GetActorResult, error)
	FollowActor(*FollowActorRequest) (*FollowActorResult, error)
}
