package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SI_RequestResources struct {
	ra_PreambleStartIndex     int64  `lb:0,ub:63,madatory`
	ra_AssociationPeriodIndex *int64 `lb:0,ub:15,optional`
	ra_ssb_OccasionMaskIndex  *int64 `lb:0,ub:15,optional`
}

func (ie *SI_RequestResources) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ra_AssociationPeriodIndex != nil, ie.ra_ssb_OccasionMaskIndex != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.ra_PreambleStartIndex, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger ra_PreambleStartIndex", err)
	}
	if ie.ra_AssociationPeriodIndex != nil {
		if err = w.WriteInteger(*ie.ra_AssociationPeriodIndex, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode ra_AssociationPeriodIndex", err)
		}
	}
	if ie.ra_ssb_OccasionMaskIndex != nil {
		if err = w.WriteInteger(*ie.ra_ssb_OccasionMaskIndex, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode ra_ssb_OccasionMaskIndex", err)
		}
	}
	return nil
}

func (ie *SI_RequestResources) Decode(r *uper.UperReader) error {
	var err error
	var ra_AssociationPeriodIndexPresent bool
	if ra_AssociationPeriodIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_ssb_OccasionMaskIndexPresent bool
	if ra_ssb_OccasionMaskIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_ra_PreambleStartIndex int64
	if tmp_int_ra_PreambleStartIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger ra_PreambleStartIndex", err)
	}
	ie.ra_PreambleStartIndex = tmp_int_ra_PreambleStartIndex
	if ra_AssociationPeriodIndexPresent {
		var tmp_int_ra_AssociationPeriodIndex int64
		if tmp_int_ra_AssociationPeriodIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode ra_AssociationPeriodIndex", err)
		}
		ie.ra_AssociationPeriodIndex = &tmp_int_ra_AssociationPeriodIndex
	}
	if ra_ssb_OccasionMaskIndexPresent {
		var tmp_int_ra_ssb_OccasionMaskIndex int64
		if tmp_int_ra_ssb_OccasionMaskIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode ra_ssb_OccasionMaskIndex", err)
		}
		ie.ra_ssb_OccasionMaskIndex = &tmp_int_ra_ssb_OccasionMaskIndex
	}
	return nil
}
