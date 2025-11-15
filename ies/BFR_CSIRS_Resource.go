package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BFR_CSIRS_Resource struct {
	csi_RS           NZP_CSI_RS_ResourceId `madatory`
	ra_OccasionList  []int64               `lb:1,ub:maxRA_OccasionsPerCSIRS,e_lb:0,e_ub:maxRA_Occasions_1,optional`
	ra_PreambleIndex *int64                `lb:0,ub:63,optional`
}

func (ie *BFR_CSIRS_Resource) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.ra_OccasionList) > 0, ie.ra_PreambleIndex != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.csi_RS.Encode(w); err != nil {
		return utils.WrapError("Encode csi_RS", err)
	}
	if len(ie.ra_OccasionList) > 0 {
		tmp_ra_OccasionList := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxRA_OccasionsPerCSIRS}, false)
		for _, i := range ie.ra_OccasionList {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: maxRA_Occasions_1}, false)
			tmp_ra_OccasionList.Value = append(tmp_ra_OccasionList.Value, &tmp_ie)
		}
		if err = tmp_ra_OccasionList.Encode(w); err != nil {
			return utils.WrapError("Encode ra_OccasionList", err)
		}
	}
	if ie.ra_PreambleIndex != nil {
		if err = w.WriteInteger(*ie.ra_PreambleIndex, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode ra_PreambleIndex", err)
		}
	}
	return nil
}

func (ie *BFR_CSIRS_Resource) Decode(r *uper.UperReader) error {
	var err error
	var ra_OccasionListPresent bool
	if ra_OccasionListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_PreambleIndexPresent bool
	if ra_PreambleIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.csi_RS.Decode(r); err != nil {
		return utils.WrapError("Decode csi_RS", err)
	}
	if ra_OccasionListPresent {
		tmp_ra_OccasionList := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxRA_OccasionsPerCSIRS}, false)
		fn_ra_OccasionList := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: maxRA_Occasions_1}, false)
			return &ie
		}
		if err = tmp_ra_OccasionList.Decode(r, fn_ra_OccasionList); err != nil {
			return utils.WrapError("Decode ra_OccasionList", err)
		}
		ie.ra_OccasionList = []int64{}
		for _, i := range tmp_ra_OccasionList.Value {
			ie.ra_OccasionList = append(ie.ra_OccasionList, int64(i.Value))
		}
	}
	if ra_PreambleIndexPresent {
		var tmp_int_ra_PreambleIndex int64
		if tmp_int_ra_PreambleIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode ra_PreambleIndex", err)
		}
		ie.ra_PreambleIndex = &tmp_int_ra_PreambleIndex
	}
	return nil
}
