package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapsInfoNR_r16 struct {
	intraFreq_needForGap_r16 NeedForGapsIntraFreqList_r16 `madatory`
	interFreq_needForGap_r16 NeedForGapsBandListNR_r16    `madatory`
}

func (ie *NeedForGapsInfoNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.intraFreq_needForGap_r16.Encode(w); err != nil {
		return utils.WrapError("Encode intraFreq_needForGap_r16", err)
	}
	if err = ie.interFreq_needForGap_r16.Encode(w); err != nil {
		return utils.WrapError("Encode interFreq_needForGap_r16", err)
	}
	return nil
}

func (ie *NeedForGapsInfoNR_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.intraFreq_needForGap_r16.Decode(r); err != nil {
		return utils.WrapError("Decode intraFreq_needForGap_r16", err)
	}
	if err = ie.interFreq_needForGap_r16.Decode(r); err != nil {
		return utils.WrapError("Decode interFreq_needForGap_r16", err)
	}
	return nil
}
