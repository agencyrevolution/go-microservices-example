package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/pborman/uuid"
)

type VulcandClient struct {
	Addr      string
	Version   string
	BackendId string
	ServerId  string
	ServerURL string
	TTL       time.Duration
}

func VulcandUpsertBackend(bid string) {
	vc := &VulcandClient{
		Addr:    os.Getenv("VULCAND_ADDR"),
		Version: os.Getenv("VULCAND_VER"),
	}

	reqBody := fmt.Sprintf(
		`{"Backend": {"Id":"%s", "Type":"http"}}`,
		bid,
	)

	gorequest.New().
		Post(vc.endpoint("backends")).
		Send(reqBody).
		End()
}

func VulcandUpsertFrontend(fid string, path string, bid string) {
	vc := &VulcandClient{
		Addr:    os.Getenv("VULCAND_ADDR"),
		Version: os.Getenv("VULCAND_VER"),
	}

	reqBody := fmt.Sprintf(
		`{"Frontend": {"Id":"%s", "Type": "http", "BackendId": "%s", "Route": "Path(\"%s\")"}}`,
		fid,
		bid,
		path,
	)

	gorequest.New().
		Post(vc.endpoint("frontends")).
		Send(reqBody).
		End()
}

func VulcandUpsertListener(lid string, scope string, addr string) {
	vc := &VulcandClient{
		Addr:    os.Getenv("VULCAND_ADDR"),
		Version: os.Getenv("VULCAND_VER"),
	}

	reqBody := fmt.Sprintf(
		`{"Listener":{"Id":"%s", "Protocol":"http", "Scope":"%s", "Address":{"Network":"tcp", "Address":"%s"}}}`,
		lid,
		scope,
		addr,
	)

	gorequest.New().
		Post(vc.endpoint("listeners")).
		Send(reqBody).
		End()
}

func NewVulcandClient(bid string, port string, ttl time.Duration) *VulcandClient {
	return &VulcandClient{
		Addr:      os.Getenv("VULCAND_ADDR"),
		Version:   os.Getenv("VULCAND_VER"),
		BackendId: bid,
		ServerId:  uuid.New(),
		ServerURL: fmt.Sprintf("http://%s:%s", os.Getenv("HOST"), port),
		TTL:       ttl,
	}
}

func (vc *VulcandClient) endpoint(params ...string) string {
	return fmt.Sprintf("%s/%s/%s", vc.Addr, vc.Version, strings.Join(params, "/"))
}

func (vc *VulcandClient) UpsertBackend() {
	reqBody := fmt.Sprintf(
		`{"Backend": {"Id":"%s", "Type":"http"}}`,
		vc.BackendId,
	)

	gorequest.New().
		Post(vc.endpoint("backends")).
		Send(reqBody).
		End()
}

func (vc *VulcandClient) Ping() {
	reqBody := fmt.Sprintf(
		`{"Server": {"Id":"%s", "URL":"%s"}, "TTL": "%s"}`,
		vc.ServerId,
		vc.ServerURL,
		vc.TTL.String(),
	)

	gorequest.New().
		Post(vc.endpoint("backends", vc.BackendId, "servers")).
		Send(reqBody).
		End()
}

func (vc *VulcandClient) KeepAlive() {
	go vc.Ping()

	ticker := time.NewTicker(vc.TTL / 2)
	go func() {
		for range ticker.C {
			vc.Ping()
		}
	}()
}
