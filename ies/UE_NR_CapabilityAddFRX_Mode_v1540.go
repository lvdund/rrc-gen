package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_CapabilityAddFRX_Mode_v1540 struct {
	ims_ParametersFRX_Diff *IMS_ParametersFRX_Diff `optional`
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ims_ParametersFRX_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ims_ParametersFRX_Diff != nil {
		if err = ie.ims_ParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode ims_ParametersFRX_Diff", err)
		}
	}
	return nil
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1540) Decode(r *uper.UperReader) error {
	var err error
	var ims_ParametersFRX_DiffPresent bool
	if ims_ParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ims_ParametersFRX_DiffPresent {
		ie.ims_ParametersFRX_Diff = new(IMS_ParametersFRX_Diff)
		if err = ie.ims_ParametersFRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode ims_ParametersFRX_Diff", err)
		}
	}
	return nil
}
