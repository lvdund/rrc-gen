package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_FSAI_InterFreq_r17 struct {
	dl_CarrierFreq_r17 ARFCN_ValueNR     `madatory`
	mbs_FSAI_List_r17  MBS_FSAI_List_r17 `madatory`
}

func (ie *MBS_FSAI_InterFreq_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.dl_CarrierFreq_r17.Encode(w); err != nil {
		return utils.WrapError("Encode dl_CarrierFreq_r17", err)
	}
	if err = ie.mbs_FSAI_List_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mbs_FSAI_List_r17", err)
	}
	return nil
}

func (ie *MBS_FSAI_InterFreq_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.dl_CarrierFreq_r17.Decode(r); err != nil {
		return utils.WrapError("Decode dl_CarrierFreq_r17", err)
	}
	if err = ie.mbs_FSAI_List_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mbs_FSAI_List_r17", err)
	}
	return nil
}
