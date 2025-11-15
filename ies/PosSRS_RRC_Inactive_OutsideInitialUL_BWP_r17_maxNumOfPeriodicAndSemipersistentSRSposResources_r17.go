package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17_Enum_n1  uper.Enumerated = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17_Enum_n2  uper.Enumerated = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17_Enum_n4  uper.Enumerated = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17_Enum_n8  uper.Enumerated = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17_Enum_n16 uper.Enumerated = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17_Enum_n32 uper.Enumerated = 5
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17_Enum_n64 uper.Enumerated = 6
)

type PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17", err)
	}
	return nil
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxNumOfPeriodicAndSemipersistentSRSposResources_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
