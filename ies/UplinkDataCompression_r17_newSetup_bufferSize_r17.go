package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UplinkDataCompression_r17_newSetup_bufferSize_r17_Enum_kbyte2 uper.Enumerated = 0
	UplinkDataCompression_r17_newSetup_bufferSize_r17_Enum_kbyte4 uper.Enumerated = 1
	UplinkDataCompression_r17_newSetup_bufferSize_r17_Enum_kbyte8 uper.Enumerated = 2
	UplinkDataCompression_r17_newSetup_bufferSize_r17_Enum_spare1 uper.Enumerated = 3
)

type UplinkDataCompression_r17_newSetup_bufferSize_r17 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *UplinkDataCompression_r17_newSetup_bufferSize_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode UplinkDataCompression_r17_newSetup_bufferSize_r17", err)
	}
	return nil
}

func (ie *UplinkDataCompression_r17_newSetup_bufferSize_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode UplinkDataCompression_r17_newSetup_bufferSize_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
