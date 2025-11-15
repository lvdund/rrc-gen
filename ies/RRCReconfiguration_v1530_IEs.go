package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_v1530_IEs struct {
	masterCellGroup                    *[]byte                                  `optional`
	fullConfig                         *RRCReconfiguration_v1530_IEs_fullConfig `optional`
	dedicatedNAS_MessageList           []DedicatedNAS_Message                   `lb:1,ub:maxDRB,optional`
	masterKeyUpdate                    *MasterKeyUpdate                         `optional`
	dedicatedSIB1_Delivery             *[]byte                                  `optional`
	dedicatedSystemInformationDelivery *[]byte                                  `optional`
	otherConfig                        *OtherConfig                             `optional`
	nonCriticalExtension               *RRCReconfiguration_v1540_IEs            `optional`
}

func (ie *RRCReconfiguration_v1530_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.masterCellGroup != nil, ie.fullConfig != nil, len(ie.dedicatedNAS_MessageList) > 0, ie.masterKeyUpdate != nil, ie.dedicatedSIB1_Delivery != nil, ie.dedicatedSystemInformationDelivery != nil, ie.otherConfig != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.masterCellGroup != nil {
		if err = w.WriteOctetString(*ie.masterCellGroup, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode masterCellGroup", err)
		}
	}
	if ie.fullConfig != nil {
		if err = ie.fullConfig.Encode(w); err != nil {
			return utils.WrapError("Encode fullConfig", err)
		}
	}
	if len(ie.dedicatedNAS_MessageList) > 0 {
		tmp_dedicatedNAS_MessageList := utils.NewSequence[*DedicatedNAS_Message]([]*DedicatedNAS_Message{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
		for _, i := range ie.dedicatedNAS_MessageList {
			tmp_dedicatedNAS_MessageList.Value = append(tmp_dedicatedNAS_MessageList.Value, &i)
		}
		if err = tmp_dedicatedNAS_MessageList.Encode(w); err != nil {
			return utils.WrapError("Encode dedicatedNAS_MessageList", err)
		}
	}
	if ie.masterKeyUpdate != nil {
		if err = ie.masterKeyUpdate.Encode(w); err != nil {
			return utils.WrapError("Encode masterKeyUpdate", err)
		}
	}
	if ie.dedicatedSIB1_Delivery != nil {
		if err = w.WriteOctetString(*ie.dedicatedSIB1_Delivery, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode dedicatedSIB1_Delivery", err)
		}
	}
	if ie.dedicatedSystemInformationDelivery != nil {
		if err = w.WriteOctetString(*ie.dedicatedSystemInformationDelivery, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode dedicatedSystemInformationDelivery", err)
		}
	}
	if ie.otherConfig != nil {
		if err = ie.otherConfig.Encode(w); err != nil {
			return utils.WrapError("Encode otherConfig", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_v1530_IEs) Decode(r *uper.UperReader) error {
	var err error
	var masterCellGroupPresent bool
	if masterCellGroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fullConfigPresent bool
	if fullConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dedicatedNAS_MessageListPresent bool
	if dedicatedNAS_MessageListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var masterKeyUpdatePresent bool
	if masterKeyUpdatePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dedicatedSIB1_DeliveryPresent bool
	if dedicatedSIB1_DeliveryPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dedicatedSystemInformationDeliveryPresent bool
	if dedicatedSystemInformationDeliveryPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var otherConfigPresent bool
	if otherConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if masterCellGroupPresent {
		var tmp_os_masterCellGroup []byte
		if tmp_os_masterCellGroup, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode masterCellGroup", err)
		}
		ie.masterCellGroup = &tmp_os_masterCellGroup
	}
	if fullConfigPresent {
		ie.fullConfig = new(RRCReconfiguration_v1530_IEs_fullConfig)
		if err = ie.fullConfig.Decode(r); err != nil {
			return utils.WrapError("Decode fullConfig", err)
		}
	}
	if dedicatedNAS_MessageListPresent {
		tmp_dedicatedNAS_MessageList := utils.NewSequence[*DedicatedNAS_Message]([]*DedicatedNAS_Message{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
		fn_dedicatedNAS_MessageList := func() *DedicatedNAS_Message {
			return new(DedicatedNAS_Message)
		}
		if err = tmp_dedicatedNAS_MessageList.Decode(r, fn_dedicatedNAS_MessageList); err != nil {
			return utils.WrapError("Decode dedicatedNAS_MessageList", err)
		}
		ie.dedicatedNAS_MessageList = []DedicatedNAS_Message{}
		for _, i := range tmp_dedicatedNAS_MessageList.Value {
			ie.dedicatedNAS_MessageList = append(ie.dedicatedNAS_MessageList, *i)
		}
	}
	if masterKeyUpdatePresent {
		ie.masterKeyUpdate = new(MasterKeyUpdate)
		if err = ie.masterKeyUpdate.Decode(r); err != nil {
			return utils.WrapError("Decode masterKeyUpdate", err)
		}
	}
	if dedicatedSIB1_DeliveryPresent {
		var tmp_os_dedicatedSIB1_Delivery []byte
		if tmp_os_dedicatedSIB1_Delivery, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode dedicatedSIB1_Delivery", err)
		}
		ie.dedicatedSIB1_Delivery = &tmp_os_dedicatedSIB1_Delivery
	}
	if dedicatedSystemInformationDeliveryPresent {
		var tmp_os_dedicatedSystemInformationDelivery []byte
		if tmp_os_dedicatedSystemInformationDelivery, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode dedicatedSystemInformationDelivery", err)
		}
		ie.dedicatedSystemInformationDelivery = &tmp_os_dedicatedSystemInformationDelivery
	}
	if otherConfigPresent {
		ie.otherConfig = new(OtherConfig)
		if err = ie.otherConfig.Decode(r); err != nil {
			return utils.WrapError("Decode otherConfig", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCReconfiguration_v1540_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
