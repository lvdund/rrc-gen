package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictInfoSCG_powerCoordination_FR1 struct {
	p_maxNR_FR1 *P_Max `optional`
	p_maxEUTRA  *P_Max `optional`
	p_maxUE_FR1 *P_Max `optional`
}

func (ie *ConfigRestrictInfoSCG_powerCoordination_FR1) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.p_maxNR_FR1 != nil, ie.p_maxEUTRA != nil, ie.p_maxUE_FR1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.p_maxNR_FR1 != nil {
		if err = ie.p_maxNR_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode p_maxNR_FR1", err)
		}
	}
	if ie.p_maxEUTRA != nil {
		if err = ie.p_maxEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode p_maxEUTRA", err)
		}
	}
	if ie.p_maxUE_FR1 != nil {
		if err = ie.p_maxUE_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode p_maxUE_FR1", err)
		}
	}
	return nil
}

func (ie *ConfigRestrictInfoSCG_powerCoordination_FR1) Decode(r *uper.UperReader) error {
	var err error
	var p_maxNR_FR1Present bool
	if p_maxNR_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var p_maxEUTRAPresent bool
	if p_maxEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var p_maxUE_FR1Present bool
	if p_maxUE_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if p_maxNR_FR1Present {
		ie.p_maxNR_FR1 = new(P_Max)
		if err = ie.p_maxNR_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode p_maxNR_FR1", err)
		}
	}
	if p_maxEUTRAPresent {
		ie.p_maxEUTRA = new(P_Max)
		if err = ie.p_maxEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode p_maxEUTRA", err)
		}
	}
	if p_maxUE_FR1Present {
		ie.p_maxUE_FR1 = new(P_Max)
		if err = ie.p_maxUE_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode p_maxUE_FR1", err)
		}
	}
	return nil
}
