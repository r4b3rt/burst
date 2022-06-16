package main

import (
	"flag"
	burst "github.com/fzdwx/burst/burst-client/client"
	"github.com/fzdwx/burst/burst-client/common"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

var (
	serverIp   = flag.String("sip", "localhost", "server ip")
	serverPort = flag.Int("sp", 10086, "server serverPort")
	token      = flag.String("t", "b92a205269d94d38808c3979615245eb", "your key, you can get it from server")
	usage      = flag.Bool("h", false, "help")
	debug      = flag.Bool("d", true, "log level use debug")
	serverAddr string
)

func init() {
	flag.Parse()
	if *usage {
		flag.Usage()
		os.Exit(0)
	}

	if strings.Compare(*token, "null") == 0 {
		log.Fatal("token is null")
		os.Exit(1)
	}

	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	log.SetFormatter(formatter)

	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	serverAddr = common.FormatToAddr(*serverIp, *serverPort)

	log.Infoln("log level:", common.WrapGreen(log.GetLevel().String()))
	log.Infoln("server address:", common.WrapGreen(serverAddr))
}

func main() {
	u := url.URL{Scheme: "ws", Host: serverAddr, Path: "/connect", RawQuery: "token=" + *token}
	client, resp, err := burst.Connect(u)
	if err != nil {
		body := resp.Body
		defer body.Close()
		data, _ := ioutil.ReadAll(body)
		log.Fatal(string(data))
	}
	defer client.Close()

	client.MountBinaryHandler(burst.HandlerBinaryData())

	go func() {
		client.React()
	}()

	select {}
}
