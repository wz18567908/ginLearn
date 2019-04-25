package services

import (
	def "ginLearn/myweb/models"
    db "ginLearn/myweb/db"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func SubmitLSFJob(jobName, cwd, cmds string) (string, error) {
	var err error
	args := []string{""}
	if jobName != "" {
		args = append(args, fmt.Sprintf("JOB_NAME=%s", jobName))
	}
	if cwd != "" {
		args = append(args, fmt.Sprintf("CURRENT_WORK_DIR=%s", cwd))
	}
	if cmds != "" {
		args = append(args, fmt.Sprintf("COMMAND=%s", cmds))
	}
	cmd := exec.Command("/bin/bash", "-c", "/apps/clustertech/chess_pro/gui/etc/application/published/generic/test.sh")
	fmt.Println("before: ", cmd.Env)

	ChangeYourCmdEnvironment(cmd, args)
	fmt.Println("after: ", cmd.Env)
	output, err := cmd.Output()
    job := def.LSFJobReq{
        JobID : string(output),
        JobName : jobName,
        WorkDir : cwd,
        Command : cmds,
    }
    db.InsertJob(&job)
	return string(output), err
}


func ChangeYourCmdEnvironment(cmd * exec.Cmd, cmdEnv []string) error {
	env := os.Environ()

	for _, e := range env {
		i := strings.Index(e, "=")
		if i > 0 && (e[:i] == "ENV_NAME") {
			// do yourself
		} else {
			cmdEnv = append(cmdEnv, e)
		}
	}
	cmd.Env = cmdEnv

	return nil
}
