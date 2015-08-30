package main

import (
	//	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

type Browser struct {
	ViewPort  *ViewPort
	UserAgent string
	Proxy     *Proxy
}

type ViewPort struct {
	Width  int
	Height int
}

type Proxy struct {
	Host string
	Port int
}

func NewMacBrowser() *Browser {
	return &Browser{
		ViewPort:  &ViewPort{Width: 1280, Height: 800},
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.94 Safari/537.36",
		Proxy:     nil,
	}
}

func NewIPhoneBrowser() *Browser {
	return &Browser{
		ViewPort:  &ViewPort{Width: 640, Height: 960},
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X; en-us) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53",
		Proxy:     nil,
	}
}

func (b *Browser) Snapshot(url string, path string, timeout int) error {
	var err error = nil
	// os.Readlink("/proc/self/exe")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)
	cmd := exec.Command(
		"casperjs",
		"--ssl-protocol=any",
		"--ignore-ssl-errors=true",
		filepath.Join(dir, "scripts", "snapshot.js"),
		"--width="+strconv.Itoa(b.ViewPort.Width),
		"--height="+strconv.Itoa(b.ViewPort.Height),
		"--useragent="+b.UserAgent,
		"--url="+url,
		"--path="+path,
	)
	env := os.Environ()
	env = append(
		env,
		"LC_CTYPE=en_US.UTF-8",
		"PATH=/usr/local/node/bin:/usr/local/phantomjs/bin",
	)
	cmd.Env = env

	log.Print(cmd.Args)

	//	outPipe, _ := cmd.StdoutPipe()
	//	errPipe, _ := cmd.StderrPipe()

	startErr := cmd.Start()

	if startErr != nil {
		log.Fatal(startErr)
		return startErr
	}

	//	stdOutput, _ := ioutil.ReadAll(outPipe)
	//	errOutput, _ := ioutil.ReadAll(errPipe)

	//	log.Print(string(stdOutput))
	//	log.Fatal(string(errOutput))

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill:", err)
		}
		<-done
		log.Println("process killed")
	case err := <-done:
		log.Printf("process exit with %v\n", err)
	}

	return nil
}
