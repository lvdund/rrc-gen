package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n10  uper.Enumerated = 0
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n12  uper.Enumerated = 1
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n15  uper.Enumerated = 2
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n20  uper.Enumerated = 3
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n25  uper.Enumerated = 4
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n50  uper.Enumerated = 5
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n75  uper.Enumerated = 6
	SL_ResourcePool_r16_sl_SubchannelSize_r16_Enum_n100 uper.Enumerated = 7
)

type SL_ResourcePool_r16_sl_SubchannelSize_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_ResourcePool_r16_sl_SubchannelSize_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_ResourcePool_r16_sl_SubchannelSize_r16", err)
	}
	return nil
}

func (ie *SL_ResourcePool_r16_sl_SubchannelSize_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_ResourcePool_r16_sl_SubchannelSize_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
