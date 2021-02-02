package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"go-game-server/proto"
	"go-game-server/proto/proto2"

	"golang.org/x/net/websocket"
)

type NetDataConn struct {
	Connection *websocket.Conn
	MD5        string
}

func (this *NetDataConn) PullFromClient() {
	for {
		var content string
		if err := websocket.Message.Receive(this.Connection, &content); err != nil {
			break
		}
		if len(content) == 0 {
			break
		}
		this.SyncMessageFun(content)
	}
	return
}

func (this *NetDataConn) SyncMessageFun(content string) {
	fmt.Println("content: ", content)

	protocolData, err := Json2map(content)
	if err != nil {
		fmt.Println("解析失败: ", err)
	}
	this.HandleCltProtocol(protocolData["protocol"], protocolData["protocol2"], protocolData)
}

func (this *NetDataConn) HandleCltProtocol(protocol interface{}, protocol2 interface{}, protocolData map[string]interface{}) {
	switch protocol {
	case float64(proto.GameDataProto):
		{
			this.HandleCltProtocol2(protocol2, protocolData)
		}
	case float64(proto.GameDataDBProto):
		{
		}
	default:
		panic("主协议不存在")
	}
}

func (this *NetDataConn) HandleCltProtocol2(protocol2 interface{}, protocolData map[string]interface{}) {
	switch protocol2 {
	case float64(proto2.C2SPlayerLoginProto2):
		{
			this.PlayerLogin(protocolData)
		}
	default:
		panic("子协议不存在")
	}

	return
}

func (this *NetDataConn) PlayerLogin(protocolData map[string]interface{}) {
	if protocolData["code"] == nil {
		panic("主协议 1，子协议 1，登录功能数据错误")
	}
	code := protocolData["code"].(string)
	fmt.Println("code: ", code)

	data := proto2.S2CPlayerLogin{
		Protocol:   proto.GameDataProto,
		Protoco2:   proto2.S2CPlayerLoginProto2,
		PlayerData: nil,
	}
	this.PlayerSendMessage(data)
	return
}

func (this *NetDataConn) PlayerSendMessage(msg interface{}) {
	bs, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("bs: ", string(bs[0:]))
	err = websocket.JSON.Send(this.Connection, bs)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	return
}

func Json2map(data string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	return result, nil
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
