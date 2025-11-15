package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_groupBasedBeamReporting_Choice_nothing uint64 = iota
	CSI_ReportConfig_groupBasedBeamReporting_Choice_enabled
	CSI_ReportConfig_groupBasedBeamReporting_Choice_disabled
)

type CSI_ReportConfig_groupBasedBeamReporting struct {
	Choice   uint64
	enabled  uper.NULL `madatory`
	disabled *CSI_ReportConfig_groupBasedBeamReporting_disabled
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_groupBasedBeamReporting_Choice_enabled:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode enabled", err)
		}
	case CSI_ReportConfig_groupBasedBeamReporting_Choice_disabled:
		if err = ie.disabled.Encode(w); err != nil {
			err = utils.WrapError("Encode disabled", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ReportConfig_groupBasedBeamReporting) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_groupBasedBeamReporting_Choice_enabled:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode enabled", err)
		}
	case CSI_ReportConfig_groupBasedBeamReporting_Choice_disabled:
		ie.disabled = new(CSI_ReportConfig_groupBasedBeamReporting_disabled)
		if err = ie.disabled.Decode(r); err != nil {
			return utils.WrapError("Decode disabled", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
