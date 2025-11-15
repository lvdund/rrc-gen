package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_ue_PowerClass_v1610_Enum_pc1dot5 uper.Enumerated = 0
)

type BandNR_ue_PowerClass_v1610 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *BandNR_ue_PowerClass_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode BandNR_ue_PowerClass_v1610", err)
	}
	return nil
}

func (ie *BandNR_ue_PowerClass_v1610) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode BandNR_ue_PowerClass_v1610", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
