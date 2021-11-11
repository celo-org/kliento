package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

type VictoropsReport struct {
	MessageType       string `json:"message_type"`
	EntityId          string `json:"entity_id"`
	EntityDisplayName string `json:"entity_display_name"`
	GethGitCommit     string `json:"geth_commit"`
	StateMessage      string `json:"state_message"`
	PodName           string `json:"pod_name"`
	ClusterName       string `json:"cluster_name,omitempty"`
	StartBlock        int    `json:"start_block,omitempty"`
	EndBlock          int    `json:"end_block,omitempty"`
	PeersCount        int    `json:"peer_count,omitempty"`
	Peers             string `json:"peers,omitempty"`
	DebugStacks       string `json:"debug_stack,omitempty"`
	ErrorMessage      string `json:"error_message,omitempty"`
	JobCount          string `json:"job_count,omitempty"`
}

type SlackReport struct {
	Text string `json:"text"`
}

func parseTimestamp(t uint64) time.Time {
	tm := time.Unix(int64(t), 0)
	return tm
}

func getenv(envvar, fallback string) string {
	value := os.Getenv(envvar)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func reportVictorops(report VictoropsReport) error {
	victoropsEndpoint := os.Getenv("VICTOROPS_ENDPOINT")
	matched, _ := regexp.MatchString("https://alert.victorops.com/.+", victoropsEndpoint)
	if !(matched) {
		log.Crit("VICTOROPS_ENDPOINT env var missmatch")
	}

	body, err := json.Marshal(report)
	if err != nil {
		return err
	}
	resp, err := http.Post(victoropsEndpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("error sending POST request to victorops")
	}
	log.Info("Report sent to victorops")
	return nil
}

func reportSlack(report SlackReport) error {
	slackEndpoint := os.Getenv("SLACK_ENDPOINT")
	matched, _ := regexp.MatchString("https://hooks.slack.com/.+", slackEndpoint)
	if !(matched) {
		log.Crit("SLACK_ENDPOINT env var missmatch")
	}

	body, err := json.Marshal(report)
	if err != nil {
		log.Error("Error marshaling report")
		return err
	}
	resp, err := http.Post(slackEndpoint, "application/json", bytes.NewBuffer(body))
	log.Info("Trace slack respone", "resp", resp)
	if err != nil {
		return err
	} else if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("error sending POST request to slack")
	}
	return nil
}
