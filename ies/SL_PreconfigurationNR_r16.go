package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PreconfigurationNR_r16 struct {
	sidelinkPreconfigNR_r16 SidelinkPreconfigNR_r16 `madatory`
}

func (ie *SL_PreconfigurationNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sidelinkPreconfigNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sidelinkPreconfigNR_r16", err)
	}
	return nil
}

func (ie *SL_PreconfigurationNR_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sidelinkPreconfigNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sidelinkPreconfigNR_r16", err)
	}
	return nil
}
