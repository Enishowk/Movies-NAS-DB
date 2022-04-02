package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hirochachacha/go-smb2"
)

func ConnectSMB() (*smb2.Session, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:445", os.Getenv("NAS_IP")))
	if err != nil {
		return nil, err
	}

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     os.Getenv("NAS_USERNAME"),
			Password: os.Getenv("NAS_PASSWORD"),
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		defer conn.Close()
		return nil, err
	}

	return s, nil
}
