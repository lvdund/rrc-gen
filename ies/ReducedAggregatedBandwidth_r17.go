package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReducedAggregatedBandwidth_r17_Enum_mhz0    uper.Enumerated = 0
	ReducedAggregatedBandwidth_r17_Enum_mhz100  uper.Enumerated = 1
	ReducedAggregatedBandwidth_r17_Enum_mhz200  uper.Enumerated = 2
	ReducedAggregatedBandwidth_r17_Enum_mhz400  uper.Enumerated = 3
	ReducedAggregatedBandwidth_r17_Enum_mhz800  uper.Enumerated = 4
	ReducedAggregatedBandwidth_r17_Enum_mhz1200 uper.Enumerated = 5
	ReducedAggregatedBandwidth_r17_Enum_mhz1600 uper.Enumerated = 6
	ReducedAggregatedBandwidth_r17_Enum_mhz2000 uper.Enumerated = 7
)

type ReducedAggregatedBandwidth_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *ReducedAggregatedBandwidth_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode ReducedAggregatedBandwidth_r17", err)
	}
	return nil
}

func (ie *ReducedAggregatedBandwidth_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode ReducedAggregatedBandwidth_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
