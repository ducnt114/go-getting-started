package log

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	log.Info().Msg(msg)
}
