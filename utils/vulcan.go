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

func NewVulcanClient(bid string, port string, ttl time.Duration) *VulcandClient {
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

	_, body, err := gorequest.New().
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
	vc.UpsertBackend()
	go vc.Ping()

	ticker := time.NewTicker(vc.TTL / 2)
	go func() {
		for range ticker.C {
			vc.Ping()
		}
	}()
}
