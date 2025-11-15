package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReportConfigToAddMod_reportConfig_Choice_nothing uint64 = iota
	ReportConfigToAddMod_reportConfig_Choice_reportConfigNR
	ReportConfigToAddMod_reportConfig_Choice_reportConfigInterRAT
	ReportConfigToAddMod_reportConfig_Choice_reportConfigNR_SL_r16
)

type ReportConfigToAddMod_reportConfig struct {
	Choice                uint64
	reportConfigNR        *ReportConfigNR
	reportConfigInterRAT  *ReportConfigInterRAT
	reportConfigNR_SL_r16 *ReportConfigNR_SL_r16
}

func (ie *ReportConfigToAddMod_reportConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigToAddMod_reportConfig_Choice_reportConfigNR:
		if err = ie.reportConfigNR.Encode(w); err != nil {
			err = utils.WrapError("Encode reportConfigNR", err)
		}
	case ReportConfigToAddMod_reportConfig_Choice_reportConfigInterRAT:
		if err = ie.reportConfigInterRAT.Encode(w); err != nil {
			err = utils.WrapError("Encode reportConfigInterRAT", err)
		}
	case ReportConfigToAddMod_reportConfig_Choice_reportConfigNR_SL_r16:
		if err = ie.reportConfigNR_SL_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode reportConfigNR_SL_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ReportConfigToAddMod_reportConfig) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigToAddMod_reportConfig_Choice_reportConfigNR:
		ie.reportConfigNR = new(ReportConfigNR)
		if err = ie.reportConfigNR.Decode(r); err != nil {
			return utils.WrapError("Decode reportConfigNR", err)
		}
	case ReportConfigToAddMod_reportConfig_Choice_reportConfigInterRAT:
		ie.reportConfigInterRAT = new(ReportConfigInterRAT)
		if err = ie.reportConfigInterRAT.Decode(r); err != nil {
			return utils.WrapError("Decode reportConfigInterRAT", err)
		}
	case ReportConfigToAddMod_reportConfig_Choice_reportConfigNR_SL_r16:
		ie.reportConfigNR_SL_r16 = new(ReportConfigNR_SL_r16)
		if err = ie.reportConfigNR_SL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode reportConfigNR_SL_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
