package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictInfoSCG_powerCoordination_FR2_r16 struct {
	p_maxNR_FR2_MCG_r16 *P_Max `optional`
	p_maxNR_FR2_SCG_r16 *P_Max `optional`
	p_maxUE_FR2_r16     *P_Max `optional`
}

func (ie *ConfigRestrictInfoSCG_powerCoordination_FR2_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.p_maxNR_FR2_MCG_r16 != nil, ie.p_maxNR_FR2_SCG_r16 != nil, ie.p_maxUE_FR2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.p_maxNR_FR2_MCG_r16 != nil {
		if err = ie.p_maxNR_FR2_MCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode p_maxNR_FR2_MCG_r16", err)
		}
	}
	if ie.p_maxNR_FR2_SCG_r16 != nil {
		if err = ie.p_maxNR_FR2_SCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode p_maxNR_FR2_SCG_r16", err)
		}
	}
	if ie.p_maxUE_FR2_r16 != nil {
		if err = ie.p_maxUE_FR2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode p_maxUE_FR2_r16", err)
		}
	}
	return nil
}

func (ie *ConfigRestrictInfoSCG_powerCoordination_FR2_r16) Decode(r *uper.UperReader) error {
	var err error
	var p_maxNR_FR2_MCG_r16Present bool
	if p_maxNR_FR2_MCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var p_maxNR_FR2_SCG_r16Present bool
	if p_maxNR_FR2_SCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var p_maxUE_FR2_r16Present bool
	if p_maxUE_FR2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if p_maxNR_FR2_MCG_r16Present {
		ie.p_maxNR_FR2_MCG_r16 = new(P_Max)
		if err = ie.p_maxNR_FR2_MCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode p_maxNR_FR2_MCG_r16", err)
		}
	}
	if p_maxNR_FR2_SCG_r16Present {
		ie.p_maxNR_FR2_SCG_r16 = new(P_Max)
		if err = ie.p_maxNR_FR2_SCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode p_maxNR_FR2_SCG_r16", err)
		}
	}
	if p_maxUE_FR2_r16Present {
		ie.p_maxUE_FR2_r16 = new(P_Max)
		if err = ie.p_maxUE_FR2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode p_maxUE_FR2_r16", err)
		}
	}
	return nil
}
