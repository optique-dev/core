package optique

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

type LogLevel string

const (
	InfoLevel  LogLevel = "info"
	ErrorLevel LogLevel = "error"
	DebugLevel LogLevel = "debug"
)

func Init() error {
	root, err := FindOptiqueJson()
	if err != nil {
		return err
	}
	manifest, err := ReadProjectManifestAt(root + "/" + PROJECT_MANIFEST)
	os.Setenv("OPTIQUE_SERVICE", manifest.Name)
	return nil
}

func Info(msg string) {
	Log(&LogOptions{
		Level:   InfoLevel,
		Service: RetrieveServiceName(),
		Message: msg,
	})
}

func RetrieveServiceName() string {
	service, ok := os.LookupEnv("OPTIQUE_SERVICE")
	if !ok {
		service = "optique"
	}
	return service
}

func Error(msg string) {
	Log(&LogOptions{
		Level:   ErrorLevel,
		Service: RetrieveServiceName(),
		Message: msg,
	})
}

func Debug(msg string) {
	Log(&LogOptions{
		Level:   DebugLevel,
		Service: RetrieveServiceName(),
		Message: msg,
	})
}

type LogOptions struct {
	Level   LogLevel
	Service string
	Message string
}

func Log(options *LogOptions) {
	formated_service := fmt.Sprintf(" @%s ", options.Service)
	switch options.Level {
	case InfoLevel:
		color.New(color.FgWhite, color.BgCyan).Print("INFO")
		color.New(color.FgCyan).Print(formated_service)
		color.New(color.FgGray.Light()).Print(options.Message)
	case ErrorLevel:
		color.New(color.FgWhite, color.BgRed).Print("ERROR")
		color.New(color.FgRed).Print(formated_service)
		color.New(color.FgGray.Light()).Print(options.Message)
	case DebugLevel:
		color.New(color.FgWhite, color.BgBlue).Print("DEBUG")
		color.New(color.FgBlue).Print(formated_service)
		color.New(color.FgGray.Light()).Print(options.Message)
	}
	fmt.Println("")
}
