package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_ConfigBroadcast_r17_mcs_Table_r17_Enum_qam256     uper.Enumerated = 0
	PDSCH_ConfigBroadcast_r17_mcs_Table_r17_Enum_qam64LowSE uper.Enumerated = 1
)

type PDSCH_ConfigBroadcast_r17_mcs_Table_r17 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PDSCH_ConfigBroadcast_r17_mcs_Table_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PDSCH_ConfigBroadcast_r17_mcs_Table_r17", err)
	}
	return nil
}

func (ie *PDSCH_ConfigBroadcast_r17_mcs_Table_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PDSCH_ConfigBroadcast_r17_mcs_Table_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
