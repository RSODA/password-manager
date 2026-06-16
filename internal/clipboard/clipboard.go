package clipboard

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type command struct {
	name string
	args []string
}

func CopyText(text string) error {
	return write(text)
}

func Clear() error {
	return write("")
}

func write(text string) error {
	cmdSpec, err := commandForEnvironment()
	if err != nil {
		return err
	}

	cmd := exec.Command(cmdSpec.name, cmdSpec.args...)
	cmd.Stdin = strings.NewReader(text)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("%s failed: %w: %s", cmdSpec.name, err, strings.TrimSpace(string(output)))
	}

	return nil
}

func commandForEnvironment() (command, error) {
	switch runtime.GOOS {
	case "darwin":
		return existingCommand([]command{{name: "pbcopy"}})
	case "windows":
		return command{name: "clip"}, nil
	case "linux":
		if os.Getenv("WAYLAND_DISPLAY") != "" {
			if cmd, err := existingCommand([]command{{name: "wl-copy"}}); err == nil {
				return cmd, nil
			}
		}

		if os.Getenv("DISPLAY") != "" {
			return existingCommand([]command{
				{name: "xclip", args: []string{"-selection", "clipboard"}},
				{name: "xsel", args: []string{"--clipboard", "--input"}},
			})
		}

		return command{}, errors.New("нет DISPLAY или WAYLAND_DISPLAY; запустите с графической сессией или установите Xvfb")
	default:
		return command{}, fmt.Errorf("буфер обмена не поддерживается для %s", runtime.GOOS)
	}
}

func existingCommand(commands []command) (command, error) {
	var tried []string
	for _, cmd := range commands {
		tried = append(tried, cmd.name)
		if _, err := exec.LookPath(cmd.name); err == nil {
			return cmd, nil
		}
	}

	return command{}, fmt.Errorf("не найдена команда для буфера обмена: %s", strings.Join(tried, ", "))
}
