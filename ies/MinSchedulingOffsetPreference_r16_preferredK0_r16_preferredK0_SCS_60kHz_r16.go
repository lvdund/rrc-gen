package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_60kHz_r16_Enum_sl2  uper.Enumerated = 0
	MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_60kHz_r16_Enum_sl4  uper.Enumerated = 1
	MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_60kHz_r16_Enum_sl8  uper.Enumerated = 2
	MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_60kHz_r16_Enum_sl12 uper.Enumerated = 3
)

type MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_60kHz_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_60kHz_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_60kHz_r16", err)
	}
	return nil
}

func (ie *MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_60kHz_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_60kHz_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
