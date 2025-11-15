package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v15g0 struct {
	simultaneousRxTxInterBandCAPerBandPair *SimultaneousRxTxPerBandPair `optional`
	simultaneousRxTxSULPerBandPair         *SimultaneousRxTxPerBandPair `optional`
}

func (ie *CA_ParametersNR_v15g0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.simultaneousRxTxInterBandCAPerBandPair != nil, ie.simultaneousRxTxSULPerBandPair != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.simultaneousRxTxInterBandCAPerBandPair != nil {
		if err = ie.simultaneousRxTxInterBandCAPerBandPair.Encode(w); err != nil {
			return utils.WrapError("Encode simultaneousRxTxInterBandCAPerBandPair", err)
		}
	}
	if ie.simultaneousRxTxSULPerBandPair != nil {
		if err = ie.simultaneousRxTxSULPerBandPair.Encode(w); err != nil {
			return utils.WrapError("Encode simultaneousRxTxSULPerBandPair", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v15g0) Decode(r *uper.UperReader) error {
	var err error
	var simultaneousRxTxInterBandCAPerBandPairPresent bool
	if simultaneousRxTxInterBandCAPerBandPairPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var simultaneousRxTxSULPerBandPairPresent bool
	if simultaneousRxTxSULPerBandPairPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if simultaneousRxTxInterBandCAPerBandPairPresent {
		ie.simultaneousRxTxInterBandCAPerBandPair = new(SimultaneousRxTxPerBandPair)
		if err = ie.simultaneousRxTxInterBandCAPerBandPair.Decode(r); err != nil {
			return utils.WrapError("Decode simultaneousRxTxInterBandCAPerBandPair", err)
		}
	}
	if simultaneousRxTxSULPerBandPairPresent {
		ie.simultaneousRxTxSULPerBandPair = new(SimultaneousRxTxPerBandPair)
		if err = ie.simultaneousRxTxSULPerBandPair.Decode(r); err != nil {
			return utils.WrapError("Decode simultaneousRxTxSULPerBandPair", err)
		}
	}
	return nil
}
