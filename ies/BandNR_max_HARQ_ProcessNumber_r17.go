package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_max_HARQ_ProcessNumber_r17_Enum_u16d32 uper.Enumerated = 0
	BandNR_max_HARQ_ProcessNumber_r17_Enum_u32d16 uper.Enumerated = 1
	BandNR_max_HARQ_ProcessNumber_r17_Enum_u32d32 uper.Enumerated = 2
)

type BandNR_max_HARQ_ProcessNumber_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *BandNR_max_HARQ_ProcessNumber_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode BandNR_max_HARQ_ProcessNumber_r17", err)
	}
	return nil
}

func (ie *BandNR_max_HARQ_ProcessNumber_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode BandNR_max_HARQ_ProcessNumber_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
