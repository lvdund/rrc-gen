package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw5   uper.Enumerated = 0
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw10  uper.Enumerated = 1
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw15  uper.Enumerated = 2
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw20  uper.Enumerated = 3
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw25  uper.Enumerated = 4
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw30  uper.Enumerated = 5
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw35  uper.Enumerated = 6
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw40  uper.Enumerated = 7
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw45  uper.Enumerated = 8
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw50  uper.Enumerated = 9
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw60  uper.Enumerated = 10
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw70  uper.Enumerated = 11
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw80  uper.Enumerated = 12
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw90  uper.Enumerated = 13
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17_Enum_bw100 uper.Enumerated = 14
)

type PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17 struct {
	Value uper.Enumerated `lb:0,ub:14,madatory`
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 14}, false); err != nil {
		return utils.WrapError("Encode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17", err)
	}
	return nil
}

func (ie *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 14}, false); err != nil {
		return utils.WrapError("Decode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17_maxSRSposBandwidthForEachSCS_withinCC_FR1_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
