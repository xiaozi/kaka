package main

import (
	// "io/ioutil"
	"log"
	"os"
	"os/exec"
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

func (b *Browser) Snapshot(url string, path string) error {
	var err error = nil
	cmd := exec.Command(
		"casperjs",
		"--ssl-protocol", "=", "any",
		"--ignore-ssl-errors", "=", "true",
		"file.js",
	)
	env := os.Environ()
	env = append(env, "LC_CTYPE=en_US.UTF-8", "PATHEXT=/usr/local/casperjs/bin:/usr/local/phantomjs/bin")
	cmd.Env = env
	// outPipe, _ := cmd.StdoutPipe()
	// errPipe, _ := cmd.StderrPipe()

	startErr := cmd.Start()

	if startErr != nil {
		log.Fatal(startErr)
		return startErr
	}

	// stdOutput, _ := ioutil.ReadAll(outPipe)

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(30 * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill:", err)
		}
		<-done
		log.Println("process killed")
	case err := <-done:
		if err != nil {
			log.Printf("process exit with %v", err)
		}
	}
	
	return err
}
