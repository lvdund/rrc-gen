package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCCH_MessageType_C1_Choice_nothing uint64 = iota
	SCCH_MessageType_C1_Choice_measurementReportSidelink
	SCCH_MessageType_C1_Choice_rrcReconfigurationSidelink
	SCCH_MessageType_C1_Choice_rrcReconfigurationCompleteSidelink
	SCCH_MessageType_C1_Choice_rrcReconfigurationFailureSidelink
	SCCH_MessageType_C1_Choice_ueCapabilityEnquirySidelink
	SCCH_MessageType_C1_Choice_ueCapabilityInformationSidelink
	SCCH_MessageType_C1_Choice_uuMessageTransferSidelink_r17
	SCCH_MessageType_C1_Choice_remoteUEInformationSidelink_r17
)

type SCCH_MessageType_C1 struct {
	Choice                             uint64
	measurementReportSidelink          *MeasurementReportSidelink
	rrcReconfigurationSidelink         *RRCReconfigurationSidelink
	rrcReconfigurationCompleteSidelink *RRCReconfigurationCompleteSidelink
	rrcReconfigurationFailureSidelink  *RRCReconfigurationFailureSidelink
	ueCapabilityEnquirySidelink        *UECapabilityEnquirySidelink
	ueCapabilityInformationSidelink    *UECapabilityInformationSidelink
	uuMessageTransferSidelink_r17      *UuMessageTransferSidelink_r17
	remoteUEInformationSidelink_r17    *RemoteUEInformationSidelink_r17
}

func (ie *SCCH_MessageType_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_C1_Choice_measurementReportSidelink:
		if err = ie.measurementReportSidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode measurementReportSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_rrcReconfigurationSidelink:
		if err = ie.rrcReconfigurationSidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReconfigurationSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_rrcReconfigurationCompleteSidelink:
		if err = ie.rrcReconfigurationCompleteSidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReconfigurationCompleteSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_rrcReconfigurationFailureSidelink:
		if err = ie.rrcReconfigurationFailureSidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReconfigurationFailureSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_ueCapabilityEnquirySidelink:
		if err = ie.ueCapabilityEnquirySidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode ueCapabilityEnquirySidelink", err)
		}
	case SCCH_MessageType_C1_Choice_ueCapabilityInformationSidelink:
		if err = ie.ueCapabilityInformationSidelink.Encode(w); err != nil {
			err = utils.WrapError("Encode ueCapabilityInformationSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_uuMessageTransferSidelink_r17:
		if err = ie.uuMessageTransferSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode uuMessageTransferSidelink_r17", err)
		}
	case SCCH_MessageType_C1_Choice_remoteUEInformationSidelink_r17:
		if err = ie.remoteUEInformationSidelink_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode remoteUEInformationSidelink_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SCCH_MessageType_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_C1_Choice_measurementReportSidelink:
		ie.measurementReportSidelink = new(MeasurementReportSidelink)
		if err = ie.measurementReportSidelink.Decode(r); err != nil {
			return utils.WrapError("Decode measurementReportSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_rrcReconfigurationSidelink:
		ie.rrcReconfigurationSidelink = new(RRCReconfigurationSidelink)
		if err = ie.rrcReconfigurationSidelink.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReconfigurationSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_rrcReconfigurationCompleteSidelink:
		ie.rrcReconfigurationCompleteSidelink = new(RRCReconfigurationCompleteSidelink)
		if err = ie.rrcReconfigurationCompleteSidelink.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReconfigurationCompleteSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_rrcReconfigurationFailureSidelink:
		ie.rrcReconfigurationFailureSidelink = new(RRCReconfigurationFailureSidelink)
		if err = ie.rrcReconfigurationFailureSidelink.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReconfigurationFailureSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_ueCapabilityEnquirySidelink:
		ie.ueCapabilityEnquirySidelink = new(UECapabilityEnquirySidelink)
		if err = ie.ueCapabilityEnquirySidelink.Decode(r); err != nil {
			return utils.WrapError("Decode ueCapabilityEnquirySidelink", err)
		}
	case SCCH_MessageType_C1_Choice_ueCapabilityInformationSidelink:
		ie.ueCapabilityInformationSidelink = new(UECapabilityInformationSidelink)
		if err = ie.ueCapabilityInformationSidelink.Decode(r); err != nil {
			return utils.WrapError("Decode ueCapabilityInformationSidelink", err)
		}
	case SCCH_MessageType_C1_Choice_uuMessageTransferSidelink_r17:
		ie.uuMessageTransferSidelink_r17 = new(UuMessageTransferSidelink_r17)
		if err = ie.uuMessageTransferSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode uuMessageTransferSidelink_r17", err)
		}
	case SCCH_MessageType_C1_Choice_remoteUEInformationSidelink_r17:
		ie.remoteUEInformationSidelink_r17 = new(RemoteUEInformationSidelink_r17)
		if err = ie.remoteUEInformationSidelink_r17.Decode(r); err != nil {
			return utils.WrapError("Decode remoteUEInformationSidelink_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
