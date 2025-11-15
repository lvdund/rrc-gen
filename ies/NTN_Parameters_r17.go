package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NTN_Parameters_r17 struct {
	inactiveStateNTN_r17               *NTN_Parameters_r17_inactiveStateNTN_r17 `optional`
	ra_SDT_NTN_r17                     *NTN_Parameters_r17_ra_SDT_NTN_r17       `optional`
	srb_SDT_NTN_r17                    *NTN_Parameters_r17_srb_SDT_NTN_r17      `optional`
	measAndMobParametersNTN_r17        *MeasAndMobParameters                    `optional`
	mac_ParametersNTN_r17              *MAC_Parameters                          `optional`
	phy_ParametersNTN_r17              *Phy_Parameters                          `optional`
	fdd_Add_UE_NR_CapabilitiesNTN_r17  *UE_NR_CapabilityAddXDD_Mode             `optional`
	fr1_Add_UE_NR_CapabilitiesNTN_r17  *UE_NR_CapabilityAddFRX_Mode             `optional`
	ue_BasedPerfMeas_ParametersNTN_r17 *UE_BasedPerfMeas_Parameters_r16         `optional`
	son_ParametersNTN_r17              *SON_Parameters_r16                      `optional`
}

func (ie *NTN_Parameters_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.inactiveStateNTN_r17 != nil, ie.ra_SDT_NTN_r17 != nil, ie.srb_SDT_NTN_r17 != nil, ie.measAndMobParametersNTN_r17 != nil, ie.mac_ParametersNTN_r17 != nil, ie.phy_ParametersNTN_r17 != nil, ie.fdd_Add_UE_NR_CapabilitiesNTN_r17 != nil, ie.fr1_Add_UE_NR_CapabilitiesNTN_r17 != nil, ie.ue_BasedPerfMeas_ParametersNTN_r17 != nil, ie.son_ParametersNTN_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.inactiveStateNTN_r17 != nil {
		if err = ie.inactiveStateNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode inactiveStateNTN_r17", err)
		}
	}
	if ie.ra_SDT_NTN_r17 != nil {
		if err = ie.ra_SDT_NTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ra_SDT_NTN_r17", err)
		}
	}
	if ie.srb_SDT_NTN_r17 != nil {
		if err = ie.srb_SDT_NTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode srb_SDT_NTN_r17", err)
		}
	}
	if ie.measAndMobParametersNTN_r17 != nil {
		if err = ie.measAndMobParametersNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersNTN_r17", err)
		}
	}
	if ie.mac_ParametersNTN_r17 != nil {
		if err = ie.mac_ParametersNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mac_ParametersNTN_r17", err)
		}
	}
	if ie.phy_ParametersNTN_r17 != nil {
		if err = ie.phy_ParametersNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode phy_ParametersNTN_r17", err)
		}
	}
	if ie.fdd_Add_UE_NR_CapabilitiesNTN_r17 != nil {
		if err = ie.fdd_Add_UE_NR_CapabilitiesNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fdd_Add_UE_NR_CapabilitiesNTN_r17", err)
		}
	}
	if ie.fr1_Add_UE_NR_CapabilitiesNTN_r17 != nil {
		if err = ie.fr1_Add_UE_NR_CapabilitiesNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_Add_UE_NR_CapabilitiesNTN_r17", err)
		}
	}
	if ie.ue_BasedPerfMeas_ParametersNTN_r17 != nil {
		if err = ie.ue_BasedPerfMeas_ParametersNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ue_BasedPerfMeas_ParametersNTN_r17", err)
		}
	}
	if ie.son_ParametersNTN_r17 != nil {
		if err = ie.son_ParametersNTN_r17.Encode(w); err != nil {
			return utils.WrapError("Encode son_ParametersNTN_r17", err)
		}
	}
	return nil
}

func (ie *NTN_Parameters_r17) Decode(r *uper.UperReader) error {
	var err error
	var inactiveStateNTN_r17Present bool
	if inactiveStateNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_SDT_NTN_r17Present bool
	if ra_SDT_NTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var srb_SDT_NTN_r17Present bool
	if srb_SDT_NTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measAndMobParametersNTN_r17Present bool
	if measAndMobParametersNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_ParametersNTN_r17Present bool
	if mac_ParametersNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var phy_ParametersNTN_r17Present bool
	if phy_ParametersNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fdd_Add_UE_NR_CapabilitiesNTN_r17Present bool
	if fdd_Add_UE_NR_CapabilitiesNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1_Add_UE_NR_CapabilitiesNTN_r17Present bool
	if fr1_Add_UE_NR_CapabilitiesNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_BasedPerfMeas_ParametersNTN_r17Present bool
	if ue_BasedPerfMeas_ParametersNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var son_ParametersNTN_r17Present bool
	if son_ParametersNTN_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if inactiveStateNTN_r17Present {
		ie.inactiveStateNTN_r17 = new(NTN_Parameters_r17_inactiveStateNTN_r17)
		if err = ie.inactiveStateNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode inactiveStateNTN_r17", err)
		}
	}
	if ra_SDT_NTN_r17Present {
		ie.ra_SDT_NTN_r17 = new(NTN_Parameters_r17_ra_SDT_NTN_r17)
		if err = ie.ra_SDT_NTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ra_SDT_NTN_r17", err)
		}
	}
	if srb_SDT_NTN_r17Present {
		ie.srb_SDT_NTN_r17 = new(NTN_Parameters_r17_srb_SDT_NTN_r17)
		if err = ie.srb_SDT_NTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode srb_SDT_NTN_r17", err)
		}
	}
	if measAndMobParametersNTN_r17Present {
		ie.measAndMobParametersNTN_r17 = new(MeasAndMobParameters)
		if err = ie.measAndMobParametersNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersNTN_r17", err)
		}
	}
	if mac_ParametersNTN_r17Present {
		ie.mac_ParametersNTN_r17 = new(MAC_Parameters)
		if err = ie.mac_ParametersNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mac_ParametersNTN_r17", err)
		}
	}
	if phy_ParametersNTN_r17Present {
		ie.phy_ParametersNTN_r17 = new(Phy_Parameters)
		if err = ie.phy_ParametersNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode phy_ParametersNTN_r17", err)
		}
	}
	if fdd_Add_UE_NR_CapabilitiesNTN_r17Present {
		ie.fdd_Add_UE_NR_CapabilitiesNTN_r17 = new(UE_NR_CapabilityAddXDD_Mode)
		if err = ie.fdd_Add_UE_NR_CapabilitiesNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fdd_Add_UE_NR_CapabilitiesNTN_r17", err)
		}
	}
	if fr1_Add_UE_NR_CapabilitiesNTN_r17Present {
		ie.fr1_Add_UE_NR_CapabilitiesNTN_r17 = new(UE_NR_CapabilityAddFRX_Mode)
		if err = ie.fr1_Add_UE_NR_CapabilitiesNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_Add_UE_NR_CapabilitiesNTN_r17", err)
		}
	}
	if ue_BasedPerfMeas_ParametersNTN_r17Present {
		ie.ue_BasedPerfMeas_ParametersNTN_r17 = new(UE_BasedPerfMeas_Parameters_r16)
		if err = ie.ue_BasedPerfMeas_ParametersNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ue_BasedPerfMeas_ParametersNTN_r17", err)
		}
	}
	if son_ParametersNTN_r17Present {
		ie.son_ParametersNTN_r17 = new(SON_Parameters_r16)
		if err = ie.son_ParametersNTN_r17.Decode(r); err != nil {
			return utils.WrapError("Decode son_ParametersNTN_r17", err)
		}
	}
	return nil
}
