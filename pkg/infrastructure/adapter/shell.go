package infrastructure

import (
	"bytes"
	"context"
	"os/exec"
	"strings"

	me "github.com/husamettinarabaci/go-resthell/core/domain/model/entity"
	mo "github.com/husamettinarabaci/go-resthell/core/domain/model/object"
	mp "github.com/husamettinarabaci/go-resthell/pkg/infrastructure/mapper"
)

type ShellAdapter struct {
}

func NewShellAdapter() ShellAdapter {
	adapter := ShellAdapter{}

	return adapter
}

func (a ShellAdapter) ExecuteCommand(ctx context.Context, commandRequest me.CommandRequest) (mo.Response, error) {
	commandMapper := mp.FromCommandObject(commandRequest.Command)
	var err error
	response := mo.Response{}
	if len(commandMapper.Values) == 0 {
		err = mo.ErrInvalidInput
		return response, err
	}
	stdout, stderr, err := a.Shellout(commandMapper.Values)
	stdoutSplt := strings.Split(stdout, "\n")
	stderrSplt := strings.Split(stderr, "\n")
	response = mo.NewResponse(stdoutSplt, stderrSplt)
	return response, err
}

func (a ShellAdapter) IsCommandExists(ctx context.Context, commandRequest me.CommandRequest) (bool, error) {
	commandMapper := mp.FromCommandObject(commandRequest.Command)
	if len(commandMapper.Values) == 0 {
		return false, mo.ErrInvalidInput
	}
	var whichCommand []string
	whichCommand = append(whichCommand, "which")
	whichCommand = append(whichCommand, commandMapper.Values[0])
	stdout, _, err := a.Shellout(whichCommand)
	if err != nil {
		return false, err
	}
	if len(stdout) > 0 {
		return true, nil
	}
	return false, nil
}

func (a ShellAdapter) Shellout(command []string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
