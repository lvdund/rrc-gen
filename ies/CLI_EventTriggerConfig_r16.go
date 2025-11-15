package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CLI_EventTriggerConfig_r16 struct {
	eventId_r16        CLI_EventTriggerConfig_r16_eventId_r16      `madatory`
	reportInterval_r16 ReportInterval                              `madatory,ext`
	reportAmount_r16   CLI_EventTriggerConfig_r16_reportAmount_r16 `madatory,ext`
	maxReportCLI_r16   int64                                       `lb:1,ub:maxCLI_Report_r16,madatory,ext`
}

func (ie *CLI_EventTriggerConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.eventId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode eventId_r16", err)
	}
	return nil
}

func (ie *CLI_EventTriggerConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.eventId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode eventId_r16", err)
	}
	return nil
}
