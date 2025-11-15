package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v1700 struct {
	bandCombination_v1700                     *BandCombination_v1700                  `optional`
	supportedBandPairListNR_v1700             []ULTxSwitchingBandPair_v1700           `lb:1,ub:maxULTxSwitchingBandPairs,optional`
	uplinkTxSwitchingBandParametersList_v1700 []UplinkTxSwitchingBandParameters_v1700 `lb:1,ub:maxSimultaneousBands,optional`
}

func (ie *BandCombination_UplinkTxSwitch_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bandCombination_v1700 != nil, len(ie.supportedBandPairListNR_v1700) > 0, len(ie.uplinkTxSwitchingBandParametersList_v1700) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bandCombination_v1700 != nil {
		if err = ie.bandCombination_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode bandCombination_v1700", err)
		}
	}
	if len(ie.supportedBandPairListNR_v1700) > 0 {
		tmp_supportedBandPairListNR_v1700 := utils.NewSequence[*ULTxSwitchingBandPair_v1700]([]*ULTxSwitchingBandPair_v1700{}, uper.Constraint{Lb: 1, Ub: maxULTxSwitchingBandPairs}, false)
		for _, i := range ie.supportedBandPairListNR_v1700 {
			tmp_supportedBandPairListNR_v1700.Value = append(tmp_supportedBandPairListNR_v1700.Value, &i)
		}
		if err = tmp_supportedBandPairListNR_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandPairListNR_v1700", err)
		}
	}
	if len(ie.uplinkTxSwitchingBandParametersList_v1700) > 0 {
		tmp_uplinkTxSwitchingBandParametersList_v1700 := utils.NewSequence[*UplinkTxSwitchingBandParameters_v1700]([]*UplinkTxSwitchingBandParameters_v1700{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		for _, i := range ie.uplinkTxSwitchingBandParametersList_v1700 {
			tmp_uplinkTxSwitchingBandParametersList_v1700.Value = append(tmp_uplinkTxSwitchingBandParametersList_v1700.Value, &i)
		}
		if err = tmp_uplinkTxSwitchingBandParametersList_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkTxSwitchingBandParametersList_v1700", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1700) Decode(r *uper.UperReader) error {
	var err error
	var bandCombination_v1700Present bool
	if bandCombination_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandPairListNR_v1700Present bool
	if supportedBandPairListNR_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkTxSwitchingBandParametersList_v1700Present bool
	if uplinkTxSwitchingBandParametersList_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bandCombination_v1700Present {
		ie.bandCombination_v1700 = new(BandCombination_v1700)
		if err = ie.bandCombination_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode bandCombination_v1700", err)
		}
	}
	if supportedBandPairListNR_v1700Present {
		tmp_supportedBandPairListNR_v1700 := utils.NewSequence[*ULTxSwitchingBandPair_v1700]([]*ULTxSwitchingBandPair_v1700{}, uper.Constraint{Lb: 1, Ub: maxULTxSwitchingBandPairs}, false)
		fn_supportedBandPairListNR_v1700 := func() *ULTxSwitchingBandPair_v1700 {
			return new(ULTxSwitchingBandPair_v1700)
		}
		if err = tmp_supportedBandPairListNR_v1700.Decode(r, fn_supportedBandPairListNR_v1700); err != nil {
			return utils.WrapError("Decode supportedBandPairListNR_v1700", err)
		}
		ie.supportedBandPairListNR_v1700 = []ULTxSwitchingBandPair_v1700{}
		for _, i := range tmp_supportedBandPairListNR_v1700.Value {
			ie.supportedBandPairListNR_v1700 = append(ie.supportedBandPairListNR_v1700, *i)
		}
	}
	if uplinkTxSwitchingBandParametersList_v1700Present {
		tmp_uplinkTxSwitchingBandParametersList_v1700 := utils.NewSequence[*UplinkTxSwitchingBandParameters_v1700]([]*UplinkTxSwitchingBandParameters_v1700{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		fn_uplinkTxSwitchingBandParametersList_v1700 := func() *UplinkTxSwitchingBandParameters_v1700 {
			return new(UplinkTxSwitchingBandParameters_v1700)
		}
		if err = tmp_uplinkTxSwitchingBandParametersList_v1700.Decode(r, fn_uplinkTxSwitchingBandParametersList_v1700); err != nil {
			return utils.WrapError("Decode uplinkTxSwitchingBandParametersList_v1700", err)
		}
		ie.uplinkTxSwitchingBandParametersList_v1700 = []UplinkTxSwitchingBandParameters_v1700{}
		for _, i := range tmp_uplinkTxSwitchingBandParametersList_v1700.Value {
			ie.uplinkTxSwitchingBandParametersList_v1700 = append(ie.uplinkTxSwitchingBandParametersList_v1700, *i)
		}
	}
	return nil
}
