package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_Gap_r17 struct {
	musim_GapId_r17   MUSIM_GapId_r17   `madatory`
	musim_GapInfo_r17 MUSIM_GapInfo_r17 `madatory`
}

func (ie *MUSIM_Gap_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.musim_GapId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode musim_GapId_r17", err)
	}
	if err = ie.musim_GapInfo_r17.Encode(w); err != nil {
		return utils.WrapError("Encode musim_GapInfo_r17", err)
	}
	return nil
}

func (ie *MUSIM_Gap_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.musim_GapId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode musim_GapId_r17", err)
	}
	if err = ie.musim_GapInfo_r17.Decode(r); err != nil {
		return utils.WrapError("Decode musim_GapInfo_r17", err)
	}
	return nil
}
