package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_xOverheadMulticast_r17_Enum_xOh6  uper.Enumerated = 0
	PDSCH_Config_xOverheadMulticast_r17_Enum_xOh12 uper.Enumerated = 1
	PDSCH_Config_xOverheadMulticast_r17_Enum_xOh18 uper.Enumerated = 2
)

type PDSCH_Config_xOverheadMulticast_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PDSCH_Config_xOverheadMulticast_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PDSCH_Config_xOverheadMulticast_r17", err)
	}
	return nil
}

func (ie *PDSCH_Config_xOverheadMulticast_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PDSCH_Config_xOverheadMulticast_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
