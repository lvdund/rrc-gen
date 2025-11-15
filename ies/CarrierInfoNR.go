package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CarrierInfoNR struct {
	carrierFreq          ARFCN_ValueNR     `madatory`
	ssbSubcarrierSpacing SubcarrierSpacing `madatory`
	smtc                 *SSB_MTC          `optional`
}

func (ie *CarrierInfoNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.smtc != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.carrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq", err)
	}
	if err = ie.ssbSubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode ssbSubcarrierSpacing", err)
	}
	if ie.smtc != nil {
		if err = ie.smtc.Encode(w); err != nil {
			return utils.WrapError("Encode smtc", err)
		}
	}
	return nil
}

func (ie *CarrierInfoNR) Decode(r *uper.UperReader) error {
	var err error
	var smtcPresent bool
	if smtcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.carrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq", err)
	}
	if err = ie.ssbSubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode ssbSubcarrierSpacing", err)
	}
	if smtcPresent {
		ie.smtc = new(SSB_MTC)
		if err = ie.smtc.Decode(r); err != nil {
			return utils.WrapError("Decode smtc", err)
		}
	}
	return nil
}
