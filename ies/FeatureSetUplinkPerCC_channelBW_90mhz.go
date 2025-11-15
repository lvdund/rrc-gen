package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplinkPerCC_channelBW_90mhz_Enum_supported uper.Enumerated = 0
)

type FeatureSetUplinkPerCC_channelBW_90mhz struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FeatureSetUplinkPerCC_channelBW_90mhz) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplinkPerCC_channelBW_90mhz", err)
	}
	return nil
}

func (ie *FeatureSetUplinkPerCC_channelBW_90mhz) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplinkPerCC_channelBW_90mhz", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
