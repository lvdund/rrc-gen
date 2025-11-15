package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_PositionQCL_Relation_r16_Enum_n1 uper.Enumerated = 0
	SSB_PositionQCL_Relation_r16_Enum_n2 uper.Enumerated = 1
	SSB_PositionQCL_Relation_r16_Enum_n4 uper.Enumerated = 2
	SSB_PositionQCL_Relation_r16_Enum_n8 uper.Enumerated = 3
)

type SSB_PositionQCL_Relation_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SSB_PositionQCL_Relation_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SSB_PositionQCL_Relation_r16", err)
	}
	return nil
}

func (ie *SSB_PositionQCL_Relation_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SSB_PositionQCL_Relation_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
