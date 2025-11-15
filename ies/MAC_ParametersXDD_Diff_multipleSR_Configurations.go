package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MAC_ParametersXDD_Diff_multipleSR_Configurations_Enum_supported uper.Enumerated = 0
)

type MAC_ParametersXDD_Diff_multipleSR_Configurations struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MAC_ParametersXDD_Diff_multipleSR_Configurations) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MAC_ParametersXDD_Diff_multipleSR_Configurations", err)
	}
	return nil
}

func (ie *MAC_ParametersXDD_Diff_multipleSR_Configurations) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MAC_ParametersXDD_Diff_multipleSR_Configurations", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
