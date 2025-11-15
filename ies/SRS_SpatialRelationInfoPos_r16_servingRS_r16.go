package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_SpatialRelationInfoPos_r16_servingRS_r16 struct {
	servingCellId       *ServCellIndex                                                   `optional`
	referenceSignal_r16 SRS_SpatialRelationInfoPos_r16_servingRS_r16_referenceSignal_r16 `madatory`
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.servingCellId != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.servingCellId != nil {
		if err = ie.servingCellId.Encode(w); err != nil {
			return utils.WrapError("Encode servingCellId", err)
		}
	}
	if err = ie.referenceSignal_r16.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignal_r16", err)
	}
	return nil
}

func (ie *SRS_SpatialRelationInfoPos_r16_servingRS_r16) Decode(r *uper.UperReader) error {
	var err error
	var servingCellIdPresent bool
	if servingCellIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if servingCellIdPresent {
		ie.servingCellId = new(ServCellIndex)
		if err = ie.servingCellId.Decode(r); err != nil {
			return utils.WrapError("Decode servingCellId", err)
		}
	}
	if err = ie.referenceSignal_r16.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignal_r16", err)
	}
	return nil
}
