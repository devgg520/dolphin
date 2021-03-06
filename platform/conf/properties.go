package conf

import (
	"fmt"
	"strings"

	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
)

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	viper.AddConfigPath("conf")
	viper.SetDefault("app.name", "dolphin")
	viper.SetDefault("app.version", "1.0")
	viper.SetDefault("app.mode", "release")
	viper.SetDefault("app.viper", false)
	viper.SetDefault("http.hash", "FF61A573-82FC-478B-9AEF-93D6F506DE9A")
	viper.SetDefault("http.port", "8081")
	viper.SetDefault("http.prefix", "/api")
	viper.SetDefault("http.static", "static")
	viper.SetDefault("grpc.port", "9081")
	viper.SetDefault("oauth.id", model.DefaultClient.Client.String)
	viper.SetDefault("oauth.secret", model.DefaultClient.Secret.String)
	viper.SetDefault("oauth.login", "/static/web/login.html")
	viper.SetDefault("oauth.affirm", "/static/web/affirm.html")
	viper.SetDefault("db.driver", "mysql")
	viper.SetDefault("db.dataSource", "root:111111@/dolphin?charset=utf8&parseTime=True&loc=Local")
	viper.SetDefault("db.connMaxLifetime", 10)
	viper.SetDefault("db.maxIdleConns", 15)
	viper.SetDefault("db.maxOpenConns", 50)
	viper.SetDefault("rd.dataSource", ":@localhost:6379/0")
	viper.SetDefault("dir.app", "app")
	viper.SetDefault("dir.doc", "doc")
	viper.SetDefault("dir.sql", "sql")
	viper.SetDefault("dir.sqlmap", "sqlmap")
	viper.SetDefault("dir.script", "script")
	viper.SetDefault("dir.log", "log")
	viper.SetDefault("dir.util", "util")
	viper.SetDefault("dir.model", "model")
	viper.SetDefault("dir.srv", "srv")
	viper.SetDefault("dir.rpc", "rpc")
	viper.SetDefault("swag.license.name", "Apache 2.0")
	viper.SetDefault("swag.license.url", "http://www.apache.org/licenses/LICENSE-2.0.html")
	viper.SetDefault("swag.securitydefinitions.oauth2.accessCode", "OAuth2AccessCode")
	viper.SetDefault("swag.tokenUrl", "/api/sys/cas/token")
	viper.SetDefault("swag.authorizationUrl", "/api/sys/cas/authorize")
	viper.SetDefault("swag.scope.read", "Grants read access")
	viper.SetDefault("swag.scope.write", "Grants write access")
	viper.SetDefault("swag.scope.admin", "Grants read and write access to administrative information")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Warn("configuration file not found")
	}
	if strings.TrimSpace(viper.GetString("oauth.server")) == "" {
		viper.SetDefault("oauth.server", fmt.Sprintf("http://localhost:%v", viper.GetString("http.port")))
	}
	if strings.TrimSpace(viper.GetString("oauth.cli")) == "" {
		viper.SetDefault("oauth.cli", fmt.Sprintf("http://localhost:%v", viper.GetString("http.port")))
	}
	if strings.TrimSpace(viper.GetString("host")) == "" {
		viper.SetDefault("host", fmt.Sprintf("localhost:%v", viper.GetString("http.port")))
	}
	if viper.GetBool("app.viper") {
		if err := viper.WriteConfig(); err != nil {
			logrus.Warn("failed to save configuration file")
		}
	}
}
