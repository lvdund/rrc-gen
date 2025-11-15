package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IMS_ParametersFR2_2_r17 struct {
	voiceOverNR_r17 *IMS_ParametersFR2_2_r17_voiceOverNR_r17 `optional`
}

func (ie *IMS_ParametersFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.voiceOverNR_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.voiceOverNR_r17 != nil {
		if err = ie.voiceOverNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode voiceOverNR_r17", err)
		}
	}
	return nil
}

func (ie *IMS_ParametersFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var voiceOverNR_r17Present bool
	if voiceOverNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if voiceOverNR_r17Present {
		ie.voiceOverNR_r17 = new(IMS_ParametersFR2_2_r17_voiceOverNR_r17)
		if err = ie.voiceOverNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode voiceOverNR_r17", err)
		}
	}
	return nil
}
