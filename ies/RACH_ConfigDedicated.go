package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigDedicated struct {
	cfra                         *CFRA              `optional`
	ra_Prioritization            *RA_Prioritization `optional`
	ra_PrioritizationTwoStep_r16 *RA_Prioritization `optional,ext-1`
	cfra_TwoStep_r16             *CFRA_TwoStep_r16  `optional,ext-1`
}

func (ie *RACH_ConfigDedicated) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ra_PrioritizationTwoStep_r16 != nil || ie.cfra_TwoStep_r16 != nil
	preambleBits := []bool{hasExtensions, ie.cfra != nil, ie.ra_Prioritization != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cfra != nil {
		if err = ie.cfra.Encode(w); err != nil {
			return utils.WrapError("Encode cfra", err)
		}
	}
	if ie.ra_Prioritization != nil {
		if err = ie.ra_Prioritization.Encode(w); err != nil {
			return utils.WrapError("Encode ra_Prioritization", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.ra_PrioritizationTwoStep_r16 != nil || ie.cfra_TwoStep_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RACH_ConfigDedicated", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ra_PrioritizationTwoStep_r16 != nil, ie.cfra_TwoStep_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ra_PrioritizationTwoStep_r16 optional
			if ie.ra_PrioritizationTwoStep_r16 != nil {
				if err = ie.ra_PrioritizationTwoStep_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ra_PrioritizationTwoStep_r16", err)
				}
			}
			// encode cfra_TwoStep_r16 optional
			if ie.cfra_TwoStep_r16 != nil {
				if err = ie.cfra_TwoStep_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cfra_TwoStep_r16", err)
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

func (ie *RACH_ConfigDedicated) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var cfraPresent bool
	if cfraPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_PrioritizationPresent bool
	if ra_PrioritizationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if cfraPresent {
		ie.cfra = new(CFRA)
		if err = ie.cfra.Decode(r); err != nil {
			return utils.WrapError("Decode cfra", err)
		}
	}
	if ra_PrioritizationPresent {
		ie.ra_Prioritization = new(RA_Prioritization)
		if err = ie.ra_Prioritization.Decode(r); err != nil {
			return utils.WrapError("Decode ra_Prioritization", err)
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

			ra_PrioritizationTwoStep_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cfra_TwoStep_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ra_PrioritizationTwoStep_r16 optional
			if ra_PrioritizationTwoStep_r16Present {
				ie.ra_PrioritizationTwoStep_r16 = new(RA_Prioritization)
				if err = ie.ra_PrioritizationTwoStep_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ra_PrioritizationTwoStep_r16", err)
				}
			}
			// decode cfra_TwoStep_r16 optional
			if cfra_TwoStep_r16Present {
				ie.cfra_TwoStep_r16 = new(CFRA_TwoStep_r16)
				if err = ie.cfra_TwoStep_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cfra_TwoStep_r16", err)
				}
			}
		}
	}
	return nil
}
