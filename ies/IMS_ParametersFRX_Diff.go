package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IMS_ParametersFRX_Diff struct {
	voiceOverNR *IMS_ParametersFRX_Diff_voiceOverNR `optional`
}

func (ie *IMS_ParametersFRX_Diff) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.voiceOverNR != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.voiceOverNR != nil {
		if err = ie.voiceOverNR.Encode(w); err != nil {
			return utils.WrapError("Encode voiceOverNR", err)
		}
	}
	return nil
}

func (ie *IMS_ParametersFRX_Diff) Decode(r *uper.UperReader) error {
	var err error
	var voiceOverNRPresent bool
	if voiceOverNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if voiceOverNRPresent {
		ie.voiceOverNR = new(IMS_ParametersFRX_Diff_voiceOverNR)
		if err = ie.voiceOverNR.Decode(r); err != nil {
			return utils.WrapError("Decode voiceOverNR", err)
		}
	}
	return nil
}
