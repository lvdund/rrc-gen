package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfig_sCellDeactivationTimer_Enum_ms20   uper.Enumerated = 0
	ServingCellConfig_sCellDeactivationTimer_Enum_ms40   uper.Enumerated = 1
	ServingCellConfig_sCellDeactivationTimer_Enum_ms80   uper.Enumerated = 2
	ServingCellConfig_sCellDeactivationTimer_Enum_ms160  uper.Enumerated = 3
	ServingCellConfig_sCellDeactivationTimer_Enum_ms200  uper.Enumerated = 4
	ServingCellConfig_sCellDeactivationTimer_Enum_ms240  uper.Enumerated = 5
	ServingCellConfig_sCellDeactivationTimer_Enum_ms320  uper.Enumerated = 6
	ServingCellConfig_sCellDeactivationTimer_Enum_ms400  uper.Enumerated = 7
	ServingCellConfig_sCellDeactivationTimer_Enum_ms480  uper.Enumerated = 8
	ServingCellConfig_sCellDeactivationTimer_Enum_ms520  uper.Enumerated = 9
	ServingCellConfig_sCellDeactivationTimer_Enum_ms640  uper.Enumerated = 10
	ServingCellConfig_sCellDeactivationTimer_Enum_ms720  uper.Enumerated = 11
	ServingCellConfig_sCellDeactivationTimer_Enum_ms840  uper.Enumerated = 12
	ServingCellConfig_sCellDeactivationTimer_Enum_ms1280 uper.Enumerated = 13
	ServingCellConfig_sCellDeactivationTimer_Enum_spare2 uper.Enumerated = 14
	ServingCellConfig_sCellDeactivationTimer_Enum_spare1 uper.Enumerated = 15
)

type ServingCellConfig_sCellDeactivationTimer struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *ServingCellConfig_sCellDeactivationTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfig_sCellDeactivationTimer", err)
	}
	return nil
}

func (ie *ServingCellConfig_sCellDeactivationTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfig_sCellDeactivationTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
