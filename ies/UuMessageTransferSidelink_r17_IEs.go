package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UuMessageTransferSidelink_r17_IEs struct {
	sl_PagingDelivery_r17            *[]byte     `optional`
	sl_SIB1_Delivery_r17             *[]byte     `optional`
	sl_SystemInformationDelivery_r17 *[]byte     `optional`
	lateNonCriticalExtension         *[]byte     `optional`
	nonCriticalExtension             interface{} `optional`
}

func (ie *UuMessageTransferSidelink_r17_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_PagingDelivery_r17 != nil, ie.sl_SIB1_Delivery_r17 != nil, ie.sl_SystemInformationDelivery_r17 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_PagingDelivery_r17 != nil {
		if err = w.WriteOctetString(*ie.sl_PagingDelivery_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sl_PagingDelivery_r17", err)
		}
	}
	if ie.sl_SIB1_Delivery_r17 != nil {
		if err = w.WriteOctetString(*ie.sl_SIB1_Delivery_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sl_SIB1_Delivery_r17", err)
		}
	}
	if ie.sl_SystemInformationDelivery_r17 != nil {
		if err = w.WriteOctetString(*ie.sl_SystemInformationDelivery_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sl_SystemInformationDelivery_r17", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UuMessageTransferSidelink_r17_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sl_PagingDelivery_r17Present bool
	if sl_PagingDelivery_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SIB1_Delivery_r17Present bool
	if sl_SIB1_Delivery_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SystemInformationDelivery_r17Present bool
	if sl_SystemInformationDelivery_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PagingDelivery_r17Present {
		var tmp_os_sl_PagingDelivery_r17 []byte
		if tmp_os_sl_PagingDelivery_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sl_PagingDelivery_r17", err)
		}
		ie.sl_PagingDelivery_r17 = &tmp_os_sl_PagingDelivery_r17
	}
	if sl_SIB1_Delivery_r17Present {
		var tmp_os_sl_SIB1_Delivery_r17 []byte
		if tmp_os_sl_SIB1_Delivery_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sl_SIB1_Delivery_r17", err)
		}
		ie.sl_SIB1_Delivery_r17 = &tmp_os_sl_SIB1_Delivery_r17
	}
	if sl_SystemInformationDelivery_r17Present {
		var tmp_os_sl_SystemInformationDelivery_r17 []byte
		if tmp_os_sl_SystemInformationDelivery_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sl_SystemInformationDelivery_r17", err)
		}
		ie.sl_SystemInformationDelivery_r17 = &tmp_os_sl_SystemInformationDelivery_r17
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
