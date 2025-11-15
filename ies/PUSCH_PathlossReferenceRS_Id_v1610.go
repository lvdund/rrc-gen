package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_PathlossReferenceRS_Id_v1610 struct {
	Value uint64 `lb:maxNrofPUSCH_PathlossReferenceRSs,ub:maxNrofPUSCH_PathlossReferenceRSs_1_r16,madatory`
}

func (ie *PUSCH_PathlossReferenceRS_Id_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: maxNrofPUSCH_PathlossReferenceRSs, Ub: maxNrofPUSCH_PathlossReferenceRSs_1_r16}, false); err != nil {
		return utils.WrapError("Encode PUSCH_PathlossReferenceRS_Id_v1610", err)
	}
	return nil
}

func (ie *PUSCH_PathlossReferenceRS_Id_v1610) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: maxNrofPUSCH_PathlossReferenceRSs, Ub: maxNrofPUSCH_PathlossReferenceRSs_1_r16}, false); err != nil {
		return utils.WrapError("Decode PUSCH_PathlossReferenceRS_Id_v1610", err)
	}
	ie.Value = uint64(v)
	return nil
}
