package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_ResourceMapping_cdm_Type_Enum_noCDM        uper.Enumerated = 0
	CSI_RS_ResourceMapping_cdm_Type_Enum_fd_CDM2      uper.Enumerated = 1
	CSI_RS_ResourceMapping_cdm_Type_Enum_cdm4_FD2_TD2 uper.Enumerated = 2
	CSI_RS_ResourceMapping_cdm_Type_Enum_cdm8_FD2_TD4 uper.Enumerated = 3
)

type CSI_RS_ResourceMapping_cdm_Type struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *CSI_RS_ResourceMapping_cdm_Type) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode CSI_RS_ResourceMapping_cdm_Type", err)
	}
	return nil
}

func (ie *CSI_RS_ResourceMapping_cdm_Type) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode CSI_RS_ResourceMapping_cdm_Type", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
