package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetupComplete_v1700_IEs struct {
	onboardingRequest_r17 *RRCSetupComplete_v1700_IEs_onboardingRequest_r17 `optional`
	nonCriticalExtension  interface{}                                       `optional`
}

func (ie *RRCSetupComplete_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.onboardingRequest_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.onboardingRequest_r17 != nil {
		if err = ie.onboardingRequest_r17.Encode(w); err != nil {
			return utils.WrapError("Encode onboardingRequest_r17", err)
		}
	}
	return nil
}

func (ie *RRCSetupComplete_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var onboardingRequest_r17Present bool
	if onboardingRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if onboardingRequest_r17Present {
		ie.onboardingRequest_r17 = new(RRCSetupComplete_v1700_IEs_onboardingRequest_r17)
		if err = ie.onboardingRequest_r17.Decode(r); err != nil {
			return utils.WrapError("Decode onboardingRequest_r17", err)
		}
	}
	return nil
}
