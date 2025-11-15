package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_AssistanceInfo struct {
	affectedCarrierFreqCombInfoListMRDC []AffectedCarrierFreqCombInfoMRDC `lb:1,ub:maxNrofCombIDC,madatory`
	overheatingAssistanceSCG_r16        *[]byte                           `optional,ext-1`
	overheatingAssistanceSCG_FR2_2_r17  *[]byte                           `optional,ext-2`
}

func (ie *MRDC_AssistanceInfo) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.overheatingAssistanceSCG_r16 != nil || ie.overheatingAssistanceSCG_FR2_2_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_affectedCarrierFreqCombInfoListMRDC := utils.NewSequence[*AffectedCarrierFreqCombInfoMRDC]([]*AffectedCarrierFreqCombInfoMRDC{}, uper.Constraint{Lb: 1, Ub: maxNrofCombIDC}, false)
	for _, i := range ie.affectedCarrierFreqCombInfoListMRDC {
		tmp_affectedCarrierFreqCombInfoListMRDC.Value = append(tmp_affectedCarrierFreqCombInfoListMRDC.Value, &i)
	}
	if err = tmp_affectedCarrierFreqCombInfoListMRDC.Encode(w); err != nil {
		return utils.WrapError("Encode affectedCarrierFreqCombInfoListMRDC", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.overheatingAssistanceSCG_r16 != nil, ie.overheatingAssistanceSCG_FR2_2_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MRDC_AssistanceInfo", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.overheatingAssistanceSCG_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode overheatingAssistanceSCG_r16 optional
			if ie.overheatingAssistanceSCG_r16 != nil {
				if err = extWriter.WriteOctetString(*ie.overheatingAssistanceSCG_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode overheatingAssistanceSCG_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.overheatingAssistanceSCG_FR2_2_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode overheatingAssistanceSCG_FR2_2_r17 optional
			if ie.overheatingAssistanceSCG_FR2_2_r17 != nil {
				if err = extWriter.WriteOctetString(*ie.overheatingAssistanceSCG_FR2_2_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode overheatingAssistanceSCG_FR2_2_r17", err)
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

func (ie *MRDC_AssistanceInfo) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_affectedCarrierFreqCombInfoListMRDC := utils.NewSequence[*AffectedCarrierFreqCombInfoMRDC]([]*AffectedCarrierFreqCombInfoMRDC{}, uper.Constraint{Lb: 1, Ub: maxNrofCombIDC}, false)
	fn_affectedCarrierFreqCombInfoListMRDC := func() *AffectedCarrierFreqCombInfoMRDC {
		return new(AffectedCarrierFreqCombInfoMRDC)
	}
	if err = tmp_affectedCarrierFreqCombInfoListMRDC.Decode(r, fn_affectedCarrierFreqCombInfoListMRDC); err != nil {
		return utils.WrapError("Decode affectedCarrierFreqCombInfoListMRDC", err)
	}
	ie.affectedCarrierFreqCombInfoListMRDC = []AffectedCarrierFreqCombInfoMRDC{}
	for _, i := range tmp_affectedCarrierFreqCombInfoListMRDC.Value {
		ie.affectedCarrierFreqCombInfoListMRDC = append(ie.affectedCarrierFreqCombInfoListMRDC, *i)
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			overheatingAssistanceSCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode overheatingAssistanceSCG_r16 optional
			if overheatingAssistanceSCG_r16Present {
				var tmp_os_overheatingAssistanceSCG_r16 []byte
				if tmp_os_overheatingAssistanceSCG_r16, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode overheatingAssistanceSCG_r16", err)
				}
				ie.overheatingAssistanceSCG_r16 = &tmp_os_overheatingAssistanceSCG_r16
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			overheatingAssistanceSCG_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode overheatingAssistanceSCG_FR2_2_r17 optional
			if overheatingAssistanceSCG_FR2_2_r17Present {
				var tmp_os_overheatingAssistanceSCG_FR2_2_r17 []byte
				if tmp_os_overheatingAssistanceSCG_FR2_2_r17, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode overheatingAssistanceSCG_FR2_2_r17", err)
				}
				ie.overheatingAssistanceSCG_FR2_2_r17 = &tmp_os_overheatingAssistanceSCG_FR2_2_r17
			}
		}
	}
	return nil
}
