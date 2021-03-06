package main

import (
	"context"
	"os"
	"github.com/sirupsen/logrus"
	"github.com/rancher/norman/signal"
	"github.com/rancher/rancher/k8s"

)


func main() {
	logrus.Info("k8s is starting")
	ctx := signal.SigTermCancelContext(context.Background())
	_, ctx, _, err := k8s.GetConfig(ctx, "auto", "")
	if err != nil {
		logrus.Infof("create k8s failed %s", err)
		return
	}

	os.Unsetenv("KUBECONFIG")
	
	<-ctx.Done()
	
	logrus.Info(ctx.Err())
	
	return
}
