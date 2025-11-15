package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentCarrierInfo_r16 struct {
	servCellIndex_r16 ServCellIndex                                         `madatory`
	servCellInfo_r16  UplinkTxDirectCurrentCarrierInfo_r16_servCellInfo_r16 `madatory`
}

func (ie *UplinkTxDirectCurrentCarrierInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.servCellIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode servCellIndex_r16", err)
	}
	if err = ie.servCellInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode servCellInfo_r16", err)
	}
	return nil
}

func (ie *UplinkTxDirectCurrentCarrierInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.servCellIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode servCellIndex_r16", err)
	}
	if err = ie.servCellInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode servCellInfo_r16", err)
	}
	return nil
}
