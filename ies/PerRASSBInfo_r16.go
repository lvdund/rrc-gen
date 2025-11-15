package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PerRASSBInfo_r16 struct {
	ssb_Index_r16                  SSB_Index                `madatory`
	numberOfPreamblesSentOnSSB_r16 int64                    `lb:1,ub:200,madatory`
	perRAAttemptInfoList_r16       PerRAAttemptInfoList_r16 `madatory`
}

func (ie *PerRASSBInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ssb_Index_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_Index_r16", err)
	}
	if err = w.WriteInteger(ie.numberOfPreamblesSentOnSSB_r16, &uper.Constraint{Lb: 1, Ub: 200}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfPreamblesSentOnSSB_r16", err)
	}
	if err = ie.perRAAttemptInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode perRAAttemptInfoList_r16", err)
	}
	return nil
}

func (ie *PerRASSBInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ssb_Index_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_Index_r16", err)
	}
	var tmp_int_numberOfPreamblesSentOnSSB_r16 int64
	if tmp_int_numberOfPreamblesSentOnSSB_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 200}, false); err != nil {
		return utils.WrapError("ReadInteger numberOfPreamblesSentOnSSB_r16", err)
	}
	ie.numberOfPreamblesSentOnSSB_r16 = tmp_int_numberOfPreamblesSentOnSSB_r16
	if err = ie.perRAAttemptInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode perRAAttemptInfoList_r16", err)
	}
	return nil
}
