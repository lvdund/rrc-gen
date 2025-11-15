package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersFRX_Diff_r16 struct {
	directMCG_SCellActivation_r16       *MAC_ParametersFRX_Diff_r16_directMCG_SCellActivation_r16       `optional`
	directMCG_SCellActivationResume_r16 *MAC_ParametersFRX_Diff_r16_directMCG_SCellActivationResume_r16 `optional`
	directSCG_SCellActivation_r16       *MAC_ParametersFRX_Diff_r16_directSCG_SCellActivation_r16       `optional`
	directSCG_SCellActivationResume_r16 *MAC_ParametersFRX_Diff_r16_directSCG_SCellActivationResume_r16 `optional`
	drx_Adaptation_r16                  *MAC_ParametersFRX_Diff_r16_drx_Adaptation_r16                  `optional`
}

func (ie *MAC_ParametersFRX_Diff_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.directMCG_SCellActivation_r16 != nil, ie.directMCG_SCellActivationResume_r16 != nil, ie.directSCG_SCellActivation_r16 != nil, ie.directSCG_SCellActivationResume_r16 != nil, ie.drx_Adaptation_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.directMCG_SCellActivation_r16 != nil {
		if err = ie.directMCG_SCellActivation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode directMCG_SCellActivation_r16", err)
		}
	}
	if ie.directMCG_SCellActivationResume_r16 != nil {
		if err = ie.directMCG_SCellActivationResume_r16.Encode(w); err != nil {
			return utils.WrapError("Encode directMCG_SCellActivationResume_r16", err)
		}
	}
	if ie.directSCG_SCellActivation_r16 != nil {
		if err = ie.directSCG_SCellActivation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode directSCG_SCellActivation_r16", err)
		}
	}
	if ie.directSCG_SCellActivationResume_r16 != nil {
		if err = ie.directSCG_SCellActivationResume_r16.Encode(w); err != nil {
			return utils.WrapError("Encode directSCG_SCellActivationResume_r16", err)
		}
	}
	if ie.drx_Adaptation_r16 != nil {
		if err = ie.drx_Adaptation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode drx_Adaptation_r16", err)
		}
	}
	return nil
}

func (ie *MAC_ParametersFRX_Diff_r16) Decode(r *uper.UperReader) error {
	var err error
	var directMCG_SCellActivation_r16Present bool
	if directMCG_SCellActivation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var directMCG_SCellActivationResume_r16Present bool
	if directMCG_SCellActivationResume_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var directSCG_SCellActivation_r16Present bool
	if directSCG_SCellActivation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var directSCG_SCellActivationResume_r16Present bool
	if directSCG_SCellActivationResume_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var drx_Adaptation_r16Present bool
	if drx_Adaptation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if directMCG_SCellActivation_r16Present {
		ie.directMCG_SCellActivation_r16 = new(MAC_ParametersFRX_Diff_r16_directMCG_SCellActivation_r16)
		if err = ie.directMCG_SCellActivation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode directMCG_SCellActivation_r16", err)
		}
	}
	if directMCG_SCellActivationResume_r16Present {
		ie.directMCG_SCellActivationResume_r16 = new(MAC_ParametersFRX_Diff_r16_directMCG_SCellActivationResume_r16)
		if err = ie.directMCG_SCellActivationResume_r16.Decode(r); err != nil {
			return utils.WrapError("Decode directMCG_SCellActivationResume_r16", err)
		}
	}
	if directSCG_SCellActivation_r16Present {
		ie.directSCG_SCellActivation_r16 = new(MAC_ParametersFRX_Diff_r16_directSCG_SCellActivation_r16)
		if err = ie.directSCG_SCellActivation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode directSCG_SCellActivation_r16", err)
		}
	}
	if directSCG_SCellActivationResume_r16Present {
		ie.directSCG_SCellActivationResume_r16 = new(MAC_ParametersFRX_Diff_r16_directSCG_SCellActivationResume_r16)
		if err = ie.directSCG_SCellActivationResume_r16.Decode(r); err != nil {
			return utils.WrapError("Decode directSCG_SCellActivationResume_r16", err)
		}
	}
	if drx_Adaptation_r16Present {
		ie.drx_Adaptation_r16 = new(MAC_ParametersFRX_Diff_r16_drx_Adaptation_r16)
		if err = ie.drx_Adaptation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode drx_Adaptation_r16", err)
		}
	}
	return nil
}
