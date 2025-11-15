package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasResultRelay_r17 struct {
	cellIdentity_r17        CellAccessRelatedInfo `madatory`
	sl_RelayUE_Identity_r17 SL_SourceIdentity_r17 `madatory`
	sl_MeasResult_r17       SL_MeasResult_r16     `madatory`
}

func (ie *SL_MeasResultRelay_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cellIdentity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode cellIdentity_r17", err)
	}
	if err = ie.sl_RelayUE_Identity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_RelayUE_Identity_r17", err)
	}
	if err = ie.sl_MeasResult_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_MeasResult_r17", err)
	}
	return nil
}

func (ie *SL_MeasResultRelay_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cellIdentity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode cellIdentity_r17", err)
	}
	if err = ie.sl_RelayUE_Identity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_RelayUE_Identity_r17", err)
	}
	if err = ie.sl_MeasResult_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_MeasResult_r17", err)
	}
	return nil
}
