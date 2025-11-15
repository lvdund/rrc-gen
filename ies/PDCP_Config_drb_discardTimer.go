package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Config_drb_discardTimer_Enum_ms10     uper.Enumerated = 0
	PDCP_Config_drb_discardTimer_Enum_ms20     uper.Enumerated = 1
	PDCP_Config_drb_discardTimer_Enum_ms30     uper.Enumerated = 2
	PDCP_Config_drb_discardTimer_Enum_ms40     uper.Enumerated = 3
	PDCP_Config_drb_discardTimer_Enum_ms50     uper.Enumerated = 4
	PDCP_Config_drb_discardTimer_Enum_ms60     uper.Enumerated = 5
	PDCP_Config_drb_discardTimer_Enum_ms75     uper.Enumerated = 6
	PDCP_Config_drb_discardTimer_Enum_ms100    uper.Enumerated = 7
	PDCP_Config_drb_discardTimer_Enum_ms150    uper.Enumerated = 8
	PDCP_Config_drb_discardTimer_Enum_ms200    uper.Enumerated = 9
	PDCP_Config_drb_discardTimer_Enum_ms250    uper.Enumerated = 10
	PDCP_Config_drb_discardTimer_Enum_ms300    uper.Enumerated = 11
	PDCP_Config_drb_discardTimer_Enum_ms500    uper.Enumerated = 12
	PDCP_Config_drb_discardTimer_Enum_ms750    uper.Enumerated = 13
	PDCP_Config_drb_discardTimer_Enum_ms1500   uper.Enumerated = 14
	PDCP_Config_drb_discardTimer_Enum_infinity uper.Enumerated = 15
)

type PDCP_Config_drb_discardTimer struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *PDCP_Config_drb_discardTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PDCP_Config_drb_discardTimer", err)
	}
	return nil
}

func (ie *PDCP_Config_drb_discardTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PDCP_Config_drb_discardTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
