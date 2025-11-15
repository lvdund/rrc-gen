package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	OtherConfig_delayBudgetReportingConfig_Choice_nothing uint64 = iota
	OtherConfig_delayBudgetReportingConfig_Choice_release
	OtherConfig_delayBudgetReportingConfig_Choice_setup
)

type OtherConfig_delayBudgetReportingConfig struct {
	Choice  uint64
	release uper.NULL `madatory`
	setup   *OtherConfig_delayBudgetReportingConfig_setup
}

func (ie *OtherConfig_delayBudgetReportingConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case OtherConfig_delayBudgetReportingConfig_Choice_release:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode release", err)
		}
	case OtherConfig_delayBudgetReportingConfig_Choice_setup:
		if err = ie.setup.Encode(w); err != nil {
			err = utils.WrapError("Encode setup", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *OtherConfig_delayBudgetReportingConfig) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case OtherConfig_delayBudgetReportingConfig_Choice_release:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode release", err)
		}
	case OtherConfig_delayBudgetReportingConfig_Choice_setup:
		ie.setup = new(OtherConfig_delayBudgetReportingConfig_setup)
		if err = ie.setup.Decode(r); err != nil {
			return utils.WrapError("Decode setup", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
