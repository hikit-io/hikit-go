package rtm

import (
	"fmt"
	"net/http"
)

const (
	BaseUrl    = "https://api.agora.io/dev/v2/project/%s"
	PeerUrl    = "/rtm/users/%s/peer_messages"
	ChannelUrl = "/rtm/users/%s/channel_messages"
)

func GenBaseUrl(appId string) string {
	return fmt.Sprintf(BaseUrl, appId)
}

func GenPeerUrl(userId Any) string {
	return fmt.Sprintf(PeerUrl, userId)
}

func GenChannelUrl(userId Any) string {
	return fmt.Sprintf(ChannelUrl, userId)
}

type TResult string

const (
	ResultSuccess TResult = "success"
	ResultFailed          = "failed"
)

type TCode string

const (
	CodeMessageSent      TCode = "message_sent"
	CodeMessageDelivered       = "message_delivered"
	CodeMessageOffline         = "message_offline"
)

type TStatus int

const (
	StatusOK                  = http.StatusOK
	StatusBadRequest          = http.StatusBadRequest
	StatusUnauthorized        = http.StatusUnauthorized
	StatusRequestTimeout      = http.StatusRequestTimeout
	StatusTooManyRequests     = http.StatusTooManyRequests
	StatusInternalServerError = http.StatusInternalServerError
)

type Response struct {
	Result    TResult `json:"result,omitempty"`
	RequestId string  `json:"request_id,omitempty"`
	Code      TCode   `json:"code,omitempty"`
}

type PeerMsg struct {
	Destination               string
	EnableOfflineMessaging    bool
	EnableHistoricalMessaging bool
	Payload                   string
}

type ChanMsg struct {
	ChannelName               string
	EnableHistoricalMessaging bool
	Payload                   string
}
