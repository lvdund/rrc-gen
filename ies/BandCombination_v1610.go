package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1610 struct {
	bandList_v1610                []BandParameters_v1610                      `lb:1,ub:maxSimultaneousBands,optional`
	ca_ParametersNR_v1610         *CA_ParametersNR_v1610                      `optional`
	ca_ParametersNRDC_v1610       *CA_ParametersNRDC_v1610                    `optional`
	powerClass_v1610              *BandCombination_v1610_powerClass_v1610     `optional`
	powerClassNRPart_r16          *BandCombination_v1610_powerClassNRPart_r16 `optional`
	featureSetCombinationDAPS_r16 *FeatureSetCombinationId                    `optional`
	mrdc_Parameters_v1620         *MRDC_Parameters_v1620                      `optional`
}

func (ie *BandCombination_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.bandList_v1610) > 0, ie.ca_ParametersNR_v1610 != nil, ie.ca_ParametersNRDC_v1610 != nil, ie.powerClass_v1610 != nil, ie.powerClassNRPart_r16 != nil, ie.featureSetCombinationDAPS_r16 != nil, ie.mrdc_Parameters_v1620 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.bandList_v1610) > 0 {
		tmp_bandList_v1610 := utils.NewSequence[*BandParameters_v1610]([]*BandParameters_v1610{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		for _, i := range ie.bandList_v1610 {
			tmp_bandList_v1610.Value = append(tmp_bandList_v1610.Value, &i)
		}
		if err = tmp_bandList_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode bandList_v1610", err)
		}
	}
	if ie.ca_ParametersNR_v1610 != nil {
		if err = ie.ca_ParametersNR_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_v1610", err)
		}
	}
	if ie.ca_ParametersNRDC_v1610 != nil {
		if err = ie.ca_ParametersNRDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNRDC_v1610", err)
		}
	}
	if ie.powerClass_v1610 != nil {
		if err = ie.powerClass_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode powerClass_v1610", err)
		}
	}
	if ie.powerClassNRPart_r16 != nil {
		if err = ie.powerClassNRPart_r16.Encode(w); err != nil {
			return utils.WrapError("Encode powerClassNRPart_r16", err)
		}
	}
	if ie.featureSetCombinationDAPS_r16 != nil {
		if err = ie.featureSetCombinationDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode featureSetCombinationDAPS_r16", err)
		}
	}
	if ie.mrdc_Parameters_v1620 != nil {
		if err = ie.mrdc_Parameters_v1620.Encode(w); err != nil {
			return utils.WrapError("Encode mrdc_Parameters_v1620", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1610) Decode(r *uper.UperReader) error {
	var err error
	var bandList_v1610Present bool
	if bandList_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNR_v1610Present bool
	if ca_ParametersNR_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ca_ParametersNRDC_v1610Present bool
	if ca_ParametersNRDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var powerClass_v1610Present bool
	if powerClass_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var powerClassNRPart_r16Present bool
	if powerClassNRPart_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var featureSetCombinationDAPS_r16Present bool
	if featureSetCombinationDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mrdc_Parameters_v1620Present bool
	if mrdc_Parameters_v1620Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bandList_v1610Present {
		tmp_bandList_v1610 := utils.NewSequence[*BandParameters_v1610]([]*BandParameters_v1610{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		fn_bandList_v1610 := func() *BandParameters_v1610 {
			return new(BandParameters_v1610)
		}
		if err = tmp_bandList_v1610.Decode(r, fn_bandList_v1610); err != nil {
			return utils.WrapError("Decode bandList_v1610", err)
		}
		ie.bandList_v1610 = []BandParameters_v1610{}
		for _, i := range tmp_bandList_v1610.Value {
			ie.bandList_v1610 = append(ie.bandList_v1610, *i)
		}
	}
	if ca_ParametersNR_v1610Present {
		ie.ca_ParametersNR_v1610 = new(CA_ParametersNR_v1610)
		if err = ie.ca_ParametersNR_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_v1610", err)
		}
	}
	if ca_ParametersNRDC_v1610Present {
		ie.ca_ParametersNRDC_v1610 = new(CA_ParametersNRDC_v1610)
		if err = ie.ca_ParametersNRDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNRDC_v1610", err)
		}
	}
	if powerClass_v1610Present {
		ie.powerClass_v1610 = new(BandCombination_v1610_powerClass_v1610)
		if err = ie.powerClass_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode powerClass_v1610", err)
		}
	}
	if powerClassNRPart_r16Present {
		ie.powerClassNRPart_r16 = new(BandCombination_v1610_powerClassNRPart_r16)
		if err = ie.powerClassNRPart_r16.Decode(r); err != nil {
			return utils.WrapError("Decode powerClassNRPart_r16", err)
		}
	}
	if featureSetCombinationDAPS_r16Present {
		ie.featureSetCombinationDAPS_r16 = new(FeatureSetCombinationId)
		if err = ie.featureSetCombinationDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode featureSetCombinationDAPS_r16", err)
		}
	}
	if mrdc_Parameters_v1620Present {
		ie.mrdc_Parameters_v1620 = new(MRDC_Parameters_v1620)
		if err = ie.mrdc_Parameters_v1620.Decode(r); err != nil {
			return utils.WrapError("Decode mrdc_Parameters_v1620", err)
		}
	}
	return nil
}
