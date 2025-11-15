package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz50   uper.Enumerated = 0
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz100  uper.Enumerated = 1
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz200  uper.Enumerated = 2
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz400  uper.Enumerated = 3
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz800  uper.Enumerated = 4
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz1600 uper.Enumerated = 5
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz2000 uper.Enumerated = 6
)

type SupportedBandwidth_v1700_fr2_r17 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SupportedBandwidth_v1700_fr2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SupportedBandwidth_v1700_fr2_r17", err)
	}
	return nil
}

func (ie *SupportedBandwidth_v1700_fr2_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SupportedBandwidth_v1700_fr2_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
