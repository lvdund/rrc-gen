package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResult2EUTRA_r16 struct {
	carrierFreq_r16    ARFCN_ValueEUTRA    `madatory`
	measResultList_r16 MeasResultListEUTRA `madatory`
}

func (ie *MeasResult2EUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.carrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq_r16", err)
	}
	if err = ie.measResultList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResultList_r16", err)
	}
	return nil
}

func (ie *MeasResult2EUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.carrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq_r16", err)
	}
	if err = ie.measResultList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode measResultList_r16", err)
	}
	return nil
}
