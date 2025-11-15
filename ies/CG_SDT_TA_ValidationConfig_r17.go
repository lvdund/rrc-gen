package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_SDT_TA_ValidationConfig_r17 struct {
	cg_SDT_RSRP_ChangeThreshold_r17 CG_SDT_TA_ValidationConfig_r17_cg_SDT_RSRP_ChangeThreshold_r17 `madatory`
}

func (ie *CG_SDT_TA_ValidationConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cg_SDT_RSRP_ChangeThreshold_r17.Encode(w); err != nil {
		return utils.WrapError("Encode cg_SDT_RSRP_ChangeThreshold_r17", err)
	}
	return nil
}

func (ie *CG_SDT_TA_ValidationConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cg_SDT_RSRP_ChangeThreshold_r17.Decode(r); err != nil {
		return utils.WrapError("Decode cg_SDT_RSRP_ChangeThreshold_r17", err)
	}
	return nil
}
