package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyD_amplitudeScalingType_Enum_wideband           uper.Enumerated = 0
	DummyD_amplitudeScalingType_Enum_widebandAndSubband uper.Enumerated = 1
)

type DummyD_amplitudeScalingType struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *DummyD_amplitudeScalingType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode DummyD_amplitudeScalingType", err)
	}
	return nil
}

func (ie *DummyD_amplitudeScalingType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode DummyD_amplitudeScalingType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
