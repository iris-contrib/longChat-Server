package consts

import (
	"fmt"
)

const (
	IdServiceAddress  = "service.id.address"
	IdServiceStartIdx = "service.id.startidx"
	IdServiceStep     = "service.id.step"

	MongoDbName = "mongodb.dbname"
	MongoDbAddr = "mongodb.addr"
	MongoDbUser = "mongodb.user"
	MongoDbPsw  = "mongodb.psw"

	ApiServiceAddress    = "service.api.address"
	ApiServiceStaticPath = "service.api.staticpath"

	MsgServiceAddress = "service.msg.address"
	RouterServiceAddress = "service.router.address"

	ErrorLogPath  = "log.error.path"
	AccessLogPath = "log.access.path"

	TlsEnable = "security.tls.enable"

	RedisAddress  = "redis.address"
	RedisPassword = "redis.password"

	SessionCookieName = "session.cookie"
	SessionPrefix     = "session.prefix"
)

func ErrGetConfigFailed(configName string, err error) string {
	return fmt.Sprintf("get config(%s) failed!err:=%s\n", configName, err.Error())
}

func ErrDialRemoteServiceFailed(addr string, err error) string {
	return fmt.Sprintf("dial to server(%s) failed!err:=%s\n", addr, err.Error())
}

func ErrRPCCallFailed(service string, function string, err error) string {
	return fmt.Sprintf("rpc call(%s,%s) failed!err:=%s\n", service, function, err.Error())
}
