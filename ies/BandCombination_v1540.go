package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1540 struct {
	bandList_v1540        []BandParameters_v1540 `lb:1,ub:maxSimultaneousBands,madatory`
	ca_ParametersNR_v1540 *CA_ParametersNR_v1540 `optional`
}

func (ie *BandCombination_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_v1540 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_bandList_v1540 := utils.NewSequence[*BandParameters_v1540]([]*BandParameters_v1540{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.bandList_v1540 {
		tmp_bandList_v1540.Value = append(tmp_bandList_v1540.Value, &i)
	}
	if err = tmp_bandList_v1540.Encode(w); err != nil {
		return utils.WrapError("Encode bandList_v1540", err)
	}
	if ie.ca_ParametersNR_v1540 != nil {
		if err = ie.ca_ParametersNR_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_v1540", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1540) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_v1540Present bool
	if ca_ParametersNR_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_bandList_v1540 := utils.NewSequence[*BandParameters_v1540]([]*BandParameters_v1540{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn_bandList_v1540 := func() *BandParameters_v1540 {
		return new(BandParameters_v1540)
	}
	if err = tmp_bandList_v1540.Decode(r, fn_bandList_v1540); err != nil {
		return utils.WrapError("Decode bandList_v1540", err)
	}
	ie.bandList_v1540 = []BandParameters_v1540{}
	for _, i := range tmp_bandList_v1540.Value {
		ie.bandList_v1540 = append(ie.bandList_v1540, *i)
	}
	if ca_ParametersNR_v1540Present {
		ie.ca_ParametersNR_v1540 = new(CA_ParametersNR_v1540)
		if err = ie.ca_ParametersNR_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_v1540", err)
		}
	}
	return nil
}
