package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_ResourceMapping_nrofPorts_Enum_p1  uper.Enumerated = 0
	CSI_RS_ResourceMapping_nrofPorts_Enum_p2  uper.Enumerated = 1
	CSI_RS_ResourceMapping_nrofPorts_Enum_p4  uper.Enumerated = 2
	CSI_RS_ResourceMapping_nrofPorts_Enum_p8  uper.Enumerated = 3
	CSI_RS_ResourceMapping_nrofPorts_Enum_p12 uper.Enumerated = 4
	CSI_RS_ResourceMapping_nrofPorts_Enum_p16 uper.Enumerated = 5
	CSI_RS_ResourceMapping_nrofPorts_Enum_p24 uper.Enumerated = 6
	CSI_RS_ResourceMapping_nrofPorts_Enum_p32 uper.Enumerated = 7
)

type CSI_RS_ResourceMapping_nrofPorts struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CSI_RS_ResourceMapping_nrofPorts) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CSI_RS_ResourceMapping_nrofPorts", err)
	}
	return nil
}

func (ie *CSI_RS_ResourceMapping_nrofPorts) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CSI_RS_ResourceMapping_nrofPorts", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
