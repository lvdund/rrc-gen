package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkTxDirectCurrentCell struct {
	servCellIndex              ServCellIndex              `madatory`
	uplinkDirectCurrentBWP     []UplinkTxDirectCurrentBWP `lb:1,ub:maxNrofBWPs,madatory`
	uplinkDirectCurrentBWP_SUL []UplinkTxDirectCurrentBWP `lb:1,ub:maxNrofBWPs,optional,ext-1`
}

func (ie *UplinkTxDirectCurrentCell) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.uplinkDirectCurrentBWP_SUL) > 0
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.servCellIndex.Encode(w); err != nil {
		return utils.WrapError("Encode servCellIndex", err)
	}
	tmp_uplinkDirectCurrentBWP := utils.NewSequence[*UplinkTxDirectCurrentBWP]([]*UplinkTxDirectCurrentBWP{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
	for _, i := range ie.uplinkDirectCurrentBWP {
		tmp_uplinkDirectCurrentBWP.Value = append(tmp_uplinkDirectCurrentBWP.Value, &i)
	}
	if err = tmp_uplinkDirectCurrentBWP.Encode(w); err != nil {
		return utils.WrapError("Encode uplinkDirectCurrentBWP", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{len(ie.uplinkDirectCurrentBWP_SUL) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap UplinkTxDirectCurrentCell", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.uplinkDirectCurrentBWP_SUL) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode uplinkDirectCurrentBWP_SUL optional
			if len(ie.uplinkDirectCurrentBWP_SUL) > 0 {
				tmp_uplinkDirectCurrentBWP_SUL := utils.NewSequence[*UplinkTxDirectCurrentBWP]([]*UplinkTxDirectCurrentBWP{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
				for _, i := range ie.uplinkDirectCurrentBWP_SUL {
					tmp_uplinkDirectCurrentBWP_SUL.Value = append(tmp_uplinkDirectCurrentBWP_SUL.Value, &i)
				}
				if err = tmp_uplinkDirectCurrentBWP_SUL.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkDirectCurrentBWP_SUL", err)
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

func (ie *UplinkTxDirectCurrentCell) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.servCellIndex.Decode(r); err != nil {
		return utils.WrapError("Decode servCellIndex", err)
	}
	tmp_uplinkDirectCurrentBWP := utils.NewSequence[*UplinkTxDirectCurrentBWP]([]*UplinkTxDirectCurrentBWP{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
	fn_uplinkDirectCurrentBWP := func() *UplinkTxDirectCurrentBWP {
		return new(UplinkTxDirectCurrentBWP)
	}
	if err = tmp_uplinkDirectCurrentBWP.Decode(r, fn_uplinkDirectCurrentBWP); err != nil {
		return utils.WrapError("Decode uplinkDirectCurrentBWP", err)
	}
	ie.uplinkDirectCurrentBWP = []UplinkTxDirectCurrentBWP{}
	for _, i := range tmp_uplinkDirectCurrentBWP.Value {
		ie.uplinkDirectCurrentBWP = append(ie.uplinkDirectCurrentBWP, *i)
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

			uplinkDirectCurrentBWP_SULPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode uplinkDirectCurrentBWP_SUL optional
			if uplinkDirectCurrentBWP_SULPresent {
				tmp_uplinkDirectCurrentBWP_SUL := utils.NewSequence[*UplinkTxDirectCurrentBWP]([]*UplinkTxDirectCurrentBWP{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
				fn_uplinkDirectCurrentBWP_SUL := func() *UplinkTxDirectCurrentBWP {
					return new(UplinkTxDirectCurrentBWP)
				}
				if err = tmp_uplinkDirectCurrentBWP_SUL.Decode(extReader, fn_uplinkDirectCurrentBWP_SUL); err != nil {
					return utils.WrapError("Decode uplinkDirectCurrentBWP_SUL", err)
				}
				ie.uplinkDirectCurrentBWP_SUL = []UplinkTxDirectCurrentBWP{}
				for _, i := range tmp_uplinkDirectCurrentBWP_SUL.Value {
					ie.uplinkDirectCurrentBWP_SUL = append(ie.uplinkDirectCurrentBWP_SUL, *i)
				}
			}
		}
	}
	return nil
}
