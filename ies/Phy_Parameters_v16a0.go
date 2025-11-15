package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_Parameters_v16a0 struct {
	phy_ParametersCommon_v16a0 *Phy_ParametersCommon_v16a0 `optional`
}

func (ie *Phy_Parameters_v16a0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.phy_ParametersCommon_v16a0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.phy_ParametersCommon_v16a0 != nil {
		if err = ie.phy_ParametersCommon_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode phy_ParametersCommon_v16a0", err)
		}
	}
	return nil
}

func (ie *Phy_Parameters_v16a0) Decode(r *uper.UperReader) error {
	var err error
	var phy_ParametersCommon_v16a0Present bool
	if phy_ParametersCommon_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if phy_ParametersCommon_v16a0Present {
		ie.phy_ParametersCommon_v16a0 = new(Phy_ParametersCommon_v16a0)
		if err = ie.phy_ParametersCommon_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode phy_ParametersCommon_v16a0", err)
		}
	}
	return nil
}
