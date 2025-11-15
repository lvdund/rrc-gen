package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedCSI_RS_Resource_maxNumberTxPortsPerResource_Enum_p2  uper.Enumerated = 0
	SupportedCSI_RS_Resource_maxNumberTxPortsPerResource_Enum_p4  uper.Enumerated = 1
	SupportedCSI_RS_Resource_maxNumberTxPortsPerResource_Enum_p8  uper.Enumerated = 2
	SupportedCSI_RS_Resource_maxNumberTxPortsPerResource_Enum_p12 uper.Enumerated = 3
	SupportedCSI_RS_Resource_maxNumberTxPortsPerResource_Enum_p16 uper.Enumerated = 4
	SupportedCSI_RS_Resource_maxNumberTxPortsPerResource_Enum_p24 uper.Enumerated = 5
	SupportedCSI_RS_Resource_maxNumberTxPortsPerResource_Enum_p32 uper.Enumerated = 6
)

type SupportedCSI_RS_Resource_maxNumberTxPortsPerResource struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SupportedCSI_RS_Resource_maxNumberTxPortsPerResource) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SupportedCSI_RS_Resource_maxNumberTxPortsPerResource", err)
	}
	return nil
}

func (ie *SupportedCSI_RS_Resource_maxNumberTxPortsPerResource) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SupportedCSI_RS_Resource_maxNumberTxPortsPerResource", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
