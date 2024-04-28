package alog

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

// 非格式化打印日志
func console(lv int, level, format string, a ...interface{}) {
	if lv >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintln(format, abc[:len(abc)-1])

		funcName, fileName, lineNo := getInfo(3)

		fmt.Printf("%s [%s] [%s=>%s:%d] %s", now, level, fileName, funcName, lineNo, msg)

	}
}

// 未知级别日志
func Unknown(format string, a ...interface{}) {
	//console(UNKNOWN, "Unknown", format, a...)
	if UNKNOWN >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintf("%s%s", format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Cyan("%s [%s] [%s=>%s:%d] %s", now, "Unknown", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Unknown", fileName, funcName, lineNo, msg)
		}

	}
}

// 用户级调试日志
func Debug(format string, a ...interface{}) {

	//console(DEBUG, "Debug", format, a...)
	if DEBUG >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintf("%s%s", format, abc[:len(abc)-1])

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.White("%s [%s] [%s=>%s:%d] %s", now, "Debug", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Debug", fileName, funcName, lineNo, msg)
		}
	}
}

// 用户级信息日志
func Info(format string, a ...interface{}) {

	//console(INFO, "Info", format, a...)
	if INFO >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintf("%s%s", format, abc[:len(abc)-1])
		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Blue("%s [%s] [%s=>%s:%d] %s", now, "Info", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Info", fileName, funcName, lineNo, msg)
		}
	}
}

// 用户级警告
func Warn(format string, a ...interface{}) {

	//console(WARN, "Warn", format, a...)
	if WARN >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintf("%s%s", format, abc[:len(abc)-1])

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Yellow("%s [%s] [%s=>%s:%d] %s", now, "Warn", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Warn", fileName, funcName, lineNo, msg)
		}
	}
}

// 用户级错误
func Error(format string, a ...interface{}) {

	//console(ERROR, "Error", format, a...)
	if ERROR >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintf("%s%s", format, abc[:len(abc)-1])

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Red("%s [%s] [%s=>%s:%d] %s", now, "Error", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Error", fileName, funcName, lineNo, msg)
		}
	}
}

// 致命错误
func Fatal(format string, a ...interface{}) {

	//console(FATAL, "Fatal", format, a...)
	if FATAL >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintf("%s%s", format, abc[:len(abc)-1])

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {

			color.Magenta("%s [%s] [%s=>%s:%d] %s", now, "Fatal", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Fatal", fileName, funcName, lineNo, msg)
		}
	}
}

// 系统级危险，比如权限出错，访问异常等
func Crit(format string, a ...interface{}) {

	//console(CRIT, "Crit", format, a...)
	if CRIT >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintf("%s%s", format, abc[:len(abc)-1])

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Magenta("%s [%s] [%s=>%s:%d] %s", now, "Crit", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Crit", fileName, funcName, lineNo, msg)
		}
	}
}

// 系统级警告，比如数据库访问异常，配置文件出错等
func Alrt(format string, a ...interface{}) {

	//console(ALRT, "Alrt", format, a...)
	if ALRT >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintf("%s%s", format, abc[:len(abc)-1])

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Magenta("%s [%s] [%s=>%s:%d] %s", now, "Alrt", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Alrt", fileName, funcName, lineNo, msg)
		}
	}
}

// 系统级紧急，比如磁盘出错，内存异常，网络不可用等
func Emer(format string, a ...interface{}) {

	//console(EMER, "Emer", format, a...)
	if EMER >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintf("%s%s", format, abc[:len(abc)-1])

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Magenta("%s [%s] [%s=>%s:%d] %s", now, "Emer", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Emer", fileName, funcName, lineNo, msg)
		}
	}
}

// 入侵警告
func Invade(format string, a ...interface{}) {

	//console(INVADE, "Invade", format, a...)
	if INVADE >= Level {
		now := time.Now().Format(TimeFormat)
		abc := fmt.Sprintln(a...)
		msg := fmt.Sprintf("%s%s", format, abc[:len(abc)-1])

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Green("%s [%s] [%s=>%s:%d] %s", now, "Invade", fileName, funcName, lineNo, msg)

		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Invade", fileName, funcName, lineNo, msg)
		}
	}
}

// 格式化打印日志
func consoleFormat(lv int, level, format string, a ...interface{}) {

	if lv >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, level, fileName, funcName, lineNo, msg)

	}
}

// Unknown Format
func Uf(format string, a ...interface{}) {

	//consoleFormat(UNKNOWN, "Unknown", format, a...)
	if UNKNOWN >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)

		funcName, fileName, lineNo := getInfo(2)

		if IsColor {
			color.Cyan("%s [%s] [%s=>%s:%d] %s", now, "Unknown", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Unknown", fileName, funcName, lineNo, msg)
		}
	}
}

// Debug Format
func Df(format string, a ...interface{}) {

	//consoleFormat(DEBUG, "Debug", format, a...)
	if DEBUG >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.White("%s [%s] [%s=>%s:%d] %s", now, "Debug", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Debug", fileName, funcName, lineNo, msg)
		}
	}
}

// Info Format
func If(format string, a ...interface{}) {

	//consoleFormat(INFO, "Info", format, a...)
	if INFO >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Blue("%s [%s] [%s=>%s:%d] %s", now, "Info", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Info", fileName, funcName, lineNo, msg)
		}
	}
}

// Warn Format
func Wf(format string, a ...interface{}) {

	//consoleFormat(WARN, "Warn", format, a...)
	if WARN >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Yellow("%s [%s] [%s=>%s:%d] %s", now, "Warn", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Warn", fileName, funcName, lineNo, msg)
		}
	}
}

// Error Format
func Errf(format string, a ...interface{}) {

	//consoleFormat(ERROR, "Error", format, a...)
	if ERROR >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Red("%s [%s] [%s=>%s:%d] %s", now, "Error", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Error", fileName, funcName, lineNo, msg)
		}
	}
}

// Fatal Format
func Ff(format string, a ...interface{}) {

	//consoleFormat(FATAL, "Fatal", format, a...)
	if FATAL >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)
		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Magenta("%s [%s] [%s=>%s:%d] %s", now, "Fatal", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Fatal", fileName, funcName, lineNo, msg)
		}
	}
}

// Crit Format  系统级危险，比如权限出错，访问异常等
func Cf(format string, a ...interface{}) {

	//consoleFormat(CRIT, "Crit", format, a...)
	if CRIT >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Magenta("%s [%s] [%s=>%s:%d] %s", now, "Crit", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Crit", fileName, funcName, lineNo, msg)
		}
	}
}

// Alrt Format  系统级警告，比如数据库访问异常，配置文件出错等
func Af(format string, a ...interface{}) {

	//consoleFormat(ALRT, "Alrt", format, a...)
	if ALRT >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Magenta("%s [%s] [%s=>%s:%d] %s", now, "Alrt", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Alrt", fileName, funcName, lineNo, msg)
		}
	}
}

// 系统级紧急，比如磁盘出错，内存异常，网络不可用等
func EmerF(format string, a ...interface{}) {

	//consoleFormat(EMER, "Emer", format, a...)
	if EMER >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Magenta("%s [%s] [%s=>%s:%d] %s", now, "Emer", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Emer", fileName, funcName, lineNo, msg)
		}
	}
}

// Invade Format，入侵警告
func InvadeF(format string, a ...interface{}) {

	//consoleFormat(INVADE, "Invade", format, a...)
	if INVADE >= Level {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format(TimeFormat)

		funcName, fileName, lineNo := getInfo(2)
		if IsColor {
			color.Green("%s [%s] [%s=>%s:%d] %s", now, "Invade", fileName, funcName, lineNo, msg)
		} else {
			fmt.Printf("%s [%s] [%s=>%s:%d] %s\n", now, "Invade", fileName, funcName, lineNo, msg)
		}
	}
}
