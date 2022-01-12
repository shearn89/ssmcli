package internal

import (
	"github.com/aws/aws-sdk-go/aws/session"
	log "github.com/sirupsen/logrus"
	"os"
	// "os/exec"
	"syscall"
)

const (
	DocumentShell      string = ""
	DocumentForwarding string = "AWS-StartPortForwardingSession"
	Reason             string = "connecting via ssmcli"
)

var (
	sess *session.Session
)

func RunSSMShell(instance string) {
	/* It would be great to do this all natively, wiring up the WS connection,
	but sadly it's too much effort for now!
	*/

	err := syscall.Exec("/opt/homebrew/bin/aws", []string{"aws", "ssm", "start-session", "--target", instance}, os.Environ())
	if err != nil {
		log.WithError(err).Error("error whilst running ssm")
	}
}

func RunSSMResume(sessionId string) {
	err := syscall.Exec("/opt/homebrew/bin/aws", []string{"aws", "ssm", "resume-session", "--session-id", sessionId}, os.Environ())
	if err != nil {
		log.WithError(err).Error("error whilst running ssm")
	}
}
