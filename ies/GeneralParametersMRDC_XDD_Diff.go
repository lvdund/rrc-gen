package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GeneralParametersMRDC_XDD_Diff struct {
	splitSRB_WithOneUL_Path      *GeneralParametersMRDC_XDD_Diff_splitSRB_WithOneUL_Path      `optional`
	splitDRB_withUL_Both_MCG_SCG *GeneralParametersMRDC_XDD_Diff_splitDRB_withUL_Both_MCG_SCG `optional`
	srb3                         *GeneralParametersMRDC_XDD_Diff_srb3                         `optional`
	dummy                        *GeneralParametersMRDC_XDD_Diff_dummy                        `optional`
}

func (ie *GeneralParametersMRDC_XDD_Diff) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.splitSRB_WithOneUL_Path != nil, ie.splitDRB_withUL_Both_MCG_SCG != nil, ie.srb3 != nil, ie.dummy != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.splitSRB_WithOneUL_Path != nil {
		if err = ie.splitSRB_WithOneUL_Path.Encode(w); err != nil {
			return utils.WrapError("Encode splitSRB_WithOneUL_Path", err)
		}
	}
	if ie.splitDRB_withUL_Both_MCG_SCG != nil {
		if err = ie.splitDRB_withUL_Both_MCG_SCG.Encode(w); err != nil {
			return utils.WrapError("Encode splitDRB_withUL_Both_MCG_SCG", err)
		}
	}
	if ie.srb3 != nil {
		if err = ie.srb3.Encode(w); err != nil {
			return utils.WrapError("Encode srb3", err)
		}
	}
	if ie.dummy != nil {
		if err = ie.dummy.Encode(w); err != nil {
			return utils.WrapError("Encode dummy", err)
		}
	}
	return nil
}

func (ie *GeneralParametersMRDC_XDD_Diff) Decode(r *uper.UperReader) error {
	var err error
	var splitSRB_WithOneUL_PathPresent bool
	if splitSRB_WithOneUL_PathPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var splitDRB_withUL_Both_MCG_SCGPresent bool
	if splitDRB_withUL_Both_MCG_SCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var srb3Present bool
	if srb3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummyPresent bool
	if dummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if splitSRB_WithOneUL_PathPresent {
		ie.splitSRB_WithOneUL_Path = new(GeneralParametersMRDC_XDD_Diff_splitSRB_WithOneUL_Path)
		if err = ie.splitSRB_WithOneUL_Path.Decode(r); err != nil {
			return utils.WrapError("Decode splitSRB_WithOneUL_Path", err)
		}
	}
	if splitDRB_withUL_Both_MCG_SCGPresent {
		ie.splitDRB_withUL_Both_MCG_SCG = new(GeneralParametersMRDC_XDD_Diff_splitDRB_withUL_Both_MCG_SCG)
		if err = ie.splitDRB_withUL_Both_MCG_SCG.Decode(r); err != nil {
			return utils.WrapError("Decode splitDRB_withUL_Both_MCG_SCG", err)
		}
	}
	if srb3Present {
		ie.srb3 = new(GeneralParametersMRDC_XDD_Diff_srb3)
		if err = ie.srb3.Decode(r); err != nil {
			return utils.WrapError("Decode srb3", err)
		}
	}
	if dummyPresent {
		ie.dummy = new(GeneralParametersMRDC_XDD_Diff_dummy)
		if err = ie.dummy.Decode(r); err != nil {
			return utils.WrapError("Decode dummy", err)
		}
	}
	return nil
}
