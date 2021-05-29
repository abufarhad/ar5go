package container

import (
	"clean/app/http/controllers"
	"clean/app/http/middlewares"
	repoImpl "clean/app/repository/impl"
	svcImpl "clean/app/svc/impl"
	"clean/infra/conn"
)

func Init(g interface{}) {
	db := conn.Db()
	redis := conn.Redis()
	acl := middlewares.ACL

	// register all repos impl, services impl, controllers
	sysRepo := repoImpl.NewSystemRepository(db, redis)
	companyRepo := repoImpl.NewMySqlCompanyRepository(db)
	userRepo := repoImpl.NewMySqlUsersRepository(db)
	locationRepo := repoImpl.NewMySqlLocationRepository(db)
	roleRepo := repoImpl.NewMySqlRolesRepository(db)
	permissionRepo := repoImpl.NewMySqlPermissionsRepository(db)

	cacheSvc := svcImpl.NewRedisService(redis)
	sysSvc := svcImpl.NewSystemService(sysRepo)
	companySvc := svcImpl.NewCompanyService(companyRepo, userRepo)
	userSvc := svcImpl.NewUsersService(userRepo, cacheSvc)
	tokenSvc := svcImpl.NewTokenService(userRepo, cacheSvc)
	authSvc := svcImpl.NewAuthService(userRepo, cacheSvc, tokenSvc)
	locationSvc := svcImpl.NewLocationService(locationRepo)
	roleSvc := svcImpl.NewRolesService(roleRepo)
	permissionSvc := svcImpl.NewPermissionsService(permissionRepo)

	controllers.NewSystemController(g, sysSvc)
	controllers.NewAuthController(g, authSvc, userSvc)
	controllers.NewCompanyController(g, acl, companySvc)
	controllers.NewUsersController(g, acl, companySvc, userSvc, locationSvc)
	controllers.NewLocationController(g, acl, locationSvc)
	controllers.NewRolesController(g, acl, roleSvc)
	controllers.NewPermissionsController(g, acl, permissionSvc)
}
