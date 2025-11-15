package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17_Enum_us100 uper.Enumerated = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17_Enum_us140 uper.Enumerated = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17_Enum_us200 uper.Enumerated = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17_Enum_us300 uper.Enumerated = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17_Enum_us500 uper.Enumerated = 4
)

type PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17 struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17", err)
	}
	return nil
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_switchingTimeSRS_TX_OtherTX_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
