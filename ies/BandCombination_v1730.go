package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1730 struct {
	ca_ParametersNR_v1730   *CA_ParametersNR_v1730   `optional`
	ca_ParametersNRDC_v1730 *CA_ParametersNRDC_v1730 `optional`
	bandList_v1730          []BandParameters_v1730   `lb:1,ub:maxSimultaneousBands,optional`
}

func (ie *BandCombination_v1730) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_v1730 != nil, ie.ca_ParametersNRDC_v1730 != nil, len(ie.bandList_v1730) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNR_v1730 != nil {
		if err = ie.ca_ParametersNR_v1730.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_v1730", err)
		}
	}
	if ie.ca_ParametersNRDC_v1730 != nil {
		if err = ie.ca_ParametersNRDC_v1730.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNRDC_v1730", err)
		}
	}
	if len(ie.bandList_v1730) > 0 {
		tmp_bandList_v1730 := utils.NewSequence[*BandParameters_v1730]([]*BandParameters_v1730{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		for _, i := range ie.bandList_v1730 {
			tmp_bandList_v1730.Value = append(tmp_bandList_v1730.Value, &i)
		}
		if err = tmp_bandList_v1730.Encode(w); err != nil {
			return utils.WrapError("Encode bandList_v1730", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1730) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_v1730Present bool
	if ca_ParametersNR_v1730Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNRDC_v1730Present bool
	if ca_ParametersNRDC_v1730Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bandList_v1730Present bool
	if bandList_v1730Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNR_v1730Present {
		ie.ca_ParametersNR_v1730 = new(CA_ParametersNR_v1730)
		if err = ie.ca_ParametersNR_v1730.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_v1730", err)
		}
	}
	if ca_ParametersNRDC_v1730Present {
		ie.ca_ParametersNRDC_v1730 = new(CA_ParametersNRDC_v1730)
		if err = ie.ca_ParametersNRDC_v1730.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNRDC_v1730", err)
		}
	}
	if bandList_v1730Present {
		tmp_bandList_v1730 := utils.NewSequence[*BandParameters_v1730]([]*BandParameters_v1730{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		fn_bandList_v1730 := func() *BandParameters_v1730 {
			return new(BandParameters_v1730)
		}
		if err = tmp_bandList_v1730.Decode(r, fn_bandList_v1730); err != nil {
			return utils.WrapError("Decode bandList_v1730", err)
		}
		ie.bandList_v1730 = []BandParameters_v1730{}
		for _, i := range tmp_bandList_v1730.Value {
			ie.bandList_v1730 = append(ie.bandList_v1730, *i)
		}
	}
	return nil
}
