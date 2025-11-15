package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v16a0 struct {
	phy_Parameters_v16a0 *Phy_Parameters_v16a0 `optional`
	rf_Parameters_v16a0  *RF_Parameters_v16a0  `optional`
	nonCriticalExtension interface{}           `optional`
}

func (ie *UE_NR_Capability_v16a0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.phy_Parameters_v16a0 != nil, ie.rf_Parameters_v16a0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.phy_Parameters_v16a0 != nil {
		if err = ie.phy_Parameters_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode phy_Parameters_v16a0", err)
		}
	}
	if ie.rf_Parameters_v16a0 != nil {
		if err = ie.rf_Parameters_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode rf_Parameters_v16a0", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v16a0) Decode(r *uper.UperReader) error {
	var err error
	var phy_Parameters_v16a0Present bool
	if phy_Parameters_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rf_Parameters_v16a0Present bool
	if rf_Parameters_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if phy_Parameters_v16a0Present {
		ie.phy_Parameters_v16a0 = new(Phy_Parameters_v16a0)
		if err = ie.phy_Parameters_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode phy_Parameters_v16a0", err)
		}
	}
	if rf_Parameters_v16a0Present {
		ie.rf_Parameters_v16a0 = new(RF_Parameters_v16a0)
		if err = ie.rf_Parameters_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode rf_Parameters_v16a0", err)
		}
	}
	return nil
}
