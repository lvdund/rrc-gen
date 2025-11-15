package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReportConfigNR_SL_r16_reportType_r16_Choice_nothing uint64 = iota
	ReportConfigNR_SL_r16_reportType_r16_Choice_periodical_r16
	ReportConfigNR_SL_r16_reportType_r16_Choice_eventTriggered_r16
)

type ReportConfigNR_SL_r16_reportType_r16 struct {
	Choice             uint64
	periodical_r16     *PeriodicalReportConfigNR_SL_r16
	eventTriggered_r16 *EventTriggerConfigNR_SL_r16
}

func (ie *ReportConfigNR_SL_r16_reportType_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigNR_SL_r16_reportType_r16_Choice_periodical_r16:
		if err = ie.periodical_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode periodical_r16", err)
		}
	case ReportConfigNR_SL_r16_reportType_r16_Choice_eventTriggered_r16:
		if err = ie.eventTriggered_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode eventTriggered_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ReportConfigNR_SL_r16_reportType_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigNR_SL_r16_reportType_r16_Choice_periodical_r16:
		ie.periodical_r16 = new(PeriodicalReportConfigNR_SL_r16)
		if err = ie.periodical_r16.Decode(r); err != nil {
			return utils.WrapError("Decode periodical_r16", err)
		}
	case ReportConfigNR_SL_r16_reportType_r16_Choice_eventTriggered_r16:
		ie.eventTriggered_r16 = new(EventTriggerConfigNR_SL_r16)
		if err = ie.eventTriggered_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eventTriggered_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
