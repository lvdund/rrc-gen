package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms1    uper.Enumerated = 0
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms10   uper.Enumerated = 1
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms40   uper.Enumerated = 2
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms160  uper.Enumerated = 3
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms500  uper.Enumerated = 4
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms1000 uper.Enumerated = 5
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms1250 uper.Enumerated = 6
	MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17_Enum_ms2750 uper.Enumerated = 7
)

type MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17", err)
	}
	return nil
}

func (ie *MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MRB_PDCP_ConfigBroadcast_r17_t_Reordering_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
