package boot

import "tiger-go/tiger"

// 重要：必须在此方法中初始化配置文件
func init() {
	tiger.Config().LoadConfig("./config/main", "./config/main-local", "applications/api/config/api")
}
