package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n1  uper.Enumerated = 0
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n2  uper.Enumerated = 1
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n4  uper.Enumerated = 2
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n8  uper.Enumerated = 3
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n16 uper.Enumerated = 4
	EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod_Enum_n32 uper.Enumerated = 5
)

type EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod", err)
	}
	return nil
}

func (ie *EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode EUTRA_MBSFN_SubframeConfig_radioframeAllocationPeriod", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
