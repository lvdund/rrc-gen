package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GeneralParametersMRDC_v1610 struct {
	f1c_OverEUTRA_r16 *GeneralParametersMRDC_v1610_f1c_OverEUTRA_r16 `optional`
}

func (ie *GeneralParametersMRDC_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.f1c_OverEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.f1c_OverEUTRA_r16 != nil {
		if err = ie.f1c_OverEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode f1c_OverEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *GeneralParametersMRDC_v1610) Decode(r *uper.UperReader) error {
	var err error
	var f1c_OverEUTRA_r16Present bool
	if f1c_OverEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if f1c_OverEUTRA_r16Present {
		ie.f1c_OverEUTRA_r16 = new(GeneralParametersMRDC_v1610_f1c_OverEUTRA_r16)
		if err = ie.f1c_OverEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode f1c_OverEUTRA_r16", err)
		}
	}
	return nil
}
