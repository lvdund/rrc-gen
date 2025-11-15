package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DownlinkPreemption_timeFrequencySet_Enum_set0 uper.Enumerated = 0
	DownlinkPreemption_timeFrequencySet_Enum_set1 uper.Enumerated = 1
)

type DownlinkPreemption_timeFrequencySet struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *DownlinkPreemption_timeFrequencySet) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode DownlinkPreemption_timeFrequencySet", err)
	}
	return nil
}

func (ie *DownlinkPreemption_timeFrequencySet) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode DownlinkPreemption_timeFrequencySet", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
