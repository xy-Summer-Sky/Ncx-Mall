package system

type ServiceGroup struct {
	JwtService
	UserService
	CasbinService
	InitDBService
	SystemConfigService
}
