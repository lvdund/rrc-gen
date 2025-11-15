package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms0p5   uper.Enumerated = 0
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms0p625 uper.Enumerated = 1
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms1     uper.Enumerated = 2
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms1p25  uper.Enumerated = 3
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms2     uper.Enumerated = 4
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms2p5   uper.Enumerated = 5
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms5     uper.Enumerated = 6
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms10    uper.Enumerated = 7
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms20    uper.Enumerated = 8
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms40    uper.Enumerated = 9
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms80    uper.Enumerated = 10
	IAB_ResourceConfig_r17_periodicitySlotList_r17_Enum_ms160   uper.Enumerated = 11
)

type IAB_ResourceConfig_r17_periodicitySlotList_r17 struct {
	Value uper.Enumerated `lb:0,ub:11,madatory`
}

func (ie *IAB_ResourceConfig_r17_periodicitySlotList_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Encode IAB_ResourceConfig_r17_periodicitySlotList_r17", err)
	}
	return nil
}

func (ie *IAB_ResourceConfig_r17_periodicitySlotList_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Decode IAB_ResourceConfig_r17_periodicitySlotList_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
