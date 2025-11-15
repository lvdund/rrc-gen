package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_aperiodicCSI_RS_FastScellActivation_r17 struct {
	maxNumberAperiodicCSI_RS_PerCC_r17     BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_PerCC_r17     `madatory`
	maxNumberAperiodicCSI_RS_AcrossCCs_r17 BandNR_aperiodicCSI_RS_FastScellActivation_r17_maxNumberAperiodicCSI_RS_AcrossCCs_r17 `madatory`
}

func (ie *BandNR_aperiodicCSI_RS_FastScellActivation_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberAperiodicCSI_RS_PerCC_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberAperiodicCSI_RS_PerCC_r17", err)
	}
	if err = ie.maxNumberAperiodicCSI_RS_AcrossCCs_r17.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberAperiodicCSI_RS_AcrossCCs_r17", err)
	}
	return nil
}

func (ie *BandNR_aperiodicCSI_RS_FastScellActivation_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberAperiodicCSI_RS_PerCC_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberAperiodicCSI_RS_PerCC_r17", err)
	}
	if err = ie.maxNumberAperiodicCSI_RS_AcrossCCs_r17.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberAperiodicCSI_RS_AcrossCCs_r17", err)
	}
	return nil
}
