package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportConfig_reportQuantity_Choice_nothing uint64 = iota
	CSI_ReportConfig_reportQuantity_Choice_none
	CSI_ReportConfig_reportQuantity_Choice_cri_RI_PMI_CQI
	CSI_ReportConfig_reportQuantity_Choice_cri_RI_i1
	CSI_ReportConfig_reportQuantity_Choice_cri_RI_i1_CQI
	CSI_ReportConfig_reportQuantity_Choice_cri_RI_CQI
	CSI_ReportConfig_reportQuantity_Choice_cri_RSRP
	CSI_ReportConfig_reportQuantity_Choice_ssb_Index_RSRP
	CSI_ReportConfig_reportQuantity_Choice_cri_RI_LI_PMI_CQI
)

type CSI_ReportConfig_reportQuantity struct {
	Choice            uint64
	none              uper.NULL `madatory`
	cri_RI_PMI_CQI    uper.NULL `madatory`
	cri_RI_i1         uper.NULL `madatory`
	cri_RI_i1_CQI     *CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI
	cri_RI_CQI        uper.NULL `madatory`
	cri_RSRP          uper.NULL `madatory`
	ssb_Index_RSRP    uper.NULL `madatory`
	cri_RI_LI_PMI_CQI uper.NULL `madatory`
}

func (ie *CSI_ReportConfig_reportQuantity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportQuantity_Choice_none:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode none", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RI_PMI_CQI:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode cri_RI_PMI_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RI_i1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode cri_RI_i1", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RI_i1_CQI:
		if err = ie.cri_RI_i1_CQI.Encode(w); err != nil {
			err = utils.WrapError("Encode cri_RI_i1_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RI_CQI:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode cri_RI_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RSRP:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode cri_RSRP", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_ssb_Index_RSRP:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode ssb_Index_RSRP", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RI_LI_PMI_CQI:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode cri_RI_LI_PMI_CQI", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ReportConfig_reportQuantity) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportConfig_reportQuantity_Choice_none:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode none", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RI_PMI_CQI:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode cri_RI_PMI_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RI_i1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode cri_RI_i1", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RI_i1_CQI:
		ie.cri_RI_i1_CQI = new(CSI_ReportConfig_reportQuantity_cri_RI_i1_CQI)
		if err = ie.cri_RI_i1_CQI.Decode(r); err != nil {
			return utils.WrapError("Decode cri_RI_i1_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RI_CQI:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode cri_RI_CQI", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RSRP:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode cri_RSRP", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_ssb_Index_RSRP:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode ssb_Index_RSRP", err)
		}
	case CSI_ReportConfig_reportQuantity_Choice_cri_RI_LI_PMI_CQI:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode cri_RI_LI_PMI_CQI", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
