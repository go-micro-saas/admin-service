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

func (s *accountDto) ToBoSendVerifyCodeParam(req *resourcev1.SendPhoneVerifyCodeReq) *bo.SendVerifyCodeParam {
	res := &bo.SendVerifyCodeParam{
		VerifyAccount: req.Phone,
		VerifyType:    enumv1.UserConfirmTypeEnum_PHONE,
	}
	return res
}

func (s *accountDto) ToBoSendVerifyCodeParam2(req *resourcev1.SendEmailVerifyCodeReq) *bo.SendVerifyCodeParam {
	res := &bo.SendVerifyCodeParam{
		VerifyAccount: req.Email,
		VerifyType:    enumv1.UserConfirmTypeEnum_EMAIL,
	}
	return res
}

func (s *accountDto) ToPbSendVerifyCodeRespData(dataModel *bo.SendVerifyCodeReply) *resourcev1.SendVerifyCodeRespData {
	res := &resourcev1.SendVerifyCodeRespData{
		Code: "",
	}
	if !dataModel.IsSendSuccess {
		res.Code = dataModel.Code
	}
	return res
}
