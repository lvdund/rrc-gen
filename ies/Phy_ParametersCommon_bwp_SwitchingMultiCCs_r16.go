package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_nothing uint64 = iota
	Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_type1_r16
	Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_type2_r16
)

type Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16 struct {
	Choice    uint64
	type1_r16 *Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_type1_r16
	type2_r16 *Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_type2_r16
}

func (ie *Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_type1_r16:
		if err = ie.type1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode type1_r16", err)
		}
	case Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_type2_r16:
		if err = ie.type2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode type2_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_type1_r16:
		ie.type1_r16 = new(Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_type1_r16)
		if err = ie.type1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode type1_r16", err)
		}
	case Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_type2_r16:
		ie.type2_r16 = new(Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_type2_r16)
		if err = ie.type2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode type2_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
