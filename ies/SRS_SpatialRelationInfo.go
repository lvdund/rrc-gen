package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_SpatialRelationInfo struct {
	servingCellId   *ServCellIndex                          `optional`
	referenceSignal SRS_SpatialRelationInfo_referenceSignal `madatory`
}

func (ie *SRS_SpatialRelationInfo) Encode(w *uper.UperWriter) error {
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
	if err = ie.referenceSignal.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignal", err)
	}
	return nil
}

func (ie *SRS_SpatialRelationInfo) Decode(r *uper.UperReader) error {
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
	if err = ie.referenceSignal.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignal", err)
	}
	return nil
}
