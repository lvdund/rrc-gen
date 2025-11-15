package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_reportQuantity_r16_Choice_nothing uint64 = iota
	CSI_ReportConfig_reportQuantity_r16_Choice_cri_SINR_r16
	CSI_ReportConfig_reportQuantity_r16_Choice_ssb_Index_SINR_r16
)

type CSI_ReportConfig_reportQuantity_r16 struct {
	Choice             uint64
	cri_SINR_r16       uper.NULL `madatory`
	ssb_Index_SINR_r16 uper.NULL `madatory`
}

func (ie *CSI_ReportConfig_reportQuantity_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportQuantity_r16_Choice_cri_SINR_r16:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode cri_SINR_r16", err)
		}
	case CSI_ReportConfig_reportQuantity_r16_Choice_ssb_Index_SINR_r16:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode ssb_Index_SINR_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ReportConfig_reportQuantity_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportQuantity_r16_Choice_cri_SINR_r16:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode cri_SINR_r16", err)
		}
	case CSI_ReportConfig_reportQuantity_r16_Choice_ssb_Index_SINR_r16:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode ssb_Index_SINR_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
