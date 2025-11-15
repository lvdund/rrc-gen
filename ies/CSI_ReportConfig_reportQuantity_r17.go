package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_reportQuantity_r17_Choice_nothing uint64 = iota
	CSI_ReportConfig_reportQuantity_r17_Choice_cri_RSRP_Index_r17
	CSI_ReportConfig_reportQuantity_r17_Choice_ssb_Index_RSRP_Index_r17
	CSI_ReportConfig_reportQuantity_r17_Choice_cri_SINR_Index_r17
	CSI_ReportConfig_reportQuantity_r17_Choice_ssb_Index_SINR_Index_r17
)

type CSI_ReportConfig_reportQuantity_r17 struct {
	Choice                   uint64
	cri_RSRP_Index_r17       uper.NULL `madatory`
	ssb_Index_RSRP_Index_r17 uper.NULL `madatory`
	cri_SINR_Index_r17       uper.NULL `madatory`
	ssb_Index_SINR_Index_r17 uper.NULL `madatory`
}

func (ie *CSI_ReportConfig_reportQuantity_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportQuantity_r17_Choice_cri_RSRP_Index_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode cri_RSRP_Index_r17", err)
		}
	case CSI_ReportConfig_reportQuantity_r17_Choice_ssb_Index_RSRP_Index_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode ssb_Index_RSRP_Index_r17", err)
		}
	case CSI_ReportConfig_reportQuantity_r17_Choice_cri_SINR_Index_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode cri_SINR_Index_r17", err)
		}
	case CSI_ReportConfig_reportQuantity_r17_Choice_ssb_Index_SINR_Index_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode ssb_Index_SINR_Index_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ReportConfig_reportQuantity_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportQuantity_r17_Choice_cri_RSRP_Index_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode cri_RSRP_Index_r17", err)
		}
	case CSI_ReportConfig_reportQuantity_r17_Choice_ssb_Index_RSRP_Index_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode ssb_Index_RSRP_Index_r17", err)
		}
	case CSI_ReportConfig_reportQuantity_r17_Choice_cri_SINR_Index_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode cri_SINR_Index_r17", err)
		}
	case CSI_ReportConfig_reportQuantity_r17_Choice_ssb_Index_SINR_Index_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode ssb_Index_SINR_Index_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
