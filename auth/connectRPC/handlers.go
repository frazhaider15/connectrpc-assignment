package connectrpc

import (
	"context"

	"connectrpc.com/connect"

	protov1 "github.com/auth/gen/authproto/v1" // generated by protoc-gen-go
	"github.com/auth/services"
)

func (s *AuthService) SignupWithPhoneNumber(ctx context.Context, req *connect.Request[protov1.SignupRequest]) (*connect.Response[protov1.Response], error) {

	err := services.UserSignup(req.Msg.Name, req.Msg.Phone)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&protov1.Response{
		Message: "User Signed up successfully",
	})

	return res, nil
}

func (s *AuthService) VerifyPhoneNumber(ctx context.Context, req *connect.Request[protov1.VerifyPhoneRequest]) (*connect.Response[protov1.Response], error) {
	err := services.VerfifySignupOtp(req.Msg.Otp, req.Msg.Phone)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&protov1.Response{
		Message: "Otp Verified successfully",
	})

	return res, nil
}

func (s *AuthService) LoginWithPhoneNumber(ctx context.Context, req *connect.Request[protov1.LoginWithPhoneNumberRequest]) (*connect.Response[protov1.Response], error) {
	err := services.UserLogin(req.Msg.Phone)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&protov1.Response{
		Message: "User Login Otp sent successfully",
	})

	return res, nil
}

func (s *AuthService) ValidatePhoneNumberLogin(ctx context.Context, req *connect.Request[protov1.ValidatePhoneNumberLoginRequest]) (*connect.Response[protov1.Response], error) {
	err := services.VerifyLoginOtp(req.Msg.Otp, req.Msg.Phone)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&protov1.Response{
		Message: "Otp Verified successfully",
	})

	return res, nil
}

func (s *AuthService) GetProfile(ctx context.Context, req *connect.Request[protov1.GetProfileRequest]) (*connect.Response[protov1.GetProfileResponse], error) {

	user, err := services.GetProfile(req.Msg.Phone)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&protov1.GetProfileResponse{
		Name:          user.Name,
		Phone:         user.Phone,
		PhoneVerified: user.PhoneVerified,
	})

	return res, nil
}