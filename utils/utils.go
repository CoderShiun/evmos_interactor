package utils

import (
	"bytes"
	"errors"
	"math/big"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func HomePath() (string, error) {

	user, err := user.Current()

	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support
	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()

}

func homeUnix() (string, error) {

	// First prefer the HOME environmental variable

	if home := os.Getenv("HOME"); home != "" {

		return home, nil

	}

	// If that fails, try the shell

	var stdout bytes.Buffer

	cmd := exec.Command("sh", "-c", "eval echo ~$USER")

	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {

		return "", err

	}

	result := strings.TrimSpace(stdout.String())

	if result == "" {

		return "", errors.New("blank output when reading home directory")

	}

	return result, nil

}

func homeWindows() (string, error) {

	drive := os.Getenv("HOMEDRIVE")

	path := os.Getenv("HOMEPATH")

	home := drive + path

	if drive == "" || path == "" {

		home = os.Getenv("USERPROFILE")

	}

	if home == "" {

		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")

	}

	return home, nil
}

func GetBigInt(amount string) *big.Int {
	newAmount, err := strconv.Atoi(amount)
	if err != nil {
		log.Fatal(err)
	}

	return big.NewInt(int64(newAmount))
}
