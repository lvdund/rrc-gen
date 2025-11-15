package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_Parameters struct {
	phy_ParametersCommon   *Phy_ParametersCommon   `optional`
	phy_ParametersXDD_Diff *Phy_ParametersXDD_Diff `optional`
	phy_ParametersFRX_Diff *Phy_ParametersFRX_Diff `optional`
	phy_ParametersFR1      *Phy_ParametersFR1      `optional`
	phy_ParametersFR2      *Phy_ParametersFR2      `optional`
}

func (ie *Phy_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.phy_ParametersCommon != nil, ie.phy_ParametersXDD_Diff != nil, ie.phy_ParametersFRX_Diff != nil, ie.phy_ParametersFR1 != nil, ie.phy_ParametersFR2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.phy_ParametersCommon != nil {
		if err = ie.phy_ParametersCommon.Encode(w); err != nil {
			return utils.WrapError("Encode phy_ParametersCommon", err)
		}
	}
	if ie.phy_ParametersXDD_Diff != nil {
		if err = ie.phy_ParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode phy_ParametersXDD_Diff", err)
		}
	}
	if ie.phy_ParametersFRX_Diff != nil {
		if err = ie.phy_ParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode phy_ParametersFRX_Diff", err)
		}
	}
	if ie.phy_ParametersFR1 != nil {
		if err = ie.phy_ParametersFR1.Encode(w); err != nil {
			return utils.WrapError("Encode phy_ParametersFR1", err)
		}
	}
	if ie.phy_ParametersFR2 != nil {
		if err = ie.phy_ParametersFR2.Encode(w); err != nil {
			return utils.WrapError("Encode phy_ParametersFR2", err)
		}
	}
	return nil
}

func (ie *Phy_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var phy_ParametersCommonPresent bool
	if phy_ParametersCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var phy_ParametersXDD_DiffPresent bool
	if phy_ParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var phy_ParametersFRX_DiffPresent bool
	if phy_ParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var phy_ParametersFR1Present bool
	if phy_ParametersFR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var phy_ParametersFR2Present bool
	if phy_ParametersFR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if phy_ParametersCommonPresent {
		ie.phy_ParametersCommon = new(Phy_ParametersCommon)
		if err = ie.phy_ParametersCommon.Decode(r); err != nil {
			return utils.WrapError("Decode phy_ParametersCommon", err)
		}
	}
	if phy_ParametersXDD_DiffPresent {
		ie.phy_ParametersXDD_Diff = new(Phy_ParametersXDD_Diff)
		if err = ie.phy_ParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode phy_ParametersXDD_Diff", err)
		}
	}
	if phy_ParametersFRX_DiffPresent {
		ie.phy_ParametersFRX_Diff = new(Phy_ParametersFRX_Diff)
		if err = ie.phy_ParametersFRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode phy_ParametersFRX_Diff", err)
		}
	}
	if phy_ParametersFR1Present {
		ie.phy_ParametersFR1 = new(Phy_ParametersFR1)
		if err = ie.phy_ParametersFR1.Decode(r); err != nil {
			return utils.WrapError("Decode phy_ParametersFR1", err)
		}
	}
	if phy_ParametersFR2Present {
		ie.phy_ParametersFR2 = new(Phy_ParametersFR2)
		if err = ie.phy_ParametersFR2.Decode(r); err != nil {
			return utils.WrapError("Decode phy_ParametersFR2", err)
		}
	}
	return nil
}
