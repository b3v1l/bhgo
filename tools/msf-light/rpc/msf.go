package rpc

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/vmihailenco/msgpack"
)

type loginRes struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type logoutReq struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	LogoutToken string
}

type logoutRes struct {
	Result string `msgpack:"result"`
}

type sessionListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type SessionListRes struct {
	ID          uint32 `msgpack:",omitempty"`
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload"`
	Description string `msgpack:"desc"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"session_host"`
	SessionPort int    `msgpack:"session_port"`
	Username    string `msgpack:"username"`
	UUID        string `msgpack:"uuid"`
	ExploitUUID string `msgpack:"exploit_uuid"`
}

type Metasploit struct {
	host  string
	user  string
	pass  string
	token string
}

//handle login
func New(host, user, pass string) *Metasploit {

	msf := &Metasploit{
		host: host,
		user: user,
		pass: pass,
	}
	return msf
}

//Deserialization / serialisation

func (msf *Metasploit) send(req interface{}, res interface{}) error {

	buf := new(bytes.Buffer)
	msgpack.NewEncoder(buf).Encode(*req)
	dest := fmt.Scanf("http://%s/api", msf.host)
	resp, err := http.Post(dest, "binary/message-pack", buf)
	if err != nil {
		return err
	}

	if err := msgpack.NewEncoder(resp.body).Decode(&res); err != nil {
		return err
	}
	return nil
}
