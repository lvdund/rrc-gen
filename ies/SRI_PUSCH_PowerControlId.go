package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRI_PUSCH_PowerControlId struct {
	Value uint64 `lb:0,ub:maxNrofSRI_PUSCH_Mappings_1,madatory`
}

func (ie *SRI_PUSCH_PowerControlId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofSRI_PUSCH_Mappings_1}, false); err != nil {
		return utils.WrapError("Encode SRI_PUSCH_PowerControlId", err)
	}
	return nil
}

func (ie *SRI_PUSCH_PowerControlId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSRI_PUSCH_Mappings_1}, false); err != nil {
		return utils.WrapError("Decode SRI_PUSCH_PowerControlId", err)
	}
	ie.Value = uint64(v)
	return nil
}
