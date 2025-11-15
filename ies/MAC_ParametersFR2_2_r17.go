package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersFR2_2_r17 struct {
	directMCG_SCellActivation_r17       *MAC_ParametersFR2_2_r17_directMCG_SCellActivation_r17       `optional`
	directMCG_SCellActivationResume_r17 *MAC_ParametersFR2_2_r17_directMCG_SCellActivationResume_r17 `optional`
	directSCG_SCellActivation_r17       *MAC_ParametersFR2_2_r17_directSCG_SCellActivation_r17       `optional`
	directSCG_SCellActivationResume_r17 *MAC_ParametersFR2_2_r17_directSCG_SCellActivationResume_r17 `optional`
	drx_Adaptation_r17                  *MAC_ParametersFR2_2_r17_drx_Adaptation_r17                  `optional`
}

func (ie *MAC_ParametersFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.directMCG_SCellActivation_r17 != nil, ie.directMCG_SCellActivationResume_r17 != nil, ie.directSCG_SCellActivation_r17 != nil, ie.directSCG_SCellActivationResume_r17 != nil, ie.drx_Adaptation_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.directMCG_SCellActivation_r17 != nil {
		if err = ie.directMCG_SCellActivation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode directMCG_SCellActivation_r17", err)
		}
	}
	if ie.directMCG_SCellActivationResume_r17 != nil {
		if err = ie.directMCG_SCellActivationResume_r17.Encode(w); err != nil {
			return utils.WrapError("Encode directMCG_SCellActivationResume_r17", err)
		}
	}
	if ie.directSCG_SCellActivation_r17 != nil {
		if err = ie.directSCG_SCellActivation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode directSCG_SCellActivation_r17", err)
		}
	}
	if ie.directSCG_SCellActivationResume_r17 != nil {
		if err = ie.directSCG_SCellActivationResume_r17.Encode(w); err != nil {
			return utils.WrapError("Encode directSCG_SCellActivationResume_r17", err)
		}
	}
	if ie.drx_Adaptation_r17 != nil {
		if err = ie.drx_Adaptation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode drx_Adaptation_r17", err)
		}
	}
	return nil
}

func (ie *MAC_ParametersFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var directMCG_SCellActivation_r17Present bool
	if directMCG_SCellActivation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var directMCG_SCellActivationResume_r17Present bool
	if directMCG_SCellActivationResume_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var directSCG_SCellActivation_r17Present bool
	if directSCG_SCellActivation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var directSCG_SCellActivationResume_r17Present bool
	if directSCG_SCellActivationResume_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var drx_Adaptation_r17Present bool
	if drx_Adaptation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if directMCG_SCellActivation_r17Present {
		ie.directMCG_SCellActivation_r17 = new(MAC_ParametersFR2_2_r17_directMCG_SCellActivation_r17)
		if err = ie.directMCG_SCellActivation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode directMCG_SCellActivation_r17", err)
		}
	}
	if directMCG_SCellActivationResume_r17Present {
		ie.directMCG_SCellActivationResume_r17 = new(MAC_ParametersFR2_2_r17_directMCG_SCellActivationResume_r17)
		if err = ie.directMCG_SCellActivationResume_r17.Decode(r); err != nil {
			return utils.WrapError("Decode directMCG_SCellActivationResume_r17", err)
		}
	}
	if directSCG_SCellActivation_r17Present {
		ie.directSCG_SCellActivation_r17 = new(MAC_ParametersFR2_2_r17_directSCG_SCellActivation_r17)
		if err = ie.directSCG_SCellActivation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode directSCG_SCellActivation_r17", err)
		}
	}
	if directSCG_SCellActivationResume_r17Present {
		ie.directSCG_SCellActivationResume_r17 = new(MAC_ParametersFR2_2_r17_directSCG_SCellActivationResume_r17)
		if err = ie.directSCG_SCellActivationResume_r17.Decode(r); err != nil {
			return utils.WrapError("Decode directSCG_SCellActivationResume_r17", err)
		}
	}
	if drx_Adaptation_r17Present {
		ie.drx_Adaptation_r17 = new(MAC_ParametersFR2_2_r17_drx_Adaptation_r17)
		if err = ie.drx_Adaptation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode drx_Adaptation_r17", err)
		}
	}
	return nil
}
