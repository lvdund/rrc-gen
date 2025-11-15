package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedBandwidth_fr2_Enum_mhz50  uper.Enumerated = 0
	SupportedBandwidth_fr2_Enum_mhz100 uper.Enumerated = 1
	SupportedBandwidth_fr2_Enum_mhz200 uper.Enumerated = 2
	SupportedBandwidth_fr2_Enum_mhz400 uper.Enumerated = 3
)

type SupportedBandwidth_fr2 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SupportedBandwidth_fr2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SupportedBandwidth_fr2", err)
	}
	return nil
}

func (ie *SupportedBandwidth_fr2) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SupportedBandwidth_fr2", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
