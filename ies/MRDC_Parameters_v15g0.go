package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v15g0 struct {
	simultaneousRxTxInterBandENDCPerBandPair *SimultaneousRxTxPerBandPair `optional`
}

func (ie *MRDC_Parameters_v15g0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.simultaneousRxTxInterBandENDCPerBandPair != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.simultaneousRxTxInterBandENDCPerBandPair != nil {
		if err = ie.simultaneousRxTxInterBandENDCPerBandPair.Encode(w); err != nil {
			return utils.WrapError("Encode simultaneousRxTxInterBandENDCPerBandPair", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v15g0) Decode(r *uper.UperReader) error {
	var err error
	var simultaneousRxTxInterBandENDCPerBandPairPresent bool
	if simultaneousRxTxInterBandENDCPerBandPairPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if simultaneousRxTxInterBandENDCPerBandPairPresent {
		ie.simultaneousRxTxInterBandENDCPerBandPair = new(SimultaneousRxTxPerBandPair)
		if err = ie.simultaneousRxTxInterBandENDCPerBandPair.Decode(r); err != nil {
			return utils.WrapError("Decode simultaneousRxTxInterBandENDCPerBandPair", err)
		}
	}
	return nil
}
