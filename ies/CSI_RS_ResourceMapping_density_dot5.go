package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_ResourceMapping_density_dot5_Enum_evenPRBs uper.Enumerated = 0
	CSI_RS_ResourceMapping_density_dot5_Enum_oddPRBs  uper.Enumerated = 1
)

type CSI_RS_ResourceMapping_density_dot5 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CSI_RS_ResourceMapping_density_dot5) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CSI_RS_ResourceMapping_density_dot5", err)
	}
	return nil
}

func (ie *CSI_RS_ResourceMapping_density_dot5) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CSI_RS_ResourceMapping_density_dot5", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
