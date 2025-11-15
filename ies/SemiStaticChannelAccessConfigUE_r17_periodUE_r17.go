package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms1     uper.Enumerated = 0
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms2     uper.Enumerated = 1
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms2dot5 uper.Enumerated = 2
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms4     uper.Enumerated = 3
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms5     uper.Enumerated = 4
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_ms10    uper.Enumerated = 5
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_spare2  uper.Enumerated = 6
	SemiStaticChannelAccessConfigUE_r17_periodUE_r17_Enum_spare1  uper.Enumerated = 7
)

type SemiStaticChannelAccessConfigUE_r17_periodUE_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SemiStaticChannelAccessConfigUE_r17_periodUE_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SemiStaticChannelAccessConfigUE_r17_periodUE_r17", err)
	}
	return nil
}

func (ie *SemiStaticChannelAccessConfigUE_r17_periodUE_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SemiStaticChannelAccessConfigUE_r17_periodUE_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
