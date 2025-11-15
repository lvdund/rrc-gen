package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_NS_PmaxValue struct {
	additionalPmax             *int64 `lb:-30,ub:33,optional`
	additionalSpectrumEmission *int64 `lb:1,ub:288,optional`
}

func (ie *EUTRA_NS_PmaxValue) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.additionalPmax != nil, ie.additionalSpectrumEmission != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.additionalPmax != nil {
		if err = w.WriteInteger(*ie.additionalPmax, &uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
			return utils.WrapError("Encode additionalPmax", err)
		}
	}
	if ie.additionalSpectrumEmission != nil {
		if err = w.WriteInteger(*ie.additionalSpectrumEmission, &uper.Constraint{Lb: 1, Ub: 288}, false); err != nil {
			return utils.WrapError("Encode additionalSpectrumEmission", err)
		}
	}
	return nil
}

func (ie *EUTRA_NS_PmaxValue) Decode(r *uper.UperReader) error {
	var err error
	var additionalPmaxPresent bool
	if additionalPmaxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var additionalSpectrumEmissionPresent bool
	if additionalSpectrumEmissionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if additionalPmaxPresent {
		var tmp_int_additionalPmax int64
		if tmp_int_additionalPmax, err = r.ReadInteger(&uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
			return utils.WrapError("Decode additionalPmax", err)
		}
		ie.additionalPmax = &tmp_int_additionalPmax
	}
	if additionalSpectrumEmissionPresent {
		var tmp_int_additionalSpectrumEmission int64
		if tmp_int_additionalSpectrumEmission, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 288}, false); err != nil {
			return utils.WrapError("Decode additionalSpectrumEmission", err)
		}
		ie.additionalSpectrumEmission = &tmp_int_additionalSpectrumEmission
	}
	return nil
}
