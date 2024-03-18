package login_service_v1

import (
	"context"
	common "github.com/acezsq/project-common"
	"github.com/acezsq/project-common/errs"
	"github.com/acezsq/project-user/pkg/dao"
	"github.com/acezsq/project-user/pkg/model"
	"github.com/acezsq/project-user/pkg/repo"
	"go.uber.org/zap"
	"log"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	// 1. 获取参数
	mobile := msg.Mobile
	// 2. 校验参数
	if !common.VerifyMobile(mobile) {
		return nil, errs.GrpcError(model.NoLegalMobile)
	}
	// 3. 生成验证码（4位1000-9999或者六位100000-999999）
	code := "123456"
	// 4. 调用短信平台（三方放入go协程中执行 接口可以快速响应）
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("短信平台调用成功，发送短信")
		// redis 假设后续缓存可能存在mysql,mongo当中
		// 5. 存储验证码 redis 当中  过期时间15分钟
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := ls.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Printf("验证码存入redis出错，cause by:%v \n", err)
		}
	}()
	return &CaptchaResponse{Code: code}, nil
}
