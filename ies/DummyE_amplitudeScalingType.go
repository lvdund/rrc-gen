package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyE_amplitudeScalingType_Enum_wideband           uper.Enumerated = 0
	DummyE_amplitudeScalingType_Enum_widebandAndSubband uper.Enumerated = 1
)

type DummyE_amplitudeScalingType struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *DummyE_amplitudeScalingType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode DummyE_amplitudeScalingType", err)
	}
	return nil
}

func (ie *DummyE_amplitudeScalingType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode DummyE_amplitudeScalingType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
