package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf20   uper.Enumerated = 0
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf40   uper.Enumerated = 1
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf64   uper.Enumerated = 2
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf128  uper.Enumerated = 3
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf512  uper.Enumerated = 4
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf1024 uper.Enumerated = 5
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_sf2560 uper.Enumerated = 6
	SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17_Enum_spare1 uper.Enumerated = 7
)

type SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17", err)
	}
	return nil
}

func (ie *SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SDT_ConfigCommonSIB_r17_sdt_LogicalChannelSR_DelayTimer_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
