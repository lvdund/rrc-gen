package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UE_CapabilityRequestFilterCommon_mrdc_Request_omitEN_DC_Enum_true uper.Enumerated = 0
)

type UE_CapabilityRequestFilterCommon_mrdc_Request_omitEN_DC struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *UE_CapabilityRequestFilterCommon_mrdc_Request_omitEN_DC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode UE_CapabilityRequestFilterCommon_mrdc_Request_omitEN_DC", err)
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterCommon_mrdc_Request_omitEN_DC) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode UE_CapabilityRequestFilterCommon_mrdc_Request_omitEN_DC", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
