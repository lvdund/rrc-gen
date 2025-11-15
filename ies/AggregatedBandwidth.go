package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	AggregatedBandwidth_Enum_mhz50  uper.Enumerated = 0
	AggregatedBandwidth_Enum_mhz100 uper.Enumerated = 1
	AggregatedBandwidth_Enum_mhz150 uper.Enumerated = 2
	AggregatedBandwidth_Enum_mhz200 uper.Enumerated = 3
	AggregatedBandwidth_Enum_mhz250 uper.Enumerated = 4
	AggregatedBandwidth_Enum_mhz300 uper.Enumerated = 5
	AggregatedBandwidth_Enum_mhz350 uper.Enumerated = 6
	AggregatedBandwidth_Enum_mhz400 uper.Enumerated = 7
	AggregatedBandwidth_Enum_mhz450 uper.Enumerated = 8
	AggregatedBandwidth_Enum_mhz500 uper.Enumerated = 9
	AggregatedBandwidth_Enum_mhz550 uper.Enumerated = 10
	AggregatedBandwidth_Enum_mhz600 uper.Enumerated = 11
	AggregatedBandwidth_Enum_mhz650 uper.Enumerated = 12
	AggregatedBandwidth_Enum_mhz700 uper.Enumerated = 13
	AggregatedBandwidth_Enum_mhz750 uper.Enumerated = 14
	AggregatedBandwidth_Enum_mhz800 uper.Enumerated = 15
)

type AggregatedBandwidth struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *AggregatedBandwidth) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode AggregatedBandwidth", err)
	}
	return nil
}

func (ie *AggregatedBandwidth) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode AggregatedBandwidth", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
