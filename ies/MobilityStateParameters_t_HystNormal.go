package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MobilityStateParameters_t_HystNormal_Enum_s30    uper.Enumerated = 0
	MobilityStateParameters_t_HystNormal_Enum_s60    uper.Enumerated = 1
	MobilityStateParameters_t_HystNormal_Enum_s120   uper.Enumerated = 2
	MobilityStateParameters_t_HystNormal_Enum_s180   uper.Enumerated = 3
	MobilityStateParameters_t_HystNormal_Enum_s240   uper.Enumerated = 4
	MobilityStateParameters_t_HystNormal_Enum_spare3 uper.Enumerated = 5
	MobilityStateParameters_t_HystNormal_Enum_spare2 uper.Enumerated = 6
	MobilityStateParameters_t_HystNormal_Enum_spare1 uper.Enumerated = 7
)

type MobilityStateParameters_t_HystNormal struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MobilityStateParameters_t_HystNormal) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MobilityStateParameters_t_HystNormal", err)
	}
	return nil
}

func (ie *MobilityStateParameters_t_HystNormal) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MobilityStateParameters_t_HystNormal", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
