package setup

import "sync"

type IAppConfig interface {
	// GetChannels - get channels list
	GetChannels() []string
	// GetMaxFileSize - maximal file size
	GetMaxFileSize() int64
	// GetBaseFolder - storage folder
	GetBaseFolder() string
}

type appConfig struct{}

func (a *appConfig) GetChannels() []string {
	//TODO implement me
	panic("implement me")
}

func (a *appConfig) GetMaxFileSize() int64 {
	//TODO implement me
	panic("implement me")
}

func (a *appConfig) GetBaseFolder() string {
	//TODO implement me
	panic("implement me")
}

// current config
var (
	_conf IAppConfig
	_once sync.Once
)

func NewConfig() IAppConfig {
	_once.Do(func() {
		_conf = &appConfig{}
	})
	return _conf
}
