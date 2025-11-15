package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BSR_Config_periodicBSR_Timer_Enum_sf1      uper.Enumerated = 0
	BSR_Config_periodicBSR_Timer_Enum_sf5      uper.Enumerated = 1
	BSR_Config_periodicBSR_Timer_Enum_sf10     uper.Enumerated = 2
	BSR_Config_periodicBSR_Timer_Enum_sf16     uper.Enumerated = 3
	BSR_Config_periodicBSR_Timer_Enum_sf20     uper.Enumerated = 4
	BSR_Config_periodicBSR_Timer_Enum_sf32     uper.Enumerated = 5
	BSR_Config_periodicBSR_Timer_Enum_sf40     uper.Enumerated = 6
	BSR_Config_periodicBSR_Timer_Enum_sf64     uper.Enumerated = 7
	BSR_Config_periodicBSR_Timer_Enum_sf80     uper.Enumerated = 8
	BSR_Config_periodicBSR_Timer_Enum_sf128    uper.Enumerated = 9
	BSR_Config_periodicBSR_Timer_Enum_sf160    uper.Enumerated = 10
	BSR_Config_periodicBSR_Timer_Enum_sf320    uper.Enumerated = 11
	BSR_Config_periodicBSR_Timer_Enum_sf640    uper.Enumerated = 12
	BSR_Config_periodicBSR_Timer_Enum_sf1280   uper.Enumerated = 13
	BSR_Config_periodicBSR_Timer_Enum_sf2560   uper.Enumerated = 14
	BSR_Config_periodicBSR_Timer_Enum_infinity uper.Enumerated = 15
)

type BSR_Config_periodicBSR_Timer struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *BSR_Config_periodicBSR_Timer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode BSR_Config_periodicBSR_Timer", err)
	}
	return nil
}

func (ie *BSR_Config_periodicBSR_Timer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode BSR_Config_periodicBSR_Timer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
