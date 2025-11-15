package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PowSav_Parameters_r16 struct {
	powSav_ParametersCommon_r16   *PowSav_ParametersCommon_r16   `optional`
	powSav_ParametersFRX_Diff_r16 *PowSav_ParametersFRX_Diff_r16 `optional`
}

func (ie *PowSav_Parameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.powSav_ParametersCommon_r16 != nil, ie.powSav_ParametersFRX_Diff_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.powSav_ParametersCommon_r16 != nil {
		if err = ie.powSav_ParametersCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode powSav_ParametersCommon_r16", err)
		}
	}
	if ie.powSav_ParametersFRX_Diff_r16 != nil {
		if err = ie.powSav_ParametersFRX_Diff_r16.Encode(w); err != nil {
			return utils.WrapError("Encode powSav_ParametersFRX_Diff_r16", err)
		}
	}
	return nil
}

func (ie *PowSav_Parameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var powSav_ParametersCommon_r16Present bool
	if powSav_ParametersCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var powSav_ParametersFRX_Diff_r16Present bool
	if powSav_ParametersFRX_Diff_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if powSav_ParametersCommon_r16Present {
		ie.powSav_ParametersCommon_r16 = new(PowSav_ParametersCommon_r16)
		if err = ie.powSav_ParametersCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode powSav_ParametersCommon_r16", err)
		}
	}
	if powSav_ParametersFRX_Diff_r16Present {
		ie.powSav_ParametersFRX_Diff_r16 = new(PowSav_ParametersFRX_Diff_r16)
		if err = ie.powSav_ParametersFRX_Diff_r16.Decode(r); err != nil {
			return utils.WrapError("Decode powSav_ParametersFRX_Diff_r16", err)
		}
	}
	return nil
}
