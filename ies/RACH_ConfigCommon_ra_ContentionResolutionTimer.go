package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_ra_ContentionResolutionTimer_Enum_sf8  uper.Enumerated = 0
	RACH_ConfigCommon_ra_ContentionResolutionTimer_Enum_sf16 uper.Enumerated = 1
	RACH_ConfigCommon_ra_ContentionResolutionTimer_Enum_sf24 uper.Enumerated = 2
	RACH_ConfigCommon_ra_ContentionResolutionTimer_Enum_sf32 uper.Enumerated = 3
	RACH_ConfigCommon_ra_ContentionResolutionTimer_Enum_sf40 uper.Enumerated = 4
	RACH_ConfigCommon_ra_ContentionResolutionTimer_Enum_sf48 uper.Enumerated = 5
	RACH_ConfigCommon_ra_ContentionResolutionTimer_Enum_sf56 uper.Enumerated = 6
	RACH_ConfigCommon_ra_ContentionResolutionTimer_Enum_sf64 uper.Enumerated = 7
)

type RACH_ConfigCommon_ra_ContentionResolutionTimer struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RACH_ConfigCommon_ra_ContentionResolutionTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommon_ra_ContentionResolutionTimer", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_ra_ContentionResolutionTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommon_ra_ContentionResolutionTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
