package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultLogging2NR_r16 struct {
	carrierFreq_r16             ARFCN_ValueNR               `madatory`
	measResultListLoggingNR_r16 MeasResultListLoggingNR_r16 `madatory`
}

func (ie *MeasResultLogging2NR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.carrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq_r16", err)
	}
	if err = ie.measResultListLoggingNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResultListLoggingNR_r16", err)
	}
	return nil
}

func (ie *MeasResultLogging2NR_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.carrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq_r16", err)
	}
	if err = ie.measResultListLoggingNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode measResultListLoggingNR_r16", err)
	}
	return nil
}
