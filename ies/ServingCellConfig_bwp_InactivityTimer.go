package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfig_bwp_InactivityTimer_Enum_ms2     uper.Enumerated = 0
	ServingCellConfig_bwp_InactivityTimer_Enum_ms3     uper.Enumerated = 1
	ServingCellConfig_bwp_InactivityTimer_Enum_ms4     uper.Enumerated = 2
	ServingCellConfig_bwp_InactivityTimer_Enum_ms5     uper.Enumerated = 3
	ServingCellConfig_bwp_InactivityTimer_Enum_ms6     uper.Enumerated = 4
	ServingCellConfig_bwp_InactivityTimer_Enum_ms8     uper.Enumerated = 5
	ServingCellConfig_bwp_InactivityTimer_Enum_ms10    uper.Enumerated = 6
	ServingCellConfig_bwp_InactivityTimer_Enum_ms20    uper.Enumerated = 7
	ServingCellConfig_bwp_InactivityTimer_Enum_ms30    uper.Enumerated = 8
	ServingCellConfig_bwp_InactivityTimer_Enum_ms40    uper.Enumerated = 9
	ServingCellConfig_bwp_InactivityTimer_Enum_ms50    uper.Enumerated = 10
	ServingCellConfig_bwp_InactivityTimer_Enum_ms60    uper.Enumerated = 11
	ServingCellConfig_bwp_InactivityTimer_Enum_ms80    uper.Enumerated = 12
	ServingCellConfig_bwp_InactivityTimer_Enum_ms100   uper.Enumerated = 13
	ServingCellConfig_bwp_InactivityTimer_Enum_ms200   uper.Enumerated = 14
	ServingCellConfig_bwp_InactivityTimer_Enum_ms300   uper.Enumerated = 15
	ServingCellConfig_bwp_InactivityTimer_Enum_ms500   uper.Enumerated = 16
	ServingCellConfig_bwp_InactivityTimer_Enum_ms750   uper.Enumerated = 17
	ServingCellConfig_bwp_InactivityTimer_Enum_ms1280  uper.Enumerated = 18
	ServingCellConfig_bwp_InactivityTimer_Enum_ms1920  uper.Enumerated = 19
	ServingCellConfig_bwp_InactivityTimer_Enum_ms2560  uper.Enumerated = 20
	ServingCellConfig_bwp_InactivityTimer_Enum_spare10 uper.Enumerated = 21
	ServingCellConfig_bwp_InactivityTimer_Enum_spare9  uper.Enumerated = 22
	ServingCellConfig_bwp_InactivityTimer_Enum_spare8  uper.Enumerated = 23
	ServingCellConfig_bwp_InactivityTimer_Enum_spare7  uper.Enumerated = 24
	ServingCellConfig_bwp_InactivityTimer_Enum_spare6  uper.Enumerated = 25
	ServingCellConfig_bwp_InactivityTimer_Enum_spare5  uper.Enumerated = 26
	ServingCellConfig_bwp_InactivityTimer_Enum_spare4  uper.Enumerated = 27
	ServingCellConfig_bwp_InactivityTimer_Enum_spare3  uper.Enumerated = 28
	ServingCellConfig_bwp_InactivityTimer_Enum_spare2  uper.Enumerated = 29
	ServingCellConfig_bwp_InactivityTimer_Enum_spare1  uper.Enumerated = 30
)

type ServingCellConfig_bwp_InactivityTimer struct {
	Value uper.Enumerated `lb:0,ub:30,madatory`
}

func (ie *ServingCellConfig_bwp_InactivityTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfig_bwp_InactivityTimer", err)
	}
	return nil
}

func (ie *ServingCellConfig_bwp_InactivityTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 30}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfig_bwp_InactivityTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
