package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_ue_PowerClass_v1700_Enum_pc5 uper.Enumerated = 0
	BandNR_ue_PowerClass_v1700_Enum_pc6 uper.Enumerated = 1
	BandNR_ue_PowerClass_v1700_Enum_pc7 uper.Enumerated = 2
)

type BandNR_ue_PowerClass_v1700 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *BandNR_ue_PowerClass_v1700) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode BandNR_ue_PowerClass_v1700", err)
	}
	return nil
}

func (ie *BandNR_ue_PowerClass_v1700) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode BandNR_ue_PowerClass_v1700", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
