package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_beamSwitchTiming_v1710 struct {
	scs_480kHz *MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz `optional`
	scs_960kHz *MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_960kHz `optional`
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_480kHz != nil, ie.scs_960kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_480kHz != nil {
		if err = ie.scs_480kHz.Encode(w); err != nil {
			return utils.WrapError("Encode scs_480kHz", err)
		}
	}
	if ie.scs_960kHz != nil {
		if err = ie.scs_960kHz.Encode(w); err != nil {
			return utils.WrapError("Encode scs_960kHz", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_v1710) Decode(r *uper.UperReader) error {
	var err error
	var scs_480kHzPresent bool
	if scs_480kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_960kHzPresent bool
	if scs_960kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_480kHzPresent {
		ie.scs_480kHz = new(MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz)
		if err = ie.scs_480kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_480kHz", err)
		}
	}
	if scs_960kHzPresent {
		ie.scs_960kHz = new(MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_960kHz)
		if err = ie.scs_960kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_960kHz", err)
		}
	}
	return nil
}
