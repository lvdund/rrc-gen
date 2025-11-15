package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_beamSwitchTiming_r16 struct {
	scs_60kHz_r16  *MIMO_ParametersPerBand_beamSwitchTiming_r16_scs_60kHz_r16  `optional`
	scs_120kHz_r16 *MIMO_ParametersPerBand_beamSwitchTiming_r16_scs_120kHz_r16 `optional`
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_60kHz_r16 != nil, ie.scs_120kHz_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_60kHz_r16 != nil {
		if err = ie.scs_60kHz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scs_60kHz_r16", err)
		}
	}
	if ie.scs_120kHz_r16 != nil {
		if err = ie.scs_120kHz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scs_120kHz_r16", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_r16) Decode(r *uper.UperReader) error {
	var err error
	var scs_60kHz_r16Present bool
	if scs_60kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_120kHz_r16Present bool
	if scs_120kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_60kHz_r16Present {
		ie.scs_60kHz_r16 = new(MIMO_ParametersPerBand_beamSwitchTiming_r16_scs_60kHz_r16)
		if err = ie.scs_60kHz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode scs_60kHz_r16", err)
		}
	}
	if scs_120kHz_r16Present {
		ie.scs_120kHz_r16 = new(MIMO_ParametersPerBand_beamSwitchTiming_r16_scs_120kHz_r16)
		if err = ie.scs_120kHz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode scs_120kHz_r16", err)
		}
	}
	return nil
}
