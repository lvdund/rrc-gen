package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_nothing uint64 = iota
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_ulDedicatedMessageSegment_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_dedicatedSIBRequest_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_mcgFailureInformation_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_ueInformationResponse_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_sidelinkUEInformationNR_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_ulInformationTransferIRAT_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_iabOtherInformation_r16
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_mbsInterestIndication_r17
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_uePositioningAssistanceInfo_r17
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_measurementReportAppLayer_r17
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare6
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare5
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare4
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare3
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare2
	UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare1
)

type UL_DCCH_MessageType_MessageClassExtension_C2 struct {
	Choice                          uint64
	ulDedicatedMessageSegment_r16   *ULDedicatedMessageSegment_r16
	dedicatedSIBRequest_r16         *DedicatedSIBRequest_r16
	mcgFailureInformation_r16       *MCGFailureInformation_r16
	ueInformationResponse_r16       *UEInformationResponse_r16
	sidelinkUEInformationNR_r16     *SidelinkUEInformationNR_r16
	ulInformationTransferIRAT_r16   *ULInformationTransferIRAT_r16
	iabOtherInformation_r16         *IABOtherInformation_r16
	mbsInterestIndication_r17       *MBSInterestIndication_r17
	uePositioningAssistanceInfo_r17 *UEPositioningAssistanceInfo_r17
	measurementReportAppLayer_r17   *MeasurementReportAppLayer_r17
	spare6                          uper.NULL `madatory`
	spare5                          uper.NULL `madatory`
	spare4                          uper.NULL `madatory`
	spare3                          uper.NULL `madatory`
	spare2                          uper.NULL `madatory`
	spare1                          uper.NULL `madatory`
}

func (ie *UL_DCCH_MessageType_MessageClassExtension_C2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_ulDedicatedMessageSegment_r16:
		if err = ie.ulDedicatedMessageSegment_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ulDedicatedMessageSegment_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_dedicatedSIBRequest_r16:
		if err = ie.dedicatedSIBRequest_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode dedicatedSIBRequest_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_mcgFailureInformation_r16:
		if err = ie.mcgFailureInformation_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode mcgFailureInformation_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_ueInformationResponse_r16:
		if err = ie.ueInformationResponse_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ueInformationResponse_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_sidelinkUEInformationNR_r16:
		if err = ie.sidelinkUEInformationNR_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode sidelinkUEInformationNR_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_ulInformationTransferIRAT_r16:
		if err = ie.ulInformationTransferIRAT_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ulInformationTransferIRAT_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_iabOtherInformation_r16:
		if err = ie.iabOtherInformation_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode iabOtherInformation_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_mbsInterestIndication_r17:
		if err = ie.mbsInterestIndication_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode mbsInterestIndication_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_uePositioningAssistanceInfo_r17:
		if err = ie.uePositioningAssistanceInfo_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode uePositioningAssistanceInfo_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_measurementReportAppLayer_r17:
		if err = ie.measurementReportAppLayer_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode measurementReportAppLayer_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare6:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare6", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare5:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare5", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare4:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare4", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare3", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare2", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UL_DCCH_MessageType_MessageClassExtension_C2) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_ulDedicatedMessageSegment_r16:
		ie.ulDedicatedMessageSegment_r16 = new(ULDedicatedMessageSegment_r16)
		if err = ie.ulDedicatedMessageSegment_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ulDedicatedMessageSegment_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_dedicatedSIBRequest_r16:
		ie.dedicatedSIBRequest_r16 = new(DedicatedSIBRequest_r16)
		if err = ie.dedicatedSIBRequest_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dedicatedSIBRequest_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_mcgFailureInformation_r16:
		ie.mcgFailureInformation_r16 = new(MCGFailureInformation_r16)
		if err = ie.mcgFailureInformation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mcgFailureInformation_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_ueInformationResponse_r16:
		ie.ueInformationResponse_r16 = new(UEInformationResponse_r16)
		if err = ie.ueInformationResponse_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ueInformationResponse_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_sidelinkUEInformationNR_r16:
		ie.sidelinkUEInformationNR_r16 = new(SidelinkUEInformationNR_r16)
		if err = ie.sidelinkUEInformationNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sidelinkUEInformationNR_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_ulInformationTransferIRAT_r16:
		ie.ulInformationTransferIRAT_r16 = new(ULInformationTransferIRAT_r16)
		if err = ie.ulInformationTransferIRAT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ulInformationTransferIRAT_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_iabOtherInformation_r16:
		ie.iabOtherInformation_r16 = new(IABOtherInformation_r16)
		if err = ie.iabOtherInformation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode iabOtherInformation_r16", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_mbsInterestIndication_r17:
		ie.mbsInterestIndication_r17 = new(MBSInterestIndication_r17)
		if err = ie.mbsInterestIndication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mbsInterestIndication_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_uePositioningAssistanceInfo_r17:
		ie.uePositioningAssistanceInfo_r17 = new(UEPositioningAssistanceInfo_r17)
		if err = ie.uePositioningAssistanceInfo_r17.Decode(r); err != nil {
			return utils.WrapError("Decode uePositioningAssistanceInfo_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_measurementReportAppLayer_r17:
		ie.measurementReportAppLayer_r17 = new(MeasurementReportAppLayer_r17)
		if err = ie.measurementReportAppLayer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode measurementReportAppLayer_r17", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare6:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare6", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare5:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare5", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare4:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare4", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare3", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare2", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_C2_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
