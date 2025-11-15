package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCSetupComplete_v1700_IEs_onboardingRequest_r17_Enum_true uper.Enumerated = 0
)

type RRCSetupComplete_v1700_IEs_onboardingRequest_r17 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *RRCSetupComplete_v1700_IEs_onboardingRequest_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode RRCSetupComplete_v1700_IEs_onboardingRequest_r17", err)
	}
	return nil
}

func (ie *RRCSetupComplete_v1700_IEs_onboardingRequest_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode RRCSetupComplete_v1700_IEs_onboardingRequest_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
