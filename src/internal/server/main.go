package server

import apiv1 "github.com/deynego-sergey/gbroker/pkg/gen"

type QServer struct {
	apiv1.UnimplementedQBrokerServer
}

func NewQServer() *QServer {

}
