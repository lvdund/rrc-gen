package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReducedAggregatedBandwidth_Enum_mhz0   uper.Enumerated = 0
	ReducedAggregatedBandwidth_Enum_mhz10  uper.Enumerated = 1
	ReducedAggregatedBandwidth_Enum_mhz20  uper.Enumerated = 2
	ReducedAggregatedBandwidth_Enum_mhz30  uper.Enumerated = 3
	ReducedAggregatedBandwidth_Enum_mhz40  uper.Enumerated = 4
	ReducedAggregatedBandwidth_Enum_mhz50  uper.Enumerated = 5
	ReducedAggregatedBandwidth_Enum_mhz60  uper.Enumerated = 6
	ReducedAggregatedBandwidth_Enum_mhz80  uper.Enumerated = 7
	ReducedAggregatedBandwidth_Enum_mhz100 uper.Enumerated = 8
	ReducedAggregatedBandwidth_Enum_mhz200 uper.Enumerated = 9
	ReducedAggregatedBandwidth_Enum_mhz300 uper.Enumerated = 10
	ReducedAggregatedBandwidth_Enum_mhz400 uper.Enumerated = 11
)

type ReducedAggregatedBandwidth struct {
	Value uper.Enumerated `lb:0,ub:11,madatory`
}

func (ie *ReducedAggregatedBandwidth) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Encode ReducedAggregatedBandwidth", err)
	}
	return nil
}

func (ie *ReducedAggregatedBandwidth) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Decode ReducedAggregatedBandwidth", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
