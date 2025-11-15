package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLC_Config_v1610 struct {
	dl_AM_RLC_v1610 DL_AM_RLC_v1610 `madatory`
}

func (ie *RLC_Config_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.dl_AM_RLC_v1610.Encode(w); err != nil {
		return utils.WrapError("Encode dl_AM_RLC_v1610", err)
	}
	return nil
}

func (ie *RLC_Config_v1610) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.dl_AM_RLC_v1610.Decode(r); err != nil {
		return utils.WrapError("Decode dl_AM_RLC_v1610", err)
	}
	return nil
}
