package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type P0_PUSCH_AlphaSetId struct {
	Value uint64 `lb:0,ub:maxNrofP0_PUSCH_AlphaSets_1,madatory`
}

func (ie *P0_PUSCH_AlphaSetId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofP0_PUSCH_AlphaSets_1}, false); err != nil {
		return utils.WrapError("Encode P0_PUSCH_AlphaSetId", err)
	}
	return nil
}

func (ie *P0_PUSCH_AlphaSetId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofP0_PUSCH_AlphaSets_1}, false); err != nil {
		return utils.WrapError("Decode P0_PUSCH_AlphaSetId", err)
	}
	ie.Value = uint64(v)
	return nil
}
