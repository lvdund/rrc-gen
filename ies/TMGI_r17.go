package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TMGI_r17 struct {
	plmn_Id_r17   TMGI_r17_plmn_Id_r17 `lb:1,ub:maxPLMN,madatory`
	serviceId_r17 []byte               `lb:3,ub:3,madatory`
}

func (ie *TMGI_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.plmn_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_Id_r17", err)
	}
	if err = w.WriteOctetString(ie.serviceId_r17, &uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteOctetString serviceId_r17", err)
	}
	return nil
}

func (ie *TMGI_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.plmn_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_Id_r17", err)
	}
	var tmp_os_serviceId_r17 []byte
	if tmp_os_serviceId_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadOctetString serviceId_r17", err)
	}
	ie.serviceId_r17 = tmp_os_serviceId_r17
	return nil
}
