package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination struct {
	bandList                         []BandParameters                  `lb:1,ub:maxSimultaneousBands,madatory`
	featureSetCombination            FeatureSetCombinationId           `madatory`
	ca_ParametersEUTRA               *CA_ParametersEUTRA               `optional`
	ca_ParametersNR                  *CA_ParametersNR                  `optional`
	mrdc_Parameters                  *MRDC_Parameters                  `optional`
	supportedBandwidthCombinationSet *uper.BitString                   `lb:1,ub:32,optional`
	powerClass_v1530                 *BandCombination_powerClass_v1530 `optional`
}

func (ie *BandCombination) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersEUTRA != nil, ie.ca_ParametersNR != nil, ie.mrdc_Parameters != nil, ie.supportedBandwidthCombinationSet != nil, ie.powerClass_v1530 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_bandList := utils.NewSequence[*BandParameters]([]*BandParameters{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.bandList {
		tmp_bandList.Value = append(tmp_bandList.Value, &i)
	}
	if err = tmp_bandList.Encode(w); err != nil {
		return utils.WrapError("Encode bandList", err)
	}
	if err = ie.featureSetCombination.Encode(w); err != nil {
		return utils.WrapError("Encode featureSetCombination", err)
	}
	if ie.ca_ParametersEUTRA != nil {
		if err = ie.ca_ParametersEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersEUTRA", err)
		}
	}
	if ie.ca_ParametersNR != nil {
		if err = ie.ca_ParametersNR.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR", err)
		}
	}
	if ie.mrdc_Parameters != nil {
		if err = ie.mrdc_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode mrdc_Parameters", err)
		}
	}
	if ie.supportedBandwidthCombinationSet != nil {
		if err = w.WriteBitString(ie.supportedBandwidthCombinationSet.Bytes, uint(ie.supportedBandwidthCombinationSet.NumBits), &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode supportedBandwidthCombinationSet", err)
		}
	}
	if ie.powerClass_v1530 != nil {
		if err = ie.powerClass_v1530.Encode(w); err != nil {
			return utils.WrapError("Encode powerClass_v1530", err)
		}
	}
	return nil
}

func (ie *BandCombination) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersEUTRAPresent bool
	if ca_ParametersEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNRPresent bool
	if ca_ParametersNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mrdc_ParametersPresent bool
	if mrdc_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandwidthCombinationSetPresent bool
	if supportedBandwidthCombinationSetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var powerClass_v1530Present bool
	if powerClass_v1530Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_bandList := utils.NewSequence[*BandParameters]([]*BandParameters{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn_bandList := func() *BandParameters {
		return new(BandParameters)
	}
	if err = tmp_bandList.Decode(r, fn_bandList); err != nil {
		return utils.WrapError("Decode bandList", err)
	}
	ie.bandList = []BandParameters{}
	for _, i := range tmp_bandList.Value {
		ie.bandList = append(ie.bandList, *i)
	}
	if err = ie.featureSetCombination.Decode(r); err != nil {
		return utils.WrapError("Decode featureSetCombination", err)
	}
	if ca_ParametersEUTRAPresent {
		ie.ca_ParametersEUTRA = new(CA_ParametersEUTRA)
		if err = ie.ca_ParametersEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersEUTRA", err)
		}
	}
	if ca_ParametersNRPresent {
		ie.ca_ParametersNR = new(CA_ParametersNR)
		if err = ie.ca_ParametersNR.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR", err)
		}
	}
	if mrdc_ParametersPresent {
		ie.mrdc_Parameters = new(MRDC_Parameters)
		if err = ie.mrdc_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode mrdc_Parameters", err)
		}
	}
	if supportedBandwidthCombinationSetPresent {
		var tmp_bs_supportedBandwidthCombinationSet []byte
		var n_supportedBandwidthCombinationSet uint
		if tmp_bs_supportedBandwidthCombinationSet, n_supportedBandwidthCombinationSet, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode supportedBandwidthCombinationSet", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_supportedBandwidthCombinationSet,
			NumBits: uint64(n_supportedBandwidthCombinationSet),
		}
		ie.supportedBandwidthCombinationSet = &tmp_bitstring
	}
	if powerClass_v1530Present {
		ie.powerClass_v1530 = new(BandCombination_powerClass_v1530)
		if err = ie.powerClass_v1530.Decode(r); err != nil {
			return utils.WrapError("Decode powerClass_v1530", err)
		}
	}
	return nil
}
