package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapNCSG_InfoNR_r17 struct {
	intraFreq_needForNCSG_r17 NeedForNCSG_IntraFreqList_r17 `madatory`
	interFreq_needForNCSG_r17 NeedForNCSG_BandListNR_r17    `madatory`
}

func (ie *NeedForGapNCSG_InfoNR_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.intraFreq_needForNCSG_r17.Encode(w); err != nil {
		return utils.WrapError("Encode intraFreq_needForNCSG_r17", err)
	}
	if err = ie.interFreq_needForNCSG_r17.Encode(w); err != nil {
		return utils.WrapError("Encode interFreq_needForNCSG_r17", err)
	}
	return nil
}

func (ie *NeedForGapNCSG_InfoNR_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.intraFreq_needForNCSG_r17.Decode(r); err != nil {
		return utils.WrapError("Decode intraFreq_needForNCSG_r17", err)
	}
	if err = ie.interFreq_needForNCSG_r17.Decode(r); err != nil {
		return utils.WrapError("Decode interFreq_needForNCSG_r17", err)
	}
	return nil
}
