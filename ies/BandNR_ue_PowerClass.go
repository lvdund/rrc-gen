package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_ue_PowerClass_Enum_pc1 uper.Enumerated = 0
	BandNR_ue_PowerClass_Enum_pc2 uper.Enumerated = 1
	BandNR_ue_PowerClass_Enum_pc3 uper.Enumerated = 2
	BandNR_ue_PowerClass_Enum_pc4 uper.Enumerated = 3
)

type BandNR_ue_PowerClass struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *BandNR_ue_PowerClass) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode BandNR_ue_PowerClass", err)
	}
	return nil
}

func (ie *BandNR_ue_PowerClass) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode BandNR_ue_PowerClass", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
