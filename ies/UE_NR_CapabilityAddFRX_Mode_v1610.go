package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_CapabilityAddFRX_Mode_v1610 struct {
	powSav_ParametersFRX_Diff_r16 *PowSav_ParametersFRX_Diff_r16 `optional`
	mac_ParametersFRX_Diff_r16    *MAC_ParametersFRX_Diff_r16    `optional`
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.powSav_ParametersFRX_Diff_r16 != nil, ie.mac_ParametersFRX_Diff_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.powSav_ParametersFRX_Diff_r16 != nil {
		if err = ie.powSav_ParametersFRX_Diff_r16.Encode(w); err != nil {
			return utils.WrapError("Encode powSav_ParametersFRX_Diff_r16", err)
		}
	}
	if ie.mac_ParametersFRX_Diff_r16 != nil {
		if err = ie.mac_ParametersFRX_Diff_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mac_ParametersFRX_Diff_r16", err)
		}
	}
	return nil
}

func (ie *UE_NR_CapabilityAddFRX_Mode_v1610) Decode(r *uper.UperReader) error {
	var err error
	var powSav_ParametersFRX_Diff_r16Present bool
	if powSav_ParametersFRX_Diff_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_ParametersFRX_Diff_r16Present bool
	if mac_ParametersFRX_Diff_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if powSav_ParametersFRX_Diff_r16Present {
		ie.powSav_ParametersFRX_Diff_r16 = new(PowSav_ParametersFRX_Diff_r16)
		if err = ie.powSav_ParametersFRX_Diff_r16.Decode(r); err != nil {
			return utils.WrapError("Decode powSav_ParametersFRX_Diff_r16", err)
		}
	}
	if mac_ParametersFRX_Diff_r16Present {
		ie.mac_ParametersFRX_Diff_r16 = new(MAC_ParametersFRX_Diff_r16)
		if err = ie.mac_ParametersFRX_Diff_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mac_ParametersFRX_Diff_r16", err)
		}
	}
	return nil
}
