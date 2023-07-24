package main

import (
	"os"
	"os/exec"
)

// здесь приходится отключать линтер, иначе ругается на неиспользованное значение
// его нельзя вывести, тк иначе выводы в test.sh не совпадут
// сигнатуру функции решил не менять, а просто заглушить

func RunCmd(cmd []string, env Environment) (returnCode int) { //nolint:all
	// теперь проверяю значение NeedRemove
	for k, v := range env {
		if v.NeedRemove {
			os.Unsetenv(k)
		} else {
			os.Setenv(k, v.Value)
		}
	}

	// здесь линтер ругается на небезопасное выполнение команды
	command := exec.Command(cmd[0], cmd[1:]...) //nolint:all
	command.Env = os.Environ()
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr
	if exitError := command.Run(); exitError != nil {
		// пытаемся извлечь код ошибки, если ошибка иная, то просто выведем 1
		if err, ok := exitError.(*exec.ExitError); ok { //nolint:all
			returnCode := err.ExitCode()
			return returnCode
		}
		return 1
	}
	return 0
}
