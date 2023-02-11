package global

import (
	"github.com/xlzhangkeke/blog-service/pkg/logger"
	"github.com/xlzhangkeke/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	EmailSetting    *setting.EmailSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
