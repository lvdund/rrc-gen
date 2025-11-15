package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FailureReportSCG_EUTRA struct {
	failureType               FailureReportSCG_EUTRA_failureType `madatory`
	measResultFreqListMRDC    *MeasResultFreqListFailMRDC        `optional`
	measResultSCG_FailureMRDC *[]byte                            `optional`
	locationInfo_r16          *LocationInfo_r16                  `optional,ext-1`
}

func (ie *FailureReportSCG_EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.locationInfo_r16 != nil
	preambleBits := []bool{hasExtensions, ie.measResultFreqListMRDC != nil, ie.measResultSCG_FailureMRDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.failureType.Encode(w); err != nil {
		return utils.WrapError("Encode failureType", err)
	}
	if ie.measResultFreqListMRDC != nil {
		if err = ie.measResultFreqListMRDC.Encode(w); err != nil {
			return utils.WrapError("Encode measResultFreqListMRDC", err)
		}
	}
	if ie.measResultSCG_FailureMRDC != nil {
		if err = w.WriteOctetString(*ie.measResultSCG_FailureMRDC, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode measResultSCG_FailureMRDC", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.locationInfo_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap FailureReportSCG_EUTRA", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.locationInfo_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode locationInfo_r16 optional
			if ie.locationInfo_r16 != nil {
				if err = ie.locationInfo_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode locationInfo_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *FailureReportSCG_EUTRA) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultFreqListMRDCPresent bool
	if measResultFreqListMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultSCG_FailureMRDCPresent bool
	if measResultSCG_FailureMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.failureType.Decode(r); err != nil {
		return utils.WrapError("Decode failureType", err)
	}
	if measResultFreqListMRDCPresent {
		ie.measResultFreqListMRDC = new(MeasResultFreqListFailMRDC)
		if err = ie.measResultFreqListMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode measResultFreqListMRDC", err)
		}
	}
	if measResultSCG_FailureMRDCPresent {
		var tmp_os_measResultSCG_FailureMRDC []byte
		if tmp_os_measResultSCG_FailureMRDC, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode measResultSCG_FailureMRDC", err)
		}
		ie.measResultSCG_FailureMRDC = &tmp_os_measResultSCG_FailureMRDC
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			locationInfo_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode locationInfo_r16 optional
			if locationInfo_r16Present {
				ie.locationInfo_r16 = new(LocationInfo_r16)
				if err = ie.locationInfo_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode locationInfo_r16", err)
				}
			}
		}
	}
	return nil
}
