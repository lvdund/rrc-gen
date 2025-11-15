package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_aperiodicCSI_RS_AdditionalBandwidth_r17_Enum_addBW_Set1 uper.Enumerated = 0
	BandNR_aperiodicCSI_RS_AdditionalBandwidth_r17_Enum_addBW_Set2 uper.Enumerated = 1
)

type BandNR_aperiodicCSI_RS_AdditionalBandwidth_r17 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *BandNR_aperiodicCSI_RS_AdditionalBandwidth_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode BandNR_aperiodicCSI_RS_AdditionalBandwidth_r17", err)
	}
	return nil
}

func (ie *BandNR_aperiodicCSI_RS_AdditionalBandwidth_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode BandNR_aperiodicCSI_RS_AdditionalBandwidth_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
