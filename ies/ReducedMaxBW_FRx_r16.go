package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReducedMaxBW_FRx_r16 struct {
	reducedBW_DL_r16 ReducedAggregatedBandwidth `madatory`
	reducedBW_UL_r16 ReducedAggregatedBandwidth `madatory`
}

func (ie *ReducedMaxBW_FRx_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.reducedBW_DL_r16.Encode(w); err != nil {
		return utils.WrapError("Encode reducedBW_DL_r16", err)
	}
	if err = ie.reducedBW_UL_r16.Encode(w); err != nil {
		return utils.WrapError("Encode reducedBW_UL_r16", err)
	}
	return nil
}

func (ie *ReducedMaxBW_FRx_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.reducedBW_DL_r16.Decode(r); err != nil {
		return utils.WrapError("Decode reducedBW_DL_r16", err)
	}
	if err = ie.reducedBW_UL_r16.Decode(r); err != nil {
		return utils.WrapError("Decode reducedBW_UL_r16", err)
	}
	return nil
}
