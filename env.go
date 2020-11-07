package ding

import "os"

const (
	envAgentId   = "DING_AGENT_ID"
	envAppKey    = "DING_APP_KEY"
	envAppSecret = "DING_APP_SECRET"
)

func init() {
	if os.Getenv(envAgentId) != "" {
		agentId = os.Getenv(envAgentId)
	}
	if os.Getenv(envAppKey) != "" {
		appKey = os.Getenv(envAppKey)
	}
	if os.Getenv(envAppSecret) != "" {
		appSecret = os.Getenv(envAppSecret)
	}
}
