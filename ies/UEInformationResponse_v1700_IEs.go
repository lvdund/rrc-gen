package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEInformationResponse_v1700_IEs struct {
	successHO_Report_r17      *SuccessHO_Report_r17      `optional`
	connEstFailReportList_r17 *ConnEstFailReportList_r17 `optional`
	coarseLocationInfo_r17    *[]byte                    `optional`
	nonCriticalExtension      interface{}                `optional`
}

func (ie *UEInformationResponse_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.successHO_Report_r17 != nil, ie.connEstFailReportList_r17 != nil, ie.coarseLocationInfo_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.successHO_Report_r17 != nil {
		if err = ie.successHO_Report_r17.Encode(w); err != nil {
			return utils.WrapError("Encode successHO_Report_r17", err)
		}
	}
	if ie.connEstFailReportList_r17 != nil {
		if err = ie.connEstFailReportList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode connEstFailReportList_r17", err)
		}
	}
	if ie.coarseLocationInfo_r17 != nil {
		if err = w.WriteOctetString(*ie.coarseLocationInfo_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode coarseLocationInfo_r17", err)
		}
	}
	return nil
}

func (ie *UEInformationResponse_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var successHO_Report_r17Present bool
	if successHO_Report_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var connEstFailReportList_r17Present bool
	if connEstFailReportList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var coarseLocationInfo_r17Present bool
	if coarseLocationInfo_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if successHO_Report_r17Present {
		ie.successHO_Report_r17 = new(SuccessHO_Report_r17)
		if err = ie.successHO_Report_r17.Decode(r); err != nil {
			return utils.WrapError("Decode successHO_Report_r17", err)
		}
	}
	if connEstFailReportList_r17Present {
		ie.connEstFailReportList_r17 = new(ConnEstFailReportList_r17)
		if err = ie.connEstFailReportList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode connEstFailReportList_r17", err)
		}
	}
	if coarseLocationInfo_r17Present {
		var tmp_os_coarseLocationInfo_r17 []byte
		if tmp_os_coarseLocationInfo_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode coarseLocationInfo_r17", err)
		}
		ie.coarseLocationInfo_r17 = &tmp_os_coarseLocationInfo_r17
	}
	return nil
}
