package app

import (
	"fmt"

	"github.com/photo-pixels/platform/server"

	"github.com/photo-pixels/user-account/internal/service/jwt_helper"
	"github.com/photo-pixels/user-account/internal/service/session_manager"
	"github.com/photo-pixels/user-account/internal/user_case/auth"
)

const (
	// ServerConfigName конфиг сервера
	ServerConfigName = "server"
	// SessionManagerName конфиг менеджера сессии
	SessionManagerName = "session_manager"
	// JwtHelperName данные для jwt
	JwtHelperName = "jwt_helper"
	// AuthName данные для авторизации
	AuthName = "auth"
)

func (a *App) getServerConfig() (server.Config, error) {

	var config server.Config
	err := a.cfgProvider.PopulateByKey(ServerConfigName, &config)
	if err != nil {
		return server.Config{}, fmt.Errorf("PopulateByKey: %w", err)
	}

	return config, nil
}

func (a *App) getAuthConfig() (auth.Config, error) {
	var config auth.Config
	err := a.cfgProvider.PopulateByKey(AuthName, &config)
	if err != nil {
		return auth.Config{}, fmt.Errorf("PopulateByKey: %w", err)
	}
	return config, nil
}

func (a *App) getSessionManagerConfig() (session_manager.Config, error) {
	var config session_manager.Config
	err := a.cfgProvider.PopulateByKey(SessionManagerName, &config)
	if err != nil {
		return session_manager.Config{}, fmt.Errorf("PopulateByKey: %w", err)
	}
	return config, nil
}

func (a *App) getJwtHelperConfig() (jwt_helper.Config, error) {
	var config jwt_helper.Config
	err := a.cfgProvider.PopulateByKey(JwtHelperName, &config)
	if err != nil {
		return jwt_helper.Config{}, fmt.Errorf("PopulateByKey: %w", err)
	}
	return config, nil
}
