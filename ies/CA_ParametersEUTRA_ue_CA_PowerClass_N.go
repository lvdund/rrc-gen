package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CA_ParametersEUTRA_ue_CA_PowerClass_N_Enum_class2 uper.Enumerated = 0
)

type CA_ParametersEUTRA_ue_CA_PowerClass_N struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *CA_ParametersEUTRA_ue_CA_PowerClass_N) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode CA_ParametersEUTRA_ue_CA_PowerClass_N", err)
	}
	return nil
}

func (ie *CA_ParametersEUTRA_ue_CA_PowerClass_N) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode CA_ParametersEUTRA_ue_CA_PowerClass_N", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
