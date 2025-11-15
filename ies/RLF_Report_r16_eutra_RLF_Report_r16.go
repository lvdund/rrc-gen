package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLF_Report_r16_eutra_RLF_Report_r16 struct {
	failedPCellId_EUTRA               CGI_InfoEUTRALogging `madatory`
	measResult_RLF_Report_EUTRA_r16   []byte               `madatory`
	measResult_RLF_Report_EUTRA_v1690 *[]byte              `optional`
}

func (ie *RLF_Report_r16_eutra_RLF_Report_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResult_RLF_Report_EUTRA_v1690 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.failedPCellId_EUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode failedPCellId_EUTRA", err)
	}
	if err = w.WriteOctetString(ie.measResult_RLF_Report_EUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString measResult_RLF_Report_EUTRA_r16", err)
	}
	if ie.measResult_RLF_Report_EUTRA_v1690 != nil {
		if err = w.WriteOctetString(*ie.measResult_RLF_Report_EUTRA_v1690, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode measResult_RLF_Report_EUTRA_v1690", err)
		}
	}
	return nil
}

func (ie *RLF_Report_r16_eutra_RLF_Report_r16) Decode(r *uper.UperReader) error {
	var err error
	var measResult_RLF_Report_EUTRA_v1690Present bool
	if measResult_RLF_Report_EUTRA_v1690Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.failedPCellId_EUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode failedPCellId_EUTRA", err)
	}
	var tmp_os_measResult_RLF_Report_EUTRA_r16 []byte
	if tmp_os_measResult_RLF_Report_EUTRA_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString measResult_RLF_Report_EUTRA_r16", err)
	}
	ie.measResult_RLF_Report_EUTRA_r16 = tmp_os_measResult_RLF_Report_EUTRA_r16
	if measResult_RLF_Report_EUTRA_v1690Present {
		var tmp_os_measResult_RLF_Report_EUTRA_v1690 []byte
		if tmp_os_measResult_RLF_Report_EUTRA_v1690, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode measResult_RLF_Report_EUTRA_v1690", err)
		}
		ie.measResult_RLF_Report_EUTRA_v1690 = &tmp_os_measResult_RLF_Report_EUTRA_v1690
	}
	return nil
}
