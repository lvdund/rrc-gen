package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n8   uper.Enumerated = 0
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n16  uper.Enumerated = 1
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n32  uper.Enumerated = 2
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n48  uper.Enumerated = 3
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n64  uper.Enumerated = 4
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n128 uper.Enumerated = 5
	BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17_Enum_n255 uper.Enumerated = 6
)

type BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17", err)
	}
	return nil
}

func (ie *BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
