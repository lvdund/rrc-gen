package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf20   uper.Enumerated = 0
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf40   uper.Enumerated = 1
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf64   uper.Enumerated = 2
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf128  uper.Enumerated = 3
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf512  uper.Enumerated = 4
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf1024 uper.Enumerated = 5
	BSR_Config_logicalChannelSR_DelayTimer_Enum_sf2560 uper.Enumerated = 6
	BSR_Config_logicalChannelSR_DelayTimer_Enum_spare1 uper.Enumerated = 7
)

type BSR_Config_logicalChannelSR_DelayTimer struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *BSR_Config_logicalChannelSR_DelayTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode BSR_Config_logicalChannelSR_DelayTimer", err)
	}
	return nil
}

func (ie *BSR_Config_logicalChannelSR_DelayTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode BSR_Config_logicalChannelSR_DelayTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
