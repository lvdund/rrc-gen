package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PH_InfoSCG struct {
	servCellIndex               ServCellIndex                           `madatory`
	ph_Uplink                   PH_UplinkCarrierSCG                     `madatory`
	ph_SupplementaryUplink      *PH_UplinkCarrierSCG                    `optional`
	twoSRS_PUSCH_Repetition_r17 *PH_InfoSCG_twoSRS_PUSCH_Repetition_r17 `optional,ext-1`
}

func (ie *PH_InfoSCG) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.twoSRS_PUSCH_Repetition_r17 != nil
	preambleBits := []bool{hasExtensions, ie.ph_SupplementaryUplink != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.servCellIndex.Encode(w); err != nil {
		return utils.WrapError("Encode servCellIndex", err)
	}
	if err = ie.ph_Uplink.Encode(w); err != nil {
		return utils.WrapError("Encode ph_Uplink", err)
	}
	if ie.ph_SupplementaryUplink != nil {
		if err = ie.ph_SupplementaryUplink.Encode(w); err != nil {
			return utils.WrapError("Encode ph_SupplementaryUplink", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.twoSRS_PUSCH_Repetition_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PH_InfoSCG", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.twoSRS_PUSCH_Repetition_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode twoSRS_PUSCH_Repetition_r17 optional
			if ie.twoSRS_PUSCH_Repetition_r17 != nil {
				if err = ie.twoSRS_PUSCH_Repetition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode twoSRS_PUSCH_Repetition_r17", err)
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

func (ie *PH_InfoSCG) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ph_SupplementaryUplinkPresent bool
	if ph_SupplementaryUplinkPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.servCellIndex.Decode(r); err != nil {
		return utils.WrapError("Decode servCellIndex", err)
	}
	if err = ie.ph_Uplink.Decode(r); err != nil {
		return utils.WrapError("Decode ph_Uplink", err)
	}
	if ph_SupplementaryUplinkPresent {
		ie.ph_SupplementaryUplink = new(PH_UplinkCarrierSCG)
		if err = ie.ph_SupplementaryUplink.Decode(r); err != nil {
			return utils.WrapError("Decode ph_SupplementaryUplink", err)
		}
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

			twoSRS_PUSCH_Repetition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode twoSRS_PUSCH_Repetition_r17 optional
			if twoSRS_PUSCH_Repetition_r17Present {
				ie.twoSRS_PUSCH_Repetition_r17 = new(PH_InfoSCG_twoSRS_PUSCH_Repetition_r17)
				if err = ie.twoSRS_PUSCH_Repetition_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode twoSRS_PUSCH_Repetition_r17", err)
				}
			}
		}
	}
	return nil
}
