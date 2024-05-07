package simplelogger

import (
	"os"

	"github.com/rs/zerolog"
)

func LogWriteFile(dest string, filename string, i uint, size uint32) {
	file, err := os.OpenFile(
		dest+filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	logger := zerolog.New(file).With().Timestamp().Logger()

	logger.Info().Uint("id", i).Uint32("size", size).Msg("Write file")
}
