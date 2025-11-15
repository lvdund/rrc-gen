package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LoggedMeasurementConfiguration_r16_IEs_reportType_Choice_nothing uint64 = iota
	LoggedMeasurementConfiguration_r16_IEs_reportType_Choice_periodical
	LoggedMeasurementConfiguration_r16_IEs_reportType_Choice_eventTriggered
)

type LoggedMeasurementConfiguration_r16_IEs_reportType struct {
	Choice         uint64
	periodical     *LoggedPeriodicalReportConfig_r16
	eventTriggered *LoggedEventTriggerConfig_r16
}

func (ie *LoggedMeasurementConfiguration_r16_IEs_reportType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LoggedMeasurementConfiguration_r16_IEs_reportType_Choice_periodical:
		if err = ie.periodical.Encode(w); err != nil {
			err = utils.WrapError("Encode periodical", err)
		}
	case LoggedMeasurementConfiguration_r16_IEs_reportType_Choice_eventTriggered:
		if err = ie.eventTriggered.Encode(w); err != nil {
			err = utils.WrapError("Encode eventTriggered", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *LoggedMeasurementConfiguration_r16_IEs_reportType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LoggedMeasurementConfiguration_r16_IEs_reportType_Choice_periodical:
		ie.periodical = new(LoggedPeriodicalReportConfig_r16)
		if err = ie.periodical.Decode(r); err != nil {
			return utils.WrapError("Decode periodical", err)
		}
	case LoggedMeasurementConfiguration_r16_IEs_reportType_Choice_eventTriggered:
		ie.eventTriggered = new(LoggedEventTriggerConfig_r16)
		if err = ie.eventTriggered.Decode(r); err != nil {
			return utils.WrapError("Decode eventTriggered", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
