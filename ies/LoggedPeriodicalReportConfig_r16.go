package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LoggedPeriodicalReportConfig_r16 struct {
	loggingInterval_r16 LoggingInterval_r16 `madatory`
}

func (ie *LoggedPeriodicalReportConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.loggingInterval_r16.Encode(w); err != nil {
		return utils.WrapError("Encode loggingInterval_r16", err)
	}
	return nil
}

func (ie *LoggedPeriodicalReportConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.loggingInterval_r16.Decode(r); err != nil {
		return utils.WrapError("Decode loggingInterval_r16", err)
	}
	return nil
}
