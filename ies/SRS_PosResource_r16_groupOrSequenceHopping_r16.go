package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PosResource_r16_groupOrSequenceHopping_r16_Enum_neither         uper.Enumerated = 0
	SRS_PosResource_r16_groupOrSequenceHopping_r16_Enum_groupHopping    uper.Enumerated = 1
	SRS_PosResource_r16_groupOrSequenceHopping_r16_Enum_sequenceHopping uper.Enumerated = 2
)

type SRS_PosResource_r16_groupOrSequenceHopping_r16 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *SRS_PosResource_r16_groupOrSequenceHopping_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode SRS_PosResource_r16_groupOrSequenceHopping_r16", err)
	}
	return nil
}

func (ie *SRS_PosResource_r16_groupOrSequenceHopping_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode SRS_PosResource_r16_groupOrSequenceHopping_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
