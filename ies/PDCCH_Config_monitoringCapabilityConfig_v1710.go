package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCCH_Config_monitoringCapabilityConfig_v1710_Enum_r17monitoringcapability uper.Enumerated = 0
)

type PDCCH_Config_monitoringCapabilityConfig_v1710 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PDCCH_Config_monitoringCapabilityConfig_v1710) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PDCCH_Config_monitoringCapabilityConfig_v1710", err)
	}
	return nil
}

func (ie *PDCCH_Config_monitoringCapabilityConfig_v1710) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PDCCH_Config_monitoringCapabilityConfig_v1710", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
