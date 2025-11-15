package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UERadioPagingInformation_IEs struct {
	supportedBandListNRForPaging []FreqBandIndicatorNR               `lb:1,ub:maxBands,optional`
	nonCriticalExtension         *UERadioPagingInformation_v15e0_IEs `optional`
}

func (ie *UERadioPagingInformation_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.supportedBandListNRForPaging) > 0, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.supportedBandListNRForPaging) > 0 {
		tmp_supportedBandListNRForPaging := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.supportedBandListNRForPaging {
			tmp_supportedBandListNRForPaging.Value = append(tmp_supportedBandListNRForPaging.Value, &i)
		}
		if err = tmp_supportedBandListNRForPaging.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandListNRForPaging", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UERadioPagingInformation_IEs) Decode(r *uper.UperReader) error {
	var err error
	var supportedBandListNRForPagingPresent bool
	if supportedBandListNRForPagingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedBandListNRForPagingPresent {
		tmp_supportedBandListNRForPaging := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_supportedBandListNRForPaging := func() *FreqBandIndicatorNR {
			return new(FreqBandIndicatorNR)
		}
		if err = tmp_supportedBandListNRForPaging.Decode(r, fn_supportedBandListNRForPaging); err != nil {
			return utils.WrapError("Decode supportedBandListNRForPaging", err)
		}
		ie.supportedBandListNRForPaging = []FreqBandIndicatorNR{}
		for _, i := range tmp_supportedBandListNRForPaging.Value {
			ie.supportedBandListNRForPaging = append(ie.supportedBandListNRForPaging, *i)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UERadioPagingInformation_v15e0_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
