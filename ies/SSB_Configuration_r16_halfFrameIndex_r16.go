package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_Configuration_r16_halfFrameIndex_r16_Enum_zero uper.Enumerated = 0
	SSB_Configuration_r16_halfFrameIndex_r16_Enum_one  uper.Enumerated = 1
)

type SSB_Configuration_r16_halfFrameIndex_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SSB_Configuration_r16_halfFrameIndex_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SSB_Configuration_r16_halfFrameIndex_r16", err)
	}
	return nil
}

func (ie *SSB_Configuration_r16_halfFrameIndex_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SSB_Configuration_r16_halfFrameIndex_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
