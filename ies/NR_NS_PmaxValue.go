package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NR_NS_PmaxValue struct {
	additionalPmax             *P_Max                     `optional`
	additionalSpectrumEmission AdditionalSpectrumEmission `madatory`
}

func (ie *NR_NS_PmaxValue) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.additionalPmax != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.additionalPmax != nil {
		if err = ie.additionalPmax.Encode(w); err != nil {
			return utils.WrapError("Encode additionalPmax", err)
		}
	}
	if err = ie.additionalSpectrumEmission.Encode(w); err != nil {
		return utils.WrapError("Encode additionalSpectrumEmission", err)
	}
	return nil
}

func (ie *NR_NS_PmaxValue) Decode(r *uper.UperReader) error {
	var err error
	var additionalPmaxPresent bool
	if additionalPmaxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if additionalPmaxPresent {
		ie.additionalPmax = new(P_Max)
		if err = ie.additionalPmax.Decode(r); err != nil {
			return utils.WrapError("Decode additionalPmax", err)
		}
	}
	if err = ie.additionalSpectrumEmission.Decode(r); err != nil {
		return utils.WrapError("Decode additionalSpectrumEmission", err)
	}
	return nil
}
