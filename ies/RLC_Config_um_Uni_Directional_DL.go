package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLC_Config_um_Uni_Directional_DL struct {
	dl_UM_RLC DL_UM_RLC `madatory`
}

func (ie *RLC_Config_um_Uni_Directional_DL) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.dl_UM_RLC.Encode(w); err != nil {
		return utils.WrapError("Encode dl_UM_RLC", err)
	}
	return nil
}

func (ie *RLC_Config_um_Uni_Directional_DL) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.dl_UM_RLC.Decode(r); err != nil {
		return utils.WrapError("Decode dl_UM_RLC", err)
	}
	return nil
}
