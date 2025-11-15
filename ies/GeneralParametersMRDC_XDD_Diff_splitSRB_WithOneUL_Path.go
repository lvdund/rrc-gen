package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	GeneralParametersMRDC_XDD_Diff_splitSRB_WithOneUL_Path_Enum_supported uper.Enumerated = 0
)

type GeneralParametersMRDC_XDD_Diff_splitSRB_WithOneUL_Path struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *GeneralParametersMRDC_XDD_Diff_splitSRB_WithOneUL_Path) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode GeneralParametersMRDC_XDD_Diff_splitSRB_WithOneUL_Path", err)
	}
	return nil
}

func (ie *GeneralParametersMRDC_XDD_Diff_splitSRB_WithOneUL_Path) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode GeneralParametersMRDC_XDD_Diff_splitSRB_WithOneUL_Path", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
