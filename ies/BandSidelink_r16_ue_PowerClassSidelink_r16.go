package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandSidelink_r16_ue_PowerClassSidelink_r16_Enum_pc2    uper.Enumerated = 0
	BandSidelink_r16_ue_PowerClassSidelink_r16_Enum_pc3    uper.Enumerated = 1
	BandSidelink_r16_ue_PowerClassSidelink_r16_Enum_spare6 uper.Enumerated = 2
	BandSidelink_r16_ue_PowerClassSidelink_r16_Enum_spare5 uper.Enumerated = 3
	BandSidelink_r16_ue_PowerClassSidelink_r16_Enum_spare4 uper.Enumerated = 4
	BandSidelink_r16_ue_PowerClassSidelink_r16_Enum_spare3 uper.Enumerated = 5
	BandSidelink_r16_ue_PowerClassSidelink_r16_Enum_spare2 uper.Enumerated = 6
	BandSidelink_r16_ue_PowerClassSidelink_r16_Enum_spare1 uper.Enumerated = 7
)

type BandSidelink_r16_ue_PowerClassSidelink_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *BandSidelink_r16_ue_PowerClassSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode BandSidelink_r16_ue_PowerClassSidelink_r16", err)
	}
	return nil
}

func (ie *BandSidelink_r16_ue_PowerClassSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode BandSidelink_r16_ue_PowerClassSidelink_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
