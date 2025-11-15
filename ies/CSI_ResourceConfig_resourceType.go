package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ResourceConfig_resourceType_Enum_aperiodic      uper.Enumerated = 0
	CSI_ResourceConfig_resourceType_Enum_semiPersistent uper.Enumerated = 1
	CSI_ResourceConfig_resourceType_Enum_periodic       uper.Enumerated = 2
)

type CSI_ResourceConfig_resourceType struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CSI_ResourceConfig_resourceType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CSI_ResourceConfig_resourceType", err)
	}
	return nil
}

func (ie *CSI_ResourceConfig_resourceType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CSI_ResourceConfig_resourceType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
