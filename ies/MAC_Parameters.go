package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_Parameters struct {
	mac_ParametersCommon   *MAC_ParametersCommon   `optional`
	mac_ParametersXDD_Diff *MAC_ParametersXDD_Diff `optional`
}

func (ie *MAC_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mac_ParametersCommon != nil, ie.mac_ParametersXDD_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mac_ParametersCommon != nil {
		if err = ie.mac_ParametersCommon.Encode(w); err != nil {
			return utils.WrapError("Encode mac_ParametersCommon", err)
		}
	}
	if ie.mac_ParametersXDD_Diff != nil {
		if err = ie.mac_ParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode mac_ParametersXDD_Diff", err)
		}
	}
	return nil
}

func (ie *MAC_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var mac_ParametersCommonPresent bool
	if mac_ParametersCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_ParametersXDD_DiffPresent bool
	if mac_ParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if mac_ParametersCommonPresent {
		ie.mac_ParametersCommon = new(MAC_ParametersCommon)
		if err = ie.mac_ParametersCommon.Decode(r); err != nil {
			return utils.WrapError("Decode mac_ParametersCommon", err)
		}
	}
	if mac_ParametersXDD_DiffPresent {
		ie.mac_ParametersXDD_Diff = new(MAC_ParametersXDD_Diff)
		if err = ie.mac_ParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode mac_ParametersXDD_Diff", err)
		}
	}
	return nil
}
