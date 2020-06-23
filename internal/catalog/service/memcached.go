package service

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/andremedeiros/loon/internal/executer"
	"github.com/andremedeiros/loon/internal/process"
)

type Memcached struct{}

func (m *Memcached) String() string {
	return "Memcached"
}

func (m *Memcached) Identifier() string {
	return "memcached"
}

func (m *Memcached) Initialize(_ executer.Executer, _, _ string, _ ...executer.Option) error {
	return nil
}

func (m *Memcached) Versions() map[string][]string {
	return map[string][]string{
		"default": {"memcached", "1.6.6"},
		"latest":  {"memcached", "1.6.6"},

		"1.6.6": {"memcached", "1.6.6"},
		"1.6.5": {"memcached", "1.6.5"},
	}
}

func (m *Memcached) Environ(ipaddr, vdpath string) []string {
	return []string{
		fmt.Sprintf("MEMCACHED_URL=%s:11211", ipaddr),
	}
}

func (m *Memcached) IsHealthy(ipaddr, _ string) bool {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:11211", ipaddr), 100*time.Millisecond)
	return err == nil
}

func (m *Memcached) Start(exe executer.Executer, ipaddr, vdpath string, opts ...executer.Option) error {
	pidPath := filepath.Join(vdpath, "pids", "memcached.pid")
	_, err := exe.Execute([]string{
		"memcached",
		"--daemon",
		"--port=11211",
		fmt.Sprintf("--listen=%s", ipaddr),
		fmt.Sprintf("--pidfile=%s", pidPath),
	}, opts...)
	return err
}

func (m *Memcached) Stop(exe executer.Executer, _, vdpath string, _ ...executer.Option) error {
	pidPath := filepath.Join(vdpath, "pids", "memcached.pid")
	p, err := process.FromPidFile(pidPath)
	if err != nil {
		return nil
	}
	_ = os.Remove(pidPath)
	return p.Signal(os.Interrupt)
}
