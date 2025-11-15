package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandNR_bwp_DiffNumerology_Enum_upto4 uper.Enumerated = 0
)

type BandNR_bwp_DiffNumerology struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *BandNR_bwp_DiffNumerology) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode BandNR_bwp_DiffNumerology", err)
	}
	return nil
}

func (ie *BandNR_bwp_DiffNumerology) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode BandNR_bwp_DiffNumerology", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
