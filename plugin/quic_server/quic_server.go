package quic_server

import (
	"strconv"
	"time"

	"github.com/hootrhino/rulex/typex"
	"github.com/hootrhino/rulex/utils"
	"github.com/quic-go/quic-go/http3"
	"gopkg.in/ini.v1"
)

type QuicServer struct {
	uuid string
	config *_quicServerConfig
	server *http3.Server
	ruleEngine *typex.RuleX
}

type _quicServerConfig struct {
	Enable bool   `ini:"enable"`
	Host   string `ini:"host"`
	Port   int    `ini:"port"`
	DtlsPublicKey string `ini:"dtls_public_key"`
	DtlsPrivateKey string `ini:"dtls_private_key"`
}

var StartedTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

func NewQuicServer() *QuicServer {
	return &QuicServer {
		uuid: "QuicServer",
	}
}

func (dm *QuicServer) Init(config *ini.Section) error {
	var quicServerConfig _quicServerConfig
	if err := utils.InIMapToStruct(config, &quicServerConfig); err != nil {
		return err
	}
	dm.config = &quicServerConfig
	return nil
}

func (dm *QuicServer) Start(r typex.RuleX) error {
	server := &http3.Server {
		Addr:    dm.config.Host + ":" +  strconv.Itoa(dm.config.Port),
		Handler: setupGenericHandler(),
	}
	dm.server = server
	dm.ruleEngine = &r
	if len(dm.config.DtlsPublicKey) > 0 && len(dm.config.DtlsPrivateKey) > 0 {
		return server.ListenAndServeTLS(dm.config.DtlsPublicKey, dm.config.DtlsPrivateKey)
	}
	return server.ListenAndServe()
}

func (dm *QuicServer) Stop() error {
	if dm.server != nil {
		return dm.server.Close()
	}
	return nil
}

func (hh *QuicServer) PluginMetaInfo() typex.XPluginMetaInfo {
	return typex.XPluginMetaInfo{
		UUID:     hh.uuid,
		Name:     "QuicServer",
		Version:  "0.0.1",
		Homepage: "https://hootrhino.github.io",
		HelpLink: "https://hootrhino.github.io",
		Author:   "wwhai",
		Email:    "cnwwhai@gmail.com",
		License:  "MIT",
	}
}

/*
*
* 服务调用接口
*
 */
func (cs *QuicServer) Service(arg typex.ServiceArg) typex.ServiceResult {
	return typex.ServiceResult{}
}
