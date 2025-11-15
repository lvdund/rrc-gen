package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IMS_Parameters struct {
	ims_ParametersCommon   *IMS_ParametersCommon   `optional`
	ims_ParametersFRX_Diff *IMS_ParametersFRX_Diff `optional`
}

func (ie *IMS_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ims_ParametersCommon != nil, ie.ims_ParametersFRX_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ims_ParametersCommon != nil {
		if err = ie.ims_ParametersCommon.Encode(w); err != nil {
			return utils.WrapError("Encode ims_ParametersCommon", err)
		}
	}
	if ie.ims_ParametersFRX_Diff != nil {
		if err = ie.ims_ParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode ims_ParametersFRX_Diff", err)
		}
	}
	return nil
}

func (ie *IMS_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var ims_ParametersCommonPresent bool
	if ims_ParametersCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ims_ParametersFRX_DiffPresent bool
	if ims_ParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ims_ParametersCommonPresent {
		ie.ims_ParametersCommon = new(IMS_ParametersCommon)
		if err = ie.ims_ParametersCommon.Decode(r); err != nil {
			return utils.WrapError("Decode ims_ParametersCommon", err)
		}
	}
	if ims_ParametersFRX_DiffPresent {
		ie.ims_ParametersFRX_Diff = new(IMS_ParametersFRX_Diff)
		if err = ie.ims_ParametersFRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode ims_ParametersFRX_Diff", err)
		}
	}
	return nil
}
