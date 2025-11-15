package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PowSav_Parameters_v1700 struct {
	powSav_ParametersFR2_2_r17 *PowSav_ParametersFR2_2_r17 `optional`
}

func (ie *PowSav_Parameters_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.powSav_ParametersFR2_2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.powSav_ParametersFR2_2_r17 != nil {
		if err = ie.powSav_ParametersFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode powSav_ParametersFR2_2_r17", err)
		}
	}
	return nil
}

func (ie *PowSav_Parameters_v1700) Decode(r *uper.UperReader) error {
	var err error
	var powSav_ParametersFR2_2_r17Present bool
	if powSav_ParametersFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if powSav_ParametersFR2_2_r17Present {
		ie.powSav_ParametersFR2_2_r17 = new(PowSav_ParametersFR2_2_r17)
		if err = ie.powSav_ParametersFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode powSav_ParametersFR2_2_r17", err)
		}
	}
	return nil
}
