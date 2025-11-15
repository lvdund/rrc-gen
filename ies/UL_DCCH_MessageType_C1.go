package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_DCCH_MessageType_C1_Choice_nothing uint64 = iota
	UL_DCCH_MessageType_C1_Choice_measurementReport
	UL_DCCH_MessageType_C1_Choice_rrcReconfigurationComplete
	UL_DCCH_MessageType_C1_Choice_rrcSetupComplete
	UL_DCCH_MessageType_C1_Choice_rrcReestablishmentComplete
	UL_DCCH_MessageType_C1_Choice_rrcResumeComplete
	UL_DCCH_MessageType_C1_Choice_securityModeComplete
	UL_DCCH_MessageType_C1_Choice_securityModeFailure
	UL_DCCH_MessageType_C1_Choice_ulInformationTransfer
	UL_DCCH_MessageType_C1_Choice_locationMeasurementIndication
	UL_DCCH_MessageType_C1_Choice_ueCapabilityInformation
	UL_DCCH_MessageType_C1_Choice_counterCheckResponse
	UL_DCCH_MessageType_C1_Choice_ueAssistanceInformation
	UL_DCCH_MessageType_C1_Choice_failureInformation
	UL_DCCH_MessageType_C1_Choice_ulInformationTransferMRDC
	UL_DCCH_MessageType_C1_Choice_scgFailureInformation
	UL_DCCH_MessageType_C1_Choice_scgFailureInformationEUTRA
)

type UL_DCCH_MessageType_C1 struct {
	Choice                        uint64
	measurementReport             *MeasurementReport
	rrcReconfigurationComplete    *RRCReconfigurationComplete
	rrcSetupComplete              *RRCSetupComplete
	rrcReestablishmentComplete    *RRCReestablishmentComplete
	rrcResumeComplete             *RRCResumeComplete
	securityModeComplete          *SecurityModeComplete
	securityModeFailure           *SecurityModeFailure
	ulInformationTransfer         *ULInformationTransfer
	locationMeasurementIndication *LocationMeasurementIndication
	ueCapabilityInformation       *UECapabilityInformation
	counterCheckResponse          *CounterCheckResponse
	ueAssistanceInformation       *UEAssistanceInformation
	failureInformation            *FailureInformation
	ulInformationTransferMRDC     *ULInformationTransferMRDC
	scgFailureInformation         *SCGFailureInformation
	scgFailureInformationEUTRA    *SCGFailureInformationEUTRA
}

func (ie *UL_DCCH_MessageType_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_DCCH_MessageType_C1_Choice_measurementReport:
		if err = ie.measurementReport.Encode(w); err != nil {
			err = utils.WrapError("Encode measurementReport", err)
		}
	case UL_DCCH_MessageType_C1_Choice_rrcReconfigurationComplete:
		if err = ie.rrcReconfigurationComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReconfigurationComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_rrcSetupComplete:
		if err = ie.rrcSetupComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcSetupComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_rrcReestablishmentComplete:
		if err = ie.rrcReestablishmentComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReestablishmentComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_rrcResumeComplete:
		if err = ie.rrcResumeComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcResumeComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_securityModeComplete:
		if err = ie.securityModeComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode securityModeComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_securityModeFailure:
		if err = ie.securityModeFailure.Encode(w); err != nil {
			err = utils.WrapError("Encode securityModeFailure", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ulInformationTransfer:
		if err = ie.ulInformationTransfer.Encode(w); err != nil {
			err = utils.WrapError("Encode ulInformationTransfer", err)
		}
	case UL_DCCH_MessageType_C1_Choice_locationMeasurementIndication:
		if err = ie.locationMeasurementIndication.Encode(w); err != nil {
			err = utils.WrapError("Encode locationMeasurementIndication", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ueCapabilityInformation:
		if err = ie.ueCapabilityInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode ueCapabilityInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_counterCheckResponse:
		if err = ie.counterCheckResponse.Encode(w); err != nil {
			err = utils.WrapError("Encode counterCheckResponse", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ueAssistanceInformation:
		if err = ie.ueAssistanceInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode ueAssistanceInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_failureInformation:
		if err = ie.failureInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode failureInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ulInformationTransferMRDC:
		if err = ie.ulInformationTransferMRDC.Encode(w); err != nil {
			err = utils.WrapError("Encode ulInformationTransferMRDC", err)
		}
	case UL_DCCH_MessageType_C1_Choice_scgFailureInformation:
		if err = ie.scgFailureInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode scgFailureInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_scgFailureInformationEUTRA:
		if err = ie.scgFailureInformationEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode scgFailureInformationEUTRA", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UL_DCCH_MessageType_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_DCCH_MessageType_C1_Choice_measurementReport:
		ie.measurementReport = new(MeasurementReport)
		if err = ie.measurementReport.Decode(r); err != nil {
			return utils.WrapError("Decode measurementReport", err)
		}
	case UL_DCCH_MessageType_C1_Choice_rrcReconfigurationComplete:
		ie.rrcReconfigurationComplete = new(RRCReconfigurationComplete)
		if err = ie.rrcReconfigurationComplete.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReconfigurationComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_rrcSetupComplete:
		ie.rrcSetupComplete = new(RRCSetupComplete)
		if err = ie.rrcSetupComplete.Decode(r); err != nil {
			return utils.WrapError("Decode rrcSetupComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_rrcReestablishmentComplete:
		ie.rrcReestablishmentComplete = new(RRCReestablishmentComplete)
		if err = ie.rrcReestablishmentComplete.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReestablishmentComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_rrcResumeComplete:
		ie.rrcResumeComplete = new(RRCResumeComplete)
		if err = ie.rrcResumeComplete.Decode(r); err != nil {
			return utils.WrapError("Decode rrcResumeComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_securityModeComplete:
		ie.securityModeComplete = new(SecurityModeComplete)
		if err = ie.securityModeComplete.Decode(r); err != nil {
			return utils.WrapError("Decode securityModeComplete", err)
		}
	case UL_DCCH_MessageType_C1_Choice_securityModeFailure:
		ie.securityModeFailure = new(SecurityModeFailure)
		if err = ie.securityModeFailure.Decode(r); err != nil {
			return utils.WrapError("Decode securityModeFailure", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ulInformationTransfer:
		ie.ulInformationTransfer = new(ULInformationTransfer)
		if err = ie.ulInformationTransfer.Decode(r); err != nil {
			return utils.WrapError("Decode ulInformationTransfer", err)
		}
	case UL_DCCH_MessageType_C1_Choice_locationMeasurementIndication:
		ie.locationMeasurementIndication = new(LocationMeasurementIndication)
		if err = ie.locationMeasurementIndication.Decode(r); err != nil {
			return utils.WrapError("Decode locationMeasurementIndication", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ueCapabilityInformation:
		ie.ueCapabilityInformation = new(UECapabilityInformation)
		if err = ie.ueCapabilityInformation.Decode(r); err != nil {
			return utils.WrapError("Decode ueCapabilityInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_counterCheckResponse:
		ie.counterCheckResponse = new(CounterCheckResponse)
		if err = ie.counterCheckResponse.Decode(r); err != nil {
			return utils.WrapError("Decode counterCheckResponse", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ueAssistanceInformation:
		ie.ueAssistanceInformation = new(UEAssistanceInformation)
		if err = ie.ueAssistanceInformation.Decode(r); err != nil {
			return utils.WrapError("Decode ueAssistanceInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_failureInformation:
		ie.failureInformation = new(FailureInformation)
		if err = ie.failureInformation.Decode(r); err != nil {
			return utils.WrapError("Decode failureInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_ulInformationTransferMRDC:
		ie.ulInformationTransferMRDC = new(ULInformationTransferMRDC)
		if err = ie.ulInformationTransferMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode ulInformationTransferMRDC", err)
		}
	case UL_DCCH_MessageType_C1_Choice_scgFailureInformation:
		ie.scgFailureInformation = new(SCGFailureInformation)
		if err = ie.scgFailureInformation.Decode(r); err != nil {
			return utils.WrapError("Decode scgFailureInformation", err)
		}
	case UL_DCCH_MessageType_C1_Choice_scgFailureInformationEUTRA:
		ie.scgFailureInformationEUTRA = new(SCGFailureInformationEUTRA)
		if err = ie.scgFailureInformationEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode scgFailureInformationEUTRA", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
