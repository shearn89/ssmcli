package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	log "github.com/sirupsen/logrus"
)

var (
	ssmSvc *ssm.SSM
)

func init() {
	ssmSvc = ssm.New(sess)
}

func GetSessions() map[string]string {
	sessionMap := map[string]string{}
	activeSessons, err := ssmSvc.DescribeSessions(&ssm.DescribeSessionsInput{
		State: aws.String("Active"),
	})
	if err != nil {
		log.WithError(err).Error("failed to list sessions")
		return sessionMap
	}
	for _, session := range activeSessons.Sessions {
		sessionId := aws.StringValue(session.SessionId)
		sessionTarget := aws.StringValue(session.Target)
		sessionMap[sessionTarget] = sessionId
	}
	return sessionMap
}
