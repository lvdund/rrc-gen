package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IMS_Parameters_v1700 struct {
	ims_ParametersFR2_2_r17 *IMS_ParametersFR2_2_r17 `optional`
}

func (ie *IMS_Parameters_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ims_ParametersFR2_2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ims_ParametersFR2_2_r17 != nil {
		if err = ie.ims_ParametersFR2_2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ims_ParametersFR2_2_r17", err)
		}
	}
	return nil
}

func (ie *IMS_Parameters_v1700) Decode(r *uper.UperReader) error {
	var err error
	var ims_ParametersFR2_2_r17Present bool
	if ims_ParametersFR2_2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ims_ParametersFR2_2_r17Present {
		ie.ims_ParametersFR2_2_r17 = new(IMS_ParametersFR2_2_r17)
		if err = ie.ims_ParametersFR2_2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ims_ParametersFR2_2_r17", err)
		}
	}
	return nil
}
