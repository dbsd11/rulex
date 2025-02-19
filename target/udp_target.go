package target

import (
	"net"
	"net/http"

	"github.com/hootrhino/rulex/common"
	"github.com/hootrhino/rulex/glogger"
	"github.com/hootrhino/rulex/typex"
	"github.com/hootrhino/rulex/utils"
)

/*
*
* 数据推到UDP
*
 */
type UdpTarget struct {
	typex.XStatus
	client     http.Client
	mainConfig common.HostConfig
	status     typex.SourceState
}

func NewUdpTarget(e typex.RuleX) typex.XTarget {
	udpt := new(UdpTarget)
	udpt.RuleEngine = e
	udpt.mainConfig = common.HostConfig{}
	udpt.status = typex.SOURCE_DOWN
	return udpt
}

func (udpt *UdpTarget) Init(outEndId string, configMap map[string]interface{}) error {
	udpt.PointId = outEndId

	if err := utils.BindSourceConfig(configMap, &udpt.mainConfig); err != nil {
		return err
	}

	return nil

}
func (udpt *UdpTarget) Start(cctx typex.CCTX) error {
	udpt.Ctx = cctx.Ctx
	udpt.CancelCTX = cctx.CancelCTX
	udpt.client = http.Client{}
	udpt.status = typex.SOURCE_UP
	glogger.GLogger.Info("UdpTarget started")
	return nil
}

func (udpt *UdpTarget) Test(outEndId string) bool {
	return true
}
func (udpt *UdpTarget) Enabled() bool {
	return udpt.Enable
}
func (udpt *UdpTarget) Reload() {

}
func (udpt *UdpTarget) Pause() {

}
func (udpt *UdpTarget) Status() typex.SourceState {
	return udpt.status

}
func (udpt *UdpTarget) To(data interface{}) (interface{}, error) {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP(udpt.mainConfig.Host),
		Port: udpt.mainConfig.Port,
	})
	if err != nil {
		return 0, err
	}
	defer socket.Close()
	switch t := data.(type) {
	case string:
		socket.Write([]byte(t))
	case []byte:
		socket.Write(t)
	default:
		glogger.GLogger.Error("unknown type:", t)
	}
	return 0, err
}

func (udpt *UdpTarget) Stop() {
	udpt.status = typex.SOURCE_STOP
	udpt.CancelCTX()
}
func (udpt *UdpTarget) Details() *typex.OutEnd {
	return udpt.RuleEngine.GetOutEnd(udpt.PointId)
}

/*
*
* 配置
*
 */
func (*UdpTarget) Configs() *typex.XConfig {
	return &typex.XConfig{}
}
