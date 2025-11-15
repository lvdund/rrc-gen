package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17_Enum_one   uper.Enumerated = 0
	RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17_Enum_two   uper.Enumerated = 1
	RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17_Enum_four  uper.Enumerated = 2
	RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17_Enum_eight uper.Enumerated = 3
)

type RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17", err)
	}
	return nil
}

func (ie *RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RA_InformationCommon_r16_msgA_RO_FDMCFRA_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
