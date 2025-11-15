package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_CapabilityAddXDD_Mode struct {
	phy_ParametersXDD_Diff       *Phy_ParametersXDD_Diff       `optional`
	mac_ParametersXDD_Diff       *MAC_ParametersXDD_Diff       `optional`
	measAndMobParametersXDD_Diff *MeasAndMobParametersXDD_Diff `optional`
}

func (ie *UE_NR_CapabilityAddXDD_Mode) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.phy_ParametersXDD_Diff != nil, ie.mac_ParametersXDD_Diff != nil, ie.measAndMobParametersXDD_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.phy_ParametersXDD_Diff != nil {
		if err = ie.phy_ParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode phy_ParametersXDD_Diff", err)
		}
	}
	if ie.mac_ParametersXDD_Diff != nil {
		if err = ie.mac_ParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode mac_ParametersXDD_Diff", err)
		}
	}
	if ie.measAndMobParametersXDD_Diff != nil {
		if err = ie.measAndMobParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersXDD_Diff", err)
		}
	}
	return nil
}

func (ie *UE_NR_CapabilityAddXDD_Mode) Decode(r *uper.UperReader) error {
	var err error
	var phy_ParametersXDD_DiffPresent bool
	if phy_ParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_ParametersXDD_DiffPresent bool
	if mac_ParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measAndMobParametersXDD_DiffPresent bool
	if measAndMobParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if phy_ParametersXDD_DiffPresent {
		ie.phy_ParametersXDD_Diff = new(Phy_ParametersXDD_Diff)
		if err = ie.phy_ParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode phy_ParametersXDD_Diff", err)
		}
	}
	if mac_ParametersXDD_DiffPresent {
		ie.mac_ParametersXDD_Diff = new(MAC_ParametersXDD_Diff)
		if err = ie.mac_ParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode mac_ParametersXDD_Diff", err)
		}
	}
	if measAndMobParametersXDD_DiffPresent {
		ie.measAndMobParametersXDD_Diff = new(MeasAndMobParametersXDD_Diff)
		if err = ie.measAndMobParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersXDD_Diff", err)
		}
	}
	return nil
}
