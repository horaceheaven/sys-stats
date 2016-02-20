package main
import (
	"github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
	"flag"
	"os/user"
	"os"
)

var log = logrus.New()
var usr, _ = user.Current()
var diskDrive = flag.String("drive", os.Getenv("STAT_DRIVE") || usr.HomeDir, "the drive location to use for stats")

func main() {
	flag.Parse()

	port := 8080

	log.Info("starting stats service...")
	log.Info("using disk drive ", *diskDrive, " for stats")

	router := NewRouter()

	log.Info("service started on port ", port)

	log.Error(http.ListenAndServe(":" + strconv.Itoa(port), router))
}