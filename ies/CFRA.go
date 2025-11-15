package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA struct {
	occasions                 *CFRA_occasions `optional`
	resources                 CFRA_resources  `lb:1,ub:maxRA_SSB_Resources,madatory`
	totalNumberOfRA_Preambles *int64          `lb:1,ub:63,optional,ext-1`
}

func (ie *CFRA) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.totalNumberOfRA_Preambles != nil
	preambleBits := []bool{hasExtensions, ie.occasions != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.occasions != nil {
		if err = ie.occasions.Encode(w); err != nil {
			return utils.WrapError("Encode occasions", err)
		}
	}
	if err = ie.resources.Encode(w); err != nil {
		return utils.WrapError("Encode resources", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.totalNumberOfRA_Preambles != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CFRA", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.totalNumberOfRA_Preambles != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode totalNumberOfRA_Preambles optional
			if ie.totalNumberOfRA_Preambles != nil {
				if err = extWriter.WriteInteger(*ie.totalNumberOfRA_Preambles, &uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
					return utils.WrapError("Encode totalNumberOfRA_Preambles", err)
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

func (ie *CFRA) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var occasionsPresent bool
	if occasionsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if occasionsPresent {
		ie.occasions = new(CFRA_occasions)
		if err = ie.occasions.Decode(r); err != nil {
			return utils.WrapError("Decode occasions", err)
		}
	}
	if err = ie.resources.Decode(r); err != nil {
		return utils.WrapError("Decode resources", err)
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

			totalNumberOfRA_PreamblesPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode totalNumberOfRA_Preambles optional
			if totalNumberOfRA_PreamblesPresent {
				var tmp_int_totalNumberOfRA_Preambles int64
				if tmp_int_totalNumberOfRA_Preambles, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
					return utils.WrapError("Decode totalNumberOfRA_Preambles", err)
				}
				ie.totalNumberOfRA_Preambles = &tmp_int_totalNumberOfRA_Preambles
			}
		}
	}
	return nil
}
