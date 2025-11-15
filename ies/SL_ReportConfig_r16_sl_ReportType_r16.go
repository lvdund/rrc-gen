package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_ReportConfig_r16_sl_ReportType_r16_Choice_nothing uint64 = iota
	SL_ReportConfig_r16_sl_ReportType_r16_Choice_sl_Periodical_r16
	SL_ReportConfig_r16_sl_ReportType_r16_Choice_sl_EventTriggered_r16
)

type SL_ReportConfig_r16_sl_ReportType_r16 struct {
	Choice                uint64
	sl_Periodical_r16     *SL_PeriodicalReportConfig_r16
	sl_EventTriggered_r16 *SL_EventTriggerConfig_r16
}

func (ie *SL_ReportConfig_r16_sl_ReportType_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_ReportConfig_r16_sl_ReportType_r16_Choice_sl_Periodical_r16:
		if err = ie.sl_Periodical_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode sl_Periodical_r16", err)
		}
	case SL_ReportConfig_r16_sl_ReportType_r16_Choice_sl_EventTriggered_r16:
		if err = ie.sl_EventTriggered_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode sl_EventTriggered_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_ReportConfig_r16_sl_ReportType_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_ReportConfig_r16_sl_ReportType_r16_Choice_sl_Periodical_r16:
		ie.sl_Periodical_r16 = new(SL_PeriodicalReportConfig_r16)
		if err = ie.sl_Periodical_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_Periodical_r16", err)
		}
	case SL_ReportConfig_r16_sl_ReportType_r16_Choice_sl_EventTriggered_r16:
		ie.sl_EventTriggered_r16 = new(SL_EventTriggerConfig_r16)
		if err = ie.sl_EventTriggered_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_EventTriggered_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
