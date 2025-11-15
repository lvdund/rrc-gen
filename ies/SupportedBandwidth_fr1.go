package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedBandwidth_fr1_Enum_mhz5   uper.Enumerated = 0
	SupportedBandwidth_fr1_Enum_mhz10  uper.Enumerated = 1
	SupportedBandwidth_fr1_Enum_mhz15  uper.Enumerated = 2
	SupportedBandwidth_fr1_Enum_mhz20  uper.Enumerated = 3
	SupportedBandwidth_fr1_Enum_mhz25  uper.Enumerated = 4
	SupportedBandwidth_fr1_Enum_mhz30  uper.Enumerated = 5
	SupportedBandwidth_fr1_Enum_mhz40  uper.Enumerated = 6
	SupportedBandwidth_fr1_Enum_mhz50  uper.Enumerated = 7
	SupportedBandwidth_fr1_Enum_mhz60  uper.Enumerated = 8
	SupportedBandwidth_fr1_Enum_mhz80  uper.Enumerated = 9
	SupportedBandwidth_fr1_Enum_mhz100 uper.Enumerated = 10
)

type SupportedBandwidth_fr1 struct {
	Value uper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *SupportedBandwidth_fr1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode SupportedBandwidth_fr1", err)
	}
	return nil
}

func (ie *SupportedBandwidth_fr1) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode SupportedBandwidth_fr1", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
