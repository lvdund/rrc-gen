package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_CapabilityAddFRX_Mode struct {
	phy_ParametersFRX_Diff       *Phy_ParametersFRX_Diff       `optional`
	measAndMobParametersFRX_Diff *MeasAndMobParametersFRX_Diff `optional`
}

func (ie *UE_NR_CapabilityAddFRX_Mode) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.phy_ParametersFRX_Diff != nil, ie.measAndMobParametersFRX_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.phy_ParametersFRX_Diff != nil {
		if err = ie.phy_ParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode phy_ParametersFRX_Diff", err)
		}
	}
	if ie.measAndMobParametersFRX_Diff != nil {
		if err = ie.measAndMobParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersFRX_Diff", err)
		}
	}
	return nil
}

func (ie *UE_NR_CapabilityAddFRX_Mode) Decode(r *uper.UperReader) error {
	var err error
	var phy_ParametersFRX_DiffPresent bool
	if phy_ParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measAndMobParametersFRX_DiffPresent bool
	if measAndMobParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if phy_ParametersFRX_DiffPresent {
		ie.phy_ParametersFRX_Diff = new(Phy_ParametersFRX_Diff)
		if err = ie.phy_ParametersFRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode phy_ParametersFRX_Diff", err)
		}
	}
	if measAndMobParametersFRX_DiffPresent {
		ie.measAndMobParametersFRX_Diff = new(MeasAndMobParametersFRX_Diff)
		if err = ie.measAndMobParametersFRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersFRX_Diff", err)
		}
	}
	return nil
}
