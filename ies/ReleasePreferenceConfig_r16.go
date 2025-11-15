package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReleasePreferenceConfig_r16 struct {
	releasePreferenceProhibitTimer_r16 ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16 `madatory`
	connectedReporting                 *ReleasePreferenceConfig_r16_connectedReporting                `optional`
}

func (ie *ReleasePreferenceConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.connectedReporting != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.releasePreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode releasePreferenceProhibitTimer_r16", err)
	}
	if ie.connectedReporting != nil {
		if err = ie.connectedReporting.Encode(w); err != nil {
			return utils.WrapError("Encode connectedReporting", err)
		}
	}
	return nil
}

func (ie *ReleasePreferenceConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var connectedReportingPresent bool
	if connectedReportingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.releasePreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode releasePreferenceProhibitTimer_r16", err)
	}
	if connectedReportingPresent {
		ie.connectedReporting = new(ReleasePreferenceConfig_r16_connectedReporting)
		if err = ie.connectedReporting.Decode(r); err != nil {
			return utils.WrapError("Decode connectedReporting", err)
		}
	}
	return nil
}
