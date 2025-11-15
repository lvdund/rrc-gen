package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReleasePreferenceConfig_r16_connectedReporting_Enum_true uper.Enumerated = 0
)

type ReleasePreferenceConfig_r16_connectedReporting struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *ReleasePreferenceConfig_r16_connectedReporting) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode ReleasePreferenceConfig_r16_connectedReporting", err)
	}
	return nil
}

func (ie *ReleasePreferenceConfig_r16_connectedReporting) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode ReleasePreferenceConfig_r16_connectedReporting", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
