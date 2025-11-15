package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms5    uper.Enumerated = 0
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms10   uper.Enumerated = 1
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms20   uper.Enumerated = 2
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms40   uper.Enumerated = 3
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms80   uper.Enumerated = 4
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_ms160  uper.Enumerated = 5
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_spare2 uper.Enumerated = 6
	ServingCellConfigCommon_ssb_periodicityServingCell_Enum_spare1 uper.Enumerated = 7
)

type ServingCellConfigCommon_ssb_periodicityServingCell struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *ServingCellConfigCommon_ssb_periodicityServingCell) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfigCommon_ssb_periodicityServingCell", err)
	}
	return nil
}

func (ie *ServingCellConfigCommon_ssb_periodicityServingCell) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfigCommon_ssb_periodicityServingCell", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
