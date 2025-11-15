package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1610_interFreqDAPS_r16 struct {
	interFreqAsyncDAPS_r16                        *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqAsyncDAPS_r16                        `optional`
	interFreqDiffSCS_DAPS_r16                     *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqDiffSCS_DAPS_r16                     `optional`
	interFreqMultiUL_TransmissionDAPS_r16         *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqMultiUL_TransmissionDAPS_r16         `optional`
	interFreqSemiStaticPowerSharingDAPS_Mode1_r16 *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqSemiStaticPowerSharingDAPS_Mode1_r16 `optional`
	interFreqSemiStaticPowerSharingDAPS_Mode2_r16 *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqSemiStaticPowerSharingDAPS_Mode2_r16 `optional`
	interFreqDynamicPowerSharingDAPS_r16          *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqDynamicPowerSharingDAPS_r16          `optional`
	interFreqUL_TransCancellationDAPS_r16         *CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqUL_TransCancellationDAPS_r16         `optional`
}

func (ie *CA_ParametersNR_v1610_interFreqDAPS_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.interFreqAsyncDAPS_r16 != nil, ie.interFreqDiffSCS_DAPS_r16 != nil, ie.interFreqMultiUL_TransmissionDAPS_r16 != nil, ie.interFreqSemiStaticPowerSharingDAPS_Mode1_r16 != nil, ie.interFreqSemiStaticPowerSharingDAPS_Mode2_r16 != nil, ie.interFreqDynamicPowerSharingDAPS_r16 != nil, ie.interFreqUL_TransCancellationDAPS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.interFreqAsyncDAPS_r16 != nil {
		if err = ie.interFreqAsyncDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqAsyncDAPS_r16", err)
		}
	}
	if ie.interFreqDiffSCS_DAPS_r16 != nil {
		if err = ie.interFreqDiffSCS_DAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqDiffSCS_DAPS_r16", err)
		}
	}
	if ie.interFreqMultiUL_TransmissionDAPS_r16 != nil {
		if err = ie.interFreqMultiUL_TransmissionDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqMultiUL_TransmissionDAPS_r16", err)
		}
	}
	if ie.interFreqSemiStaticPowerSharingDAPS_Mode1_r16 != nil {
		if err = ie.interFreqSemiStaticPowerSharingDAPS_Mode1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqSemiStaticPowerSharingDAPS_Mode1_r16", err)
		}
	}
	if ie.interFreqSemiStaticPowerSharingDAPS_Mode2_r16 != nil {
		if err = ie.interFreqSemiStaticPowerSharingDAPS_Mode2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqSemiStaticPowerSharingDAPS_Mode2_r16", err)
		}
	}
	if ie.interFreqDynamicPowerSharingDAPS_r16 != nil {
		if err = ie.interFreqDynamicPowerSharingDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqDynamicPowerSharingDAPS_r16", err)
		}
	}
	if ie.interFreqUL_TransCancellationDAPS_r16 != nil {
		if err = ie.interFreqUL_TransCancellationDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqUL_TransCancellationDAPS_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1610_interFreqDAPS_r16) Decode(r *uper.UperReader) error {
	var err error
	var interFreqAsyncDAPS_r16Present bool
	if interFreqAsyncDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqDiffSCS_DAPS_r16Present bool
	if interFreqDiffSCS_DAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqMultiUL_TransmissionDAPS_r16Present bool
	if interFreqMultiUL_TransmissionDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqSemiStaticPowerSharingDAPS_Mode1_r16Present bool
	if interFreqSemiStaticPowerSharingDAPS_Mode1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqSemiStaticPowerSharingDAPS_Mode2_r16Present bool
	if interFreqSemiStaticPowerSharingDAPS_Mode2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqDynamicPowerSharingDAPS_r16Present bool
	if interFreqDynamicPowerSharingDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqUL_TransCancellationDAPS_r16Present bool
	if interFreqUL_TransCancellationDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if interFreqAsyncDAPS_r16Present {
		ie.interFreqAsyncDAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqAsyncDAPS_r16)
		if err = ie.interFreqAsyncDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqAsyncDAPS_r16", err)
		}
	}
	if interFreqDiffSCS_DAPS_r16Present {
		ie.interFreqDiffSCS_DAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqDiffSCS_DAPS_r16)
		if err = ie.interFreqDiffSCS_DAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqDiffSCS_DAPS_r16", err)
		}
	}
	if interFreqMultiUL_TransmissionDAPS_r16Present {
		ie.interFreqMultiUL_TransmissionDAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqMultiUL_TransmissionDAPS_r16)
		if err = ie.interFreqMultiUL_TransmissionDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqMultiUL_TransmissionDAPS_r16", err)
		}
	}
	if interFreqSemiStaticPowerSharingDAPS_Mode1_r16Present {
		ie.interFreqSemiStaticPowerSharingDAPS_Mode1_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqSemiStaticPowerSharingDAPS_Mode1_r16)
		if err = ie.interFreqSemiStaticPowerSharingDAPS_Mode1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqSemiStaticPowerSharingDAPS_Mode1_r16", err)
		}
	}
	if interFreqSemiStaticPowerSharingDAPS_Mode2_r16Present {
		ie.interFreqSemiStaticPowerSharingDAPS_Mode2_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqSemiStaticPowerSharingDAPS_Mode2_r16)
		if err = ie.interFreqSemiStaticPowerSharingDAPS_Mode2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqSemiStaticPowerSharingDAPS_Mode2_r16", err)
		}
	}
	if interFreqDynamicPowerSharingDAPS_r16Present {
		ie.interFreqDynamicPowerSharingDAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqDynamicPowerSharingDAPS_r16)
		if err = ie.interFreqDynamicPowerSharingDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqDynamicPowerSharingDAPS_r16", err)
		}
	}
	if interFreqUL_TransCancellationDAPS_r16Present {
		ie.interFreqUL_TransCancellationDAPS_r16 = new(CA_ParametersNR_v1610_interFreqDAPS_r16_interFreqUL_TransCancellationDAPS_r16)
		if err = ie.interFreqUL_TransCancellationDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interFreqUL_TransCancellationDAPS_r16", err)
		}
	}
	return nil
}
