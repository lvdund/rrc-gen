package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FrequencyConfig_NR_r16 struct {
	freqBandIndicatorNR_r16  FreqBandIndicatorNR `madatory`
	carrierCenterFreq_NR_r16 ARFCN_ValueNR       `madatory`
	carrierBandwidth_NR_r16  int64               `lb:1,ub:maxNrofPhysicalResourceBlocks,madatory`
	subcarrierSpacing_NR_r16 SubcarrierSpacing   `madatory`
}

func (ie *FrequencyConfig_NR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.freqBandIndicatorNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode freqBandIndicatorNR_r16", err)
	}
	if err = ie.carrierCenterFreq_NR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierCenterFreq_NR_r16", err)
	}
	if err = w.WriteInteger(ie.carrierBandwidth_NR_r16, &uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks}, false); err != nil {
		return utils.WrapError("WriteInteger carrierBandwidth_NR_r16", err)
	}
	if err = ie.subcarrierSpacing_NR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode subcarrierSpacing_NR_r16", err)
	}
	return nil
}

func (ie *FrequencyConfig_NR_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.freqBandIndicatorNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode freqBandIndicatorNR_r16", err)
	}
	if err = ie.carrierCenterFreq_NR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierCenterFreq_NR_r16", err)
	}
	var tmp_int_carrierBandwidth_NR_r16 int64
	if tmp_int_carrierBandwidth_NR_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofPhysicalResourceBlocks}, false); err != nil {
		return utils.WrapError("ReadInteger carrierBandwidth_NR_r16", err)
	}
	ie.carrierBandwidth_NR_r16 = tmp_int_carrierBandwidth_NR_r16
	if err = ie.subcarrierSpacing_NR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode subcarrierSpacing_NR_r16", err)
	}
	return nil
}
