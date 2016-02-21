package main

import (
	"flag"
	"github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
)

var log = logrus.New()

var diskDrive = flag.String("drive", "/", "the drive location to use for stats")

func main() {
	flag.Parse()

	port := 8080

	log.Info("starting stats service...")
	log.Info("using disk drive ", *diskDrive, " for stats")

	router := NewRouter()

	log.Info("service started on port ", port)

	log.Error(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
