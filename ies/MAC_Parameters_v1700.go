package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_Parameters_v1700 struct {
	mac_ParametersFR2_2_r17 *MAC_ParametersFR2_2_r17 `optional`
}

func (ie *MAC_Parameters_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mac_ParametersFR2_2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mac_ParametersFR2_2_r17 != nil {
		if err = ie.mac_ParametersFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mac_ParametersFR2_2_r17", err)
		}
	}
	return nil
}

func (ie *MAC_Parameters_v1700) Decode(r *uper.UperReader) error {
	var err error
	var mac_ParametersFR2_2_r17Present bool
	if mac_ParametersFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if mac_ParametersFR2_2_r17Present {
		ie.mac_ParametersFR2_2_r17 = new(MAC_ParametersFR2_2_r17)
		if err = ie.mac_ParametersFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mac_ParametersFR2_2_r17", err)
		}
	}
	return nil
}
