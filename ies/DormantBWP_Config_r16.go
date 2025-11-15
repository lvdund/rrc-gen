package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DormantBWP_Config_r16 struct {
	dormantBWP_Id_r16           *BWP_Id                      `optional`
	withinActiveTimeConfig_r16  *WithinActiveTimeConfig_r16  `optional,setuprelease`
	outsideActiveTimeConfig_r16 *OutsideActiveTimeConfig_r16 `optional,setuprelease`
}

func (ie *DormantBWP_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dormantBWP_Id_r16 != nil, ie.withinActiveTimeConfig_r16 != nil, ie.outsideActiveTimeConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dormantBWP_Id_r16 != nil {
		if err = ie.dormantBWP_Id_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dormantBWP_Id_r16", err)
		}
	}
	if ie.withinActiveTimeConfig_r16 != nil {
		tmp_withinActiveTimeConfig_r16 := utils.SetupRelease[*WithinActiveTimeConfig_r16]{
			Setup: ie.withinActiveTimeConfig_r16,
		}
		if err = tmp_withinActiveTimeConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode withinActiveTimeConfig_r16", err)
		}
	}
	if ie.outsideActiveTimeConfig_r16 != nil {
		tmp_outsideActiveTimeConfig_r16 := utils.SetupRelease[*OutsideActiveTimeConfig_r16]{
			Setup: ie.outsideActiveTimeConfig_r16,
		}
		if err = tmp_outsideActiveTimeConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode outsideActiveTimeConfig_r16", err)
		}
	}
	return nil
}

func (ie *DormantBWP_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var dormantBWP_Id_r16Present bool
	if dormantBWP_Id_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var withinActiveTimeConfig_r16Present bool
	if withinActiveTimeConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var outsideActiveTimeConfig_r16Present bool
	if outsideActiveTimeConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dormantBWP_Id_r16Present {
		ie.dormantBWP_Id_r16 = new(BWP_Id)
		if err = ie.dormantBWP_Id_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dormantBWP_Id_r16", err)
		}
	}
	if withinActiveTimeConfig_r16Present {
		tmp_withinActiveTimeConfig_r16 := utils.SetupRelease[*WithinActiveTimeConfig_r16]{}
		if err = tmp_withinActiveTimeConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode withinActiveTimeConfig_r16", err)
		}
		ie.withinActiveTimeConfig_r16 = tmp_withinActiveTimeConfig_r16.Setup
	}
	if outsideActiveTimeConfig_r16Present {
		tmp_outsideActiveTimeConfig_r16 := utils.SetupRelease[*OutsideActiveTimeConfig_r16]{}
		if err = tmp_outsideActiveTimeConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode outsideActiveTimeConfig_r16", err)
		}
		ie.outsideActiveTimeConfig_r16 = tmp_outsideActiveTimeConfig_r16.Setup
	}
	return nil
}
