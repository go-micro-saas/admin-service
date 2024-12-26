package dto

import (
	enumv1 "github.com/go-micro-saas/account-service/api/account-service/v1/enums"
	resourcev1 "github.com/go-micro-saas/account-service/api/account-service/v1/resources"
	"github.com/go-micro-saas/account-service/app/account-service/internal/biz/bo"
	"github.com/go-micro-saas/account-service/app/account-service/internal/data/po"
	timepkg "github.com/ikaiguang/go-srv-kit/kit/time"
)

var (
	AccountDto accountDto
)

type accountDto struct{}

func (s *accountDto) ToPbLoginRespData(userModel *po.User, tokenResp *bo.SignTokenResp) *resourcev1.LoginRespData {
	res := &resourcev1.LoginRespData{
		UserInfo: &resourcev1.UserInfo{
			// Id:           userModel.Id,
			UserId:       userModel.UserId,
			UserNickname: userModel.UserNickname,
			UserAvatar:   userModel.UserAvatar,
			// UserGender:   userModel.UserGender,
			// UserStatus:   userModel.UserStatus,
		},
		AccessToken:           tokenResp.AccessToken,
		AccessTokenExpiredAt:  tokenResp.AccessTokenItem.ExpiredAt,
		RefreshToken:          tokenResp.RefreshToken,
		RefreshTokenExpiredAt: tokenResp.RefreshTokenItem.ExpiredAt,
	}
	return res
}

func (s *accountDto) ToPbUser(dataModel *po.User) *resourcev1.User {
	newDataModel := &resourcev1.User{
		Id:          dataModel.Id,                                 // ID
		UserId:      dataModel.UserId,                             // uid
		CreatedTime: dataModel.CreatedTime.Format(timepkg.YmdHms), // 创建时间
		UpdatedTime: dataModel.UpdatedTime.Format(timepkg.YmdHms), // 最后修改时间
		// DeletedTime:  dataModel.DeletedTime,                        // 删除时间
		UserEmail:    dataModel.UserEmail,    // 邮箱
		UserNickname: dataModel.UserNickname, // 昵称
		UserAvatar:   dataModel.UserAvatar,   // 头像
		UserGender:   dataModel.UserGender,   // 性别
		UserStatus:   dataModel.UserStatus,   // 状态
		//RegisterType: dataModel.RegisterType, // 注册类型
		//PasswordHash: dataModel.PasswordHash, // 密码
	}
	// ActiveBeginTime 激活开始时间
	//if dataModel.ActiveBeginTime != nil {
	//	newDataModel.ActiveBeginTime = dataModel.ActiveBeginTime.Format(timepkg.YmdHms)
	//}
	// ActiveEndTime 激活结束时间
	//if dataModel.ActiveEndTime != nil {
	//	newDataModel.ActiveEndTime = dataModel.ActiveEndTime.Format(timepkg.YmdHms)
	//}
	// DisableTime 禁用时间
	//if dataModel.DisableTime != nil {
	//	newDataModel.DisableTime = dataModel.DisableTime.Format(timepkg.YmdHms)
	//}
	// BlacklistTime 黑名单时间
	//if dataModel.BlacklistTime != nil {
	//	newDataModel.BlacklistTime = dataModel.BlacklistTime.Format(timepkg.YmdHms)
	//}

	return newDataModel
}

func (s *accountDto) ToBoSendVerifyCodeParam(req *resourcev1.SendPhoneSignupCodeReq) *bo.SendVerifyCodeParam {
	res := &bo.SendVerifyCodeParam{
		VerifyAccount: req.Phone,
		VerifyType:    enumv1.UserVerifyTypeEnum_SIGNUP_BY_PHONE,
	}
	return res
}

func (s *accountDto) ToBoSendVerifyCodeParam2(req *resourcev1.SendEmailSignupCodeReq) *bo.SendVerifyCodeParam {
	res := &bo.SendVerifyCodeParam{
		VerifyAccount: req.Email,
		VerifyType:    enumv1.UserVerifyTypeEnum_SIGNUP_BY_EMAIL,
	}
	return res
}

func (s *accountDto) ToPbSendSignupCodeRespData(dataModel *bo.SendVerifyCodeReply) *resourcev1.SendSignupCodeRespData {
	res := &resourcev1.SendSignupCodeRespData{
		Code: "",
	}
	if !dataModel.IsSendToMQ {
		res.Code = dataModel.VerifyCode
	}
	return res
}

func (s *accountDto) ToBoSendEmailCodeParam(dataModel *bo.SendVerifyCodeReply) *bo.SendEmailCodeParam {
	res := &bo.SendEmailCodeParam{
		VerifyAccount: dataModel.VerifyAccount,
		VerifyCode:    dataModel.VerifyCode,
	}
	return res
}

func (s *accountDto) ToBoSignupByPhoneParam(req *resourcev1.SignupByPhoneReq) *bo.SignupByPhoneParam {
	res := &bo.SignupByPhoneParam{
		Phone:           req.Phone,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		Code:            req.Code,
		SkipVerifyCode:  false,
	}
	return res
}

func (s *accountDto) ToBoSignupByEmailParam(req *resourcev1.SignupByEmailReq) *bo.SignupByEmailParam {
	res := &bo.SignupByEmailParam{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
		Code:            req.Code,
		SkipVerifyCode:  false,
	}
	return res
}

func (s *accountDto) ToBoLoginByPhoneParam(req *resourcev1.LoginByPhoneReq) *bo.LoginByPhoneParam {
	res := &bo.LoginByPhoneParam{
		Phone:          req.Phone,
		Password:       req.Password,
		Code:           req.Code,
		SkipVerifyCode: false,
	}
	return res
}

func (s *accountDto) ToBoLoginByEmailParam(req *resourcev1.LoginByEmailReq) *bo.LoginByEmailParam {
	res := &bo.LoginByEmailParam{
		Email:          req.Email,
		Password:       req.Password,
		Code:           req.Code,
		SkipVerifyCode: false,
	}
	return res
}
