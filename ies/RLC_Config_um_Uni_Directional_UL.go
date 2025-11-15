package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLC_Config_um_Uni_Directional_UL struct {
	ul_UM_RLC UL_UM_RLC `madatory`
}

func (ie *RLC_Config_um_Uni_Directional_UL) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ul_UM_RLC.Encode(w); err != nil {
		return utils.WrapError("Encode ul_UM_RLC", err)
	}
	return nil
}

func (ie *RLC_Config_um_Uni_Directional_UL) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ul_UM_RLC.Decode(r); err != nil {
		return utils.WrapError("Decode ul_UM_RLC", err)
	}
	return nil
}
