package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_SidelinkCapabilityAddXDD_Mode_r16 struct {
	mac_ParametersSidelinkXDD_Diff_r16 *MAC_ParametersSidelinkXDD_Diff_r16 `optional`
}

func (ie *UE_SidelinkCapabilityAddXDD_Mode_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mac_ParametersSidelinkXDD_Diff_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mac_ParametersSidelinkXDD_Diff_r16 != nil {
		if err = ie.mac_ParametersSidelinkXDD_Diff_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mac_ParametersSidelinkXDD_Diff_r16", err)
		}
	}
	return nil
}

func (ie *UE_SidelinkCapabilityAddXDD_Mode_r16) Decode(r *uper.UperReader) error {
	var err error
	var mac_ParametersSidelinkXDD_Diff_r16Present bool
	if mac_ParametersSidelinkXDD_Diff_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if mac_ParametersSidelinkXDD_Diff_r16Present {
		ie.mac_ParametersSidelinkXDD_Diff_r16 = new(MAC_ParametersSidelinkXDD_Diff_r16)
		if err = ie.mac_ParametersSidelinkXDD_Diff_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mac_ParametersSidelinkXDD_Diff_r16", err)
		}
	}
	return nil
}
