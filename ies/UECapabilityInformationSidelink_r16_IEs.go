package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UECapabilityInformationSidelink_r16_IEs struct {
	accessStratumReleaseSidelink_r16           AccessStratumReleaseSidelink_r16           `madatory`
	pdcp_ParametersSidelink_r16                *PDCP_ParametersSidelink_r16               `optional`
	rlc_ParametersSidelink_r16                 *RLC_ParametersSidelink_r16                `optional`
	supportedBandCombinationListSidelinkNR_r16 *BandCombinationListSidelinkNR_r16         `optional`
	supportedBandListSidelink_r16              []BandSidelinkPC5_r16                      `lb:1,ub:maxBands,optional`
	appliedFreqBandListFilter_r16              *FreqBandList                              `optional`
	lateNonCriticalExtension                   *[]byte                                    `optional`
	nonCriticalExtension                       *UECapabilityInformationSidelink_v1700_IEs `optional`
}

func (ie *UECapabilityInformationSidelink_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdcp_ParametersSidelink_r16 != nil, ie.rlc_ParametersSidelink_r16 != nil, ie.supportedBandCombinationListSidelinkNR_r16 != nil, len(ie.supportedBandListSidelink_r16) > 0, ie.appliedFreqBandListFilter_r16 != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.accessStratumReleaseSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode accessStratumReleaseSidelink_r16", err)
	}
	if ie.pdcp_ParametersSidelink_r16 != nil {
		if err = ie.pdcp_ParametersSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_ParametersSidelink_r16", err)
		}
	}
	if ie.rlc_ParametersSidelink_r16 != nil {
		if err = ie.rlc_ParametersSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rlc_ParametersSidelink_r16", err)
		}
	}
	if ie.supportedBandCombinationListSidelinkNR_r16 != nil {
		if err = ie.supportedBandCombinationListSidelinkNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandCombinationListSidelinkNR_r16", err)
		}
	}
	if len(ie.supportedBandListSidelink_r16) > 0 {
		tmp_supportedBandListSidelink_r16 := utils.NewSequence[*BandSidelinkPC5_r16]([]*BandSidelinkPC5_r16{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		for _, i := range ie.supportedBandListSidelink_r16 {
			tmp_supportedBandListSidelink_r16.Value = append(tmp_supportedBandListSidelink_r16.Value, &i)
		}
		if err = tmp_supportedBandListSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportedBandListSidelink_r16", err)
		}
	}
	if ie.appliedFreqBandListFilter_r16 != nil {
		if err = ie.appliedFreqBandListFilter_r16.Encode(w); err != nil {
			return utils.WrapError("Encode appliedFreqBandListFilter_r16", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UECapabilityInformationSidelink_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var pdcp_ParametersSidelink_r16Present bool
	if pdcp_ParametersSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rlc_ParametersSidelink_r16Present bool
	if rlc_ParametersSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandCombinationListSidelinkNR_r16Present bool
	if supportedBandCombinationListSidelinkNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedBandListSidelink_r16Present bool
	if supportedBandListSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var appliedFreqBandListFilter_r16Present bool
	if appliedFreqBandListFilter_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.accessStratumReleaseSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode accessStratumReleaseSidelink_r16", err)
	}
	if pdcp_ParametersSidelink_r16Present {
		ie.pdcp_ParametersSidelink_r16 = new(PDCP_ParametersSidelink_r16)
		if err = ie.pdcp_ParametersSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_ParametersSidelink_r16", err)
		}
	}
	if rlc_ParametersSidelink_r16Present {
		ie.rlc_ParametersSidelink_r16 = new(RLC_ParametersSidelink_r16)
		if err = ie.rlc_ParametersSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rlc_ParametersSidelink_r16", err)
		}
	}
	if supportedBandCombinationListSidelinkNR_r16Present {
		ie.supportedBandCombinationListSidelinkNR_r16 = new(BandCombinationListSidelinkNR_r16)
		if err = ie.supportedBandCombinationListSidelinkNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportedBandCombinationListSidelinkNR_r16", err)
		}
	}
	if supportedBandListSidelink_r16Present {
		tmp_supportedBandListSidelink_r16 := utils.NewSequence[*BandSidelinkPC5_r16]([]*BandSidelinkPC5_r16{}, uper.Constraint{Lb: 1, Ub: maxBands}, false)
		fn_supportedBandListSidelink_r16 := func() *BandSidelinkPC5_r16 {
			return new(BandSidelinkPC5_r16)
		}
		if err = tmp_supportedBandListSidelink_r16.Decode(r, fn_supportedBandListSidelink_r16); err != nil {
			return utils.WrapError("Decode supportedBandListSidelink_r16", err)
		}
		ie.supportedBandListSidelink_r16 = []BandSidelinkPC5_r16{}
		for _, i := range tmp_supportedBandListSidelink_r16.Value {
			ie.supportedBandListSidelink_r16 = append(ie.supportedBandListSidelink_r16, *i)
		}
	}
	if appliedFreqBandListFilter_r16Present {
		ie.appliedFreqBandListFilter_r16 = new(FreqBandList)
		if err = ie.appliedFreqBandListFilter_r16.Decode(r); err != nil {
			return utils.WrapError("Decode appliedFreqBandListFilter_r16", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UECapabilityInformationSidelink_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
