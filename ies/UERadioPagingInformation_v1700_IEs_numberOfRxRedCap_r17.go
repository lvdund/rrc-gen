package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UERadioPagingInformation_v1700_IEs_numberOfRxRedCap_r17_Enum_one uper.Enumerated = 0
	UERadioPagingInformation_v1700_IEs_numberOfRxRedCap_r17_Enum_two uper.Enumerated = 1
)

type UERadioPagingInformation_v1700_IEs_numberOfRxRedCap_r17 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *UERadioPagingInformation_v1700_IEs_numberOfRxRedCap_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode UERadioPagingInformation_v1700_IEs_numberOfRxRedCap_r17", err)
	}
	return nil
}

func (ie *UERadioPagingInformation_v1700_IEs_numberOfRxRedCap_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode UERadioPagingInformation_v1700_IEs_numberOfRxRedCap_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
