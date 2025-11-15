package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceExt_v1610 struct {
	interlaceAllocation_r16       *PUCCH_ResourceExt_v1610_interlaceAllocation_r16       `lb:0,ub:4,optional`
	format_v1610                  *PUCCH_ResourceExt_v1610_format_v1610                  `lb:0,ub:9,optional`
	formatExt_v1700               *PUCCH_ResourceExt_v1610_formatExt_v1700               `lb:1,ub:16,optional,ext-1`
	pucch_RepetitionNrofSlots_r17 *PUCCH_ResourceExt_v1610_pucch_RepetitionNrofSlots_r17 `optional,ext-1`
}

func (ie *PUCCH_ResourceExt_v1610) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.formatExt_v1700 != nil || ie.pucch_RepetitionNrofSlots_r17 != nil
	preambleBits := []bool{hasExtensions, ie.interlaceAllocation_r16 != nil, ie.format_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.interlaceAllocation_r16 != nil {
		if err = ie.interlaceAllocation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interlaceAllocation_r16", err)
		}
	}
	if ie.format_v1610 != nil {
		if err = ie.format_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode format_v1610", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.formatExt_v1700 != nil || ie.pucch_RepetitionNrofSlots_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUCCH_ResourceExt_v1610", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.formatExt_v1700 != nil, ie.pucch_RepetitionNrofSlots_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode formatExt_v1700 optional
			if ie.formatExt_v1700 != nil {
				if err = ie.formatExt_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode formatExt_v1700", err)
				}
			}
			// encode pucch_RepetitionNrofSlots_r17 optional
			if ie.pucch_RepetitionNrofSlots_r17 != nil {
				if err = ie.pucch_RepetitionNrofSlots_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_RepetitionNrofSlots_r17", err)
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

func (ie *PUCCH_ResourceExt_v1610) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var interlaceAllocation_r16Present bool
	if interlaceAllocation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var format_v1610Present bool
	if format_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if interlaceAllocation_r16Present {
		ie.interlaceAllocation_r16 = new(PUCCH_ResourceExt_v1610_interlaceAllocation_r16)
		if err = ie.interlaceAllocation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interlaceAllocation_r16", err)
		}
	}
	if format_v1610Present {
		ie.format_v1610 = new(PUCCH_ResourceExt_v1610_format_v1610)
		if err = ie.format_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode format_v1610", err)
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

			formatExt_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_RepetitionNrofSlots_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode formatExt_v1700 optional
			if formatExt_v1700Present {
				ie.formatExt_v1700 = new(PUCCH_ResourceExt_v1610_formatExt_v1700)
				if err = ie.formatExt_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode formatExt_v1700", err)
				}
			}
			// decode pucch_RepetitionNrofSlots_r17 optional
			if pucch_RepetitionNrofSlots_r17Present {
				ie.pucch_RepetitionNrofSlots_r17 = new(PUCCH_ResourceExt_v1610_pucch_RepetitionNrofSlots_r17)
				if err = ie.pucch_RepetitionNrofSlots_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_RepetitionNrofSlots_r17", err)
				}
			}
		}
	}
	return nil
}
