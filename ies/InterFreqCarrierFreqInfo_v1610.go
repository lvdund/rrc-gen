package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqInfo_v1610 struct {
	interFreqNeighCellList_v1610 *InterFreqNeighCellList_v1610      `optional`
	smtc2_LP_r16                 *SSB_MTC2_LP_r16                   `optional`
	interFreqAllowedCellList_r16 *InterFreqAllowedCellList_r16      `optional`
	ssb_PositionQCL_Common_r16   *SSB_PositionQCL_Relation_r16      `optional`
	interFreqCAG_CellList_r16    []InterFreqCAG_CellListPerPLMN_r16 `lb:1,ub:maxPLMN,optional`
}

func (ie *InterFreqCarrierFreqInfo_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.interFreqNeighCellList_v1610 != nil, ie.smtc2_LP_r16 != nil, ie.interFreqAllowedCellList_r16 != nil, ie.ssb_PositionQCL_Common_r16 != nil, len(ie.interFreqCAG_CellList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.interFreqNeighCellList_v1610 != nil {
		if err = ie.interFreqNeighCellList_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqNeighCellList_v1610", err)
		}
	}
	if ie.smtc2_LP_r16 != nil {
		if err = ie.smtc2_LP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode smtc2_LP_r16", err)
		}
	}
	if ie.interFreqAllowedCellList_r16 != nil {
		if err = ie.interFreqAllowedCellList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqAllowedCellList_r16", err)
		}
	}
	if ie.ssb_PositionQCL_Common_r16 != nil {
		if err = ie.ssb_PositionQCL_Common_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_PositionQCL_Common_r16", err)
		}
	}
	if len(ie.interFreqCAG_CellList_r16) > 0 {
		tmp_interFreqCAG_CellList_r16 := utils.NewSequence[*InterFreqCAG_CellListPerPLMN_r16]([]*InterFreqCAG_CellListPerPLMN_r16{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		for _, i := range ie.interFreqCAG_CellList_r16 {
			tmp_interFreqCAG_CellList_r16.Value = append(tmp_interFreqCAG_CellList_r16.Value, &i)
		}
		if err = tmp_interFreqCAG_CellList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqCAG_CellList_r16", err)
		}
	}
	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1610) Decode(r *uper.UperReader) error {
	var err error
	var interFreqNeighCellList_v1610Present bool
	if interFreqNeighCellList_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var smtc2_LP_r16Present bool
	if smtc2_LP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqAllowedCellList_r16Present bool
	if interFreqAllowedCellList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_PositionQCL_Common_r16Present bool
	if ssb_PositionQCL_Common_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqCAG_CellList_r16Present bool
	if interFreqCAG_CellList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if interFreqNeighCellList_v1610Present {
		ie.interFreqNeighCellList_v1610 = new(InterFreqNeighCellList_v1610)
		if err = ie.interFreqNeighCellList_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqNeighCellList_v1610", err)
		}
	}
	if smtc2_LP_r16Present {
		ie.smtc2_LP_r16 = new(SSB_MTC2_LP_r16)
		if err = ie.smtc2_LP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode smtc2_LP_r16", err)
		}
	}
	if interFreqAllowedCellList_r16Present {
		ie.interFreqAllowedCellList_r16 = new(InterFreqAllowedCellList_r16)
		if err = ie.interFreqAllowedCellList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqAllowedCellList_r16", err)
		}
	}
	if ssb_PositionQCL_Common_r16Present {
		ie.ssb_PositionQCL_Common_r16 = new(SSB_PositionQCL_Relation_r16)
		if err = ie.ssb_PositionQCL_Common_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_PositionQCL_Common_r16", err)
		}
	}
	if interFreqCAG_CellList_r16Present {
		tmp_interFreqCAG_CellList_r16 := utils.NewSequence[*InterFreqCAG_CellListPerPLMN_r16]([]*InterFreqCAG_CellListPerPLMN_r16{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		fn_interFreqCAG_CellList_r16 := func() *InterFreqCAG_CellListPerPLMN_r16 {
			return new(InterFreqCAG_CellListPerPLMN_r16)
		}
		if err = tmp_interFreqCAG_CellList_r16.Decode(r, fn_interFreqCAG_CellList_r16); err != nil {
			return utils.WrapError("Decode interFreqCAG_CellList_r16", err)
		}
		ie.interFreqCAG_CellList_r16 = []InterFreqCAG_CellListPerPLMN_r16{}
		for _, i := range tmp_interFreqCAG_CellList_r16.Value {
			ie.interFreqCAG_CellList_r16 = append(ie.interFreqCAG_CellList_r16, *i)
		}
	}
	return nil
}
