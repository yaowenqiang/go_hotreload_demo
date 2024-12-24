package main
import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Starting service...")
	logrus.Debug("Detailed debug information here.")
}

