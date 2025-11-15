package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SDT_Config_r17 struct {
	sdt_DRB_List_r17          []DRB_Identity                           `lb:0,ub:maxDRB,optional`
	sdt_SRB2_Indication_r17   *SDT_Config_r17_sdt_SRB2_Indication_r17  `optional`
	sdt_MAC_PHY_CG_Config_r17 *SDT_CG_Config_r17                       `optional,setuprelease`
	sdt_DRB_ContinueROHC_r17  *SDT_Config_r17_sdt_DRB_ContinueROHC_r17 `optional`
}

func (ie *SDT_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.sdt_DRB_List_r17) > 0, ie.sdt_SRB2_Indication_r17 != nil, ie.sdt_MAC_PHY_CG_Config_r17 != nil, ie.sdt_DRB_ContinueROHC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.sdt_DRB_List_r17) > 0 {
		tmp_sdt_DRB_List_r17 := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, uper.Constraint{Lb: 0, Ub: maxDRB}, false)
		for _, i := range ie.sdt_DRB_List_r17 {
			tmp_sdt_DRB_List_r17.Value = append(tmp_sdt_DRB_List_r17.Value, &i)
		}
		if err = tmp_sdt_DRB_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_DRB_List_r17", err)
		}
	}
	if ie.sdt_SRB2_Indication_r17 != nil {
		if err = ie.sdt_SRB2_Indication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_SRB2_Indication_r17", err)
		}
	}
	if ie.sdt_MAC_PHY_CG_Config_r17 != nil {
		tmp_sdt_MAC_PHY_CG_Config_r17 := utils.SetupRelease[*SDT_CG_Config_r17]{
			Setup: ie.sdt_MAC_PHY_CG_Config_r17,
		}
		if err = tmp_sdt_MAC_PHY_CG_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_MAC_PHY_CG_Config_r17", err)
		}
	}
	if ie.sdt_DRB_ContinueROHC_r17 != nil {
		if err = ie.sdt_DRB_ContinueROHC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sdt_DRB_ContinueROHC_r17", err)
		}
	}
	return nil
}

func (ie *SDT_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var sdt_DRB_List_r17Present bool
	if sdt_DRB_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_SRB2_Indication_r17Present bool
	if sdt_SRB2_Indication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_MAC_PHY_CG_Config_r17Present bool
	if sdt_MAC_PHY_CG_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sdt_DRB_ContinueROHC_r17Present bool
	if sdt_DRB_ContinueROHC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sdt_DRB_List_r17Present {
		tmp_sdt_DRB_List_r17 := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, uper.Constraint{Lb: 0, Ub: maxDRB}, false)
		fn_sdt_DRB_List_r17 := func() *DRB_Identity {
			return new(DRB_Identity)
		}
		if err = tmp_sdt_DRB_List_r17.Decode(r, fn_sdt_DRB_List_r17); err != nil {
			return utils.WrapError("Decode sdt_DRB_List_r17", err)
		}
		ie.sdt_DRB_List_r17 = []DRB_Identity{}
		for _, i := range tmp_sdt_DRB_List_r17.Value {
			ie.sdt_DRB_List_r17 = append(ie.sdt_DRB_List_r17, *i)
		}
	}
	if sdt_SRB2_Indication_r17Present {
		ie.sdt_SRB2_Indication_r17 = new(SDT_Config_r17_sdt_SRB2_Indication_r17)
		if err = ie.sdt_SRB2_Indication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sdt_SRB2_Indication_r17", err)
		}
	}
	if sdt_MAC_PHY_CG_Config_r17Present {
		tmp_sdt_MAC_PHY_CG_Config_r17 := utils.SetupRelease[*SDT_CG_Config_r17]{}
		if err = tmp_sdt_MAC_PHY_CG_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sdt_MAC_PHY_CG_Config_r17", err)
		}
		ie.sdt_MAC_PHY_CG_Config_r17 = tmp_sdt_MAC_PHY_CG_Config_r17.Setup
	}
	if sdt_DRB_ContinueROHC_r17Present {
		ie.sdt_DRB_ContinueROHC_r17 = new(SDT_Config_r17_sdt_DRB_ContinueROHC_r17)
		if err = ie.sdt_DRB_ContinueROHC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sdt_DRB_ContinueROHC_r17", err)
		}
	}
	return nil
}
