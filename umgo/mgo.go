// Package umgo provides ...
package umgo

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"time"

	"github.com/globalsign/mgo"
)

type ConnConfig struct {
	Hosts      []string      `json:"hosts"`
	DataBase   string        `json:"database"`
	UserName   string        `json:"username"`
	Password   string        `json:"password"`
	SSLCrtFile string        `json:"sslCrtFile"`
	Timeout    time.Duration `json:"timeout"`
}

var mgoSession *mgo.Session

func NewConfig(confFile string) (*ConnConfig, error) {
	configBytes, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	var connConfig ConnConfig
	err = json.Unmarshal(configBytes, &connConfig)
	return &connConfig, err
}

func NewSession(config *ConnConfig) (*mgo.Session, error) {
	var dialServer func(*mgo.ServerAddr) (net.Conn, error)
	if config.SSLCrtFile != "" {
		rootCrtBytes, err := ioutil.ReadFile(config.SSLCrtFile)
		if err != nil {
			return nil, err
		}
		rootCAs := x509.NewCertPool()
		ok := rootCAs.AppendCertsFromPEM(rootCrtBytes)
		if !ok {
			return nil, fmt.Errorf("Failed to parse root certificate from %s", config.SSLCrtFile)
		}
		tlsConfig := &tls.Config{
			RootCAs:            rootCAs,
			InsecureSkipVerify: true,
		}
		dialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), tlsConfig)
		}
	}
	dialInfo := &mgo.DialInfo{
		Addrs:      config.Hosts,
		Timeout:    config.Timeout * time.Second,
		Database:   config.DataBase,
		Username:   config.UserName,
		Password:   config.Password,
		DialServer: dialServer,
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	mgoSession = session
	return session, nil
}

func Close() {
	mgoSession.Close()
}

func Exec(f func(*mgo.Session)) {
	sessionCopy := mgoSession.Clone()
	defer func() {
		sessionCopy.Close()
	}()
	f(sessionCopy)
}
