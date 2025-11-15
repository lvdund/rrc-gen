package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReleasePreference_r16_preferredRRC_State_r16_Enum_idle           uper.Enumerated = 0
	ReleasePreference_r16_preferredRRC_State_r16_Enum_inactive       uper.Enumerated = 1
	ReleasePreference_r16_preferredRRC_State_r16_Enum_connected      uper.Enumerated = 2
	ReleasePreference_r16_preferredRRC_State_r16_Enum_outOfConnected uper.Enumerated = 3
)

type ReleasePreference_r16_preferredRRC_State_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ReleasePreference_r16_preferredRRC_State_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ReleasePreference_r16_preferredRRC_State_r16", err)
	}
	return nil
}

func (ie *ReleasePreference_r16_preferredRRC_State_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ReleasePreference_r16_preferredRRC_State_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
