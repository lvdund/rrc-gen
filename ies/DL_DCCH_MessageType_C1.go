package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_DCCH_MessageType_C1_Choice_nothing uint64 = iota
	DL_DCCH_MessageType_C1_Choice_rrcReconfiguration
	DL_DCCH_MessageType_C1_Choice_rrcResume
	DL_DCCH_MessageType_C1_Choice_rrcRelease
	DL_DCCH_MessageType_C1_Choice_rrcReestablishment
	DL_DCCH_MessageType_C1_Choice_securityModeCommand
	DL_DCCH_MessageType_C1_Choice_dlInformationTransfer
	DL_DCCH_MessageType_C1_Choice_ueCapabilityEnquiry
	DL_DCCH_MessageType_C1_Choice_counterCheck
	DL_DCCH_MessageType_C1_Choice_mobilityFromNRCommand
	DL_DCCH_MessageType_C1_Choice_dlDedicatedMessageSegment_r16
	DL_DCCH_MessageType_C1_Choice_ueInformationRequest_r16
	DL_DCCH_MessageType_C1_Choice_dlInformationTransferMRDC_r16
	DL_DCCH_MessageType_C1_Choice_loggedMeasurementConfiguration_r16
	DL_DCCH_MessageType_C1_Choice_spare3
	DL_DCCH_MessageType_C1_Choice_spare2
	DL_DCCH_MessageType_C1_Choice_spare1
)

type DL_DCCH_MessageType_C1 struct {
	Choice                             uint64
	rrcReconfiguration                 *RRCReconfiguration
	rrcResume                          *RRCResume
	rrcRelease                         *RRCRelease
	rrcReestablishment                 *RRCReestablishment
	securityModeCommand                *SecurityModeCommand
	dlInformationTransfer              *DLInformationTransfer
	ueCapabilityEnquiry                *UECapabilityEnquiry
	counterCheck                       *CounterCheck
	mobilityFromNRCommand              *MobilityFromNRCommand
	dlDedicatedMessageSegment_r16      *DLDedicatedMessageSegment_r16
	ueInformationRequest_r16           *UEInformationRequest_r16
	dlInformationTransferMRDC_r16      *DLInformationTransferMRDC_r16
	loggedMeasurementConfiguration_r16 *LoggedMeasurementConfiguration_r16
	spare3                             uper.NULL `madatory`
	spare2                             uper.NULL `madatory`
	spare1                             uper.NULL `madatory`
}

func (ie *DL_DCCH_MessageType_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_DCCH_MessageType_C1_Choice_rrcReconfiguration:
		if err = ie.rrcReconfiguration.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReconfiguration", err)
		}
	case DL_DCCH_MessageType_C1_Choice_rrcResume:
		if err = ie.rrcResume.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcResume", err)
		}
	case DL_DCCH_MessageType_C1_Choice_rrcRelease:
		if err = ie.rrcRelease.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcRelease", err)
		}
	case DL_DCCH_MessageType_C1_Choice_rrcReestablishment:
		if err = ie.rrcReestablishment.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReestablishment", err)
		}
	case DL_DCCH_MessageType_C1_Choice_securityModeCommand:
		if err = ie.securityModeCommand.Encode(w); err != nil {
			err = utils.WrapError("Encode securityModeCommand", err)
		}
	case DL_DCCH_MessageType_C1_Choice_dlInformationTransfer:
		if err = ie.dlInformationTransfer.Encode(w); err != nil {
			err = utils.WrapError("Encode dlInformationTransfer", err)
		}
	case DL_DCCH_MessageType_C1_Choice_ueCapabilityEnquiry:
		if err = ie.ueCapabilityEnquiry.Encode(w); err != nil {
			err = utils.WrapError("Encode ueCapabilityEnquiry", err)
		}
	case DL_DCCH_MessageType_C1_Choice_counterCheck:
		if err = ie.counterCheck.Encode(w); err != nil {
			err = utils.WrapError("Encode counterCheck", err)
		}
	case DL_DCCH_MessageType_C1_Choice_mobilityFromNRCommand:
		if err = ie.mobilityFromNRCommand.Encode(w); err != nil {
			err = utils.WrapError("Encode mobilityFromNRCommand", err)
		}
	case DL_DCCH_MessageType_C1_Choice_dlDedicatedMessageSegment_r16:
		if err = ie.dlDedicatedMessageSegment_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode dlDedicatedMessageSegment_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_ueInformationRequest_r16:
		if err = ie.ueInformationRequest_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ueInformationRequest_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_dlInformationTransferMRDC_r16:
		if err = ie.dlInformationTransferMRDC_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode dlInformationTransferMRDC_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_loggedMeasurementConfiguration_r16:
		if err = ie.loggedMeasurementConfiguration_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode loggedMeasurementConfiguration_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare3", err)
		}
	case DL_DCCH_MessageType_C1_Choice_spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare2", err)
		}
	case DL_DCCH_MessageType_C1_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DL_DCCH_MessageType_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(16, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_DCCH_MessageType_C1_Choice_rrcReconfiguration:
		ie.rrcReconfiguration = new(RRCReconfiguration)
		if err = ie.rrcReconfiguration.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReconfiguration", err)
		}
	case DL_DCCH_MessageType_C1_Choice_rrcResume:
		ie.rrcResume = new(RRCResume)
		if err = ie.rrcResume.Decode(r); err != nil {
			return utils.WrapError("Decode rrcResume", err)
		}
	case DL_DCCH_MessageType_C1_Choice_rrcRelease:
		ie.rrcRelease = new(RRCRelease)
		if err = ie.rrcRelease.Decode(r); err != nil {
			return utils.WrapError("Decode rrcRelease", err)
		}
	case DL_DCCH_MessageType_C1_Choice_rrcReestablishment:
		ie.rrcReestablishment = new(RRCReestablishment)
		if err = ie.rrcReestablishment.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReestablishment", err)
		}
	case DL_DCCH_MessageType_C1_Choice_securityModeCommand:
		ie.securityModeCommand = new(SecurityModeCommand)
		if err = ie.securityModeCommand.Decode(r); err != nil {
			return utils.WrapError("Decode securityModeCommand", err)
		}
	case DL_DCCH_MessageType_C1_Choice_dlInformationTransfer:
		ie.dlInformationTransfer = new(DLInformationTransfer)
		if err = ie.dlInformationTransfer.Decode(r); err != nil {
			return utils.WrapError("Decode dlInformationTransfer", err)
		}
	case DL_DCCH_MessageType_C1_Choice_ueCapabilityEnquiry:
		ie.ueCapabilityEnquiry = new(UECapabilityEnquiry)
		if err = ie.ueCapabilityEnquiry.Decode(r); err != nil {
			return utils.WrapError("Decode ueCapabilityEnquiry", err)
		}
	case DL_DCCH_MessageType_C1_Choice_counterCheck:
		ie.counterCheck = new(CounterCheck)
		if err = ie.counterCheck.Decode(r); err != nil {
			return utils.WrapError("Decode counterCheck", err)
		}
	case DL_DCCH_MessageType_C1_Choice_mobilityFromNRCommand:
		ie.mobilityFromNRCommand = new(MobilityFromNRCommand)
		if err = ie.mobilityFromNRCommand.Decode(r); err != nil {
			return utils.WrapError("Decode mobilityFromNRCommand", err)
		}
	case DL_DCCH_MessageType_C1_Choice_dlDedicatedMessageSegment_r16:
		ie.dlDedicatedMessageSegment_r16 = new(DLDedicatedMessageSegment_r16)
		if err = ie.dlDedicatedMessageSegment_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dlDedicatedMessageSegment_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_ueInformationRequest_r16:
		ie.ueInformationRequest_r16 = new(UEInformationRequest_r16)
		if err = ie.ueInformationRequest_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ueInformationRequest_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_dlInformationTransferMRDC_r16:
		ie.dlInformationTransferMRDC_r16 = new(DLInformationTransferMRDC_r16)
		if err = ie.dlInformationTransferMRDC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dlInformationTransferMRDC_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_loggedMeasurementConfiguration_r16:
		ie.loggedMeasurementConfiguration_r16 = new(LoggedMeasurementConfiguration_r16)
		if err = ie.loggedMeasurementConfiguration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode loggedMeasurementConfiguration_r16", err)
		}
	case DL_DCCH_MessageType_C1_Choice_spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare3", err)
		}
	case DL_DCCH_MessageType_C1_Choice_spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare2", err)
		}
	case DL_DCCH_MessageType_C1_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
