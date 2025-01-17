// Code generated by fastssz. DO NOT EDIT.
// Hash: 733221663201260573e0519b81dc7f95c41564034ef35065ab0463a3f8fa4063
// Version: 0.1.3
package deneb

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SignedBlobSidecar object
func (s *SignedBlobSidecar) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedBlobSidecar object to a target array
func (s *SignedBlobSidecar) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(BlobSidecar)
	}
	if dst, err = s.Message.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Signature'
	dst = append(dst, s.Signature[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the SignedBlobSidecar object
func (s *SignedBlobSidecar) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 131352 {
		return ssz.ErrSize
	}

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(BlobSidecar)
	}
	if err = s.Message.UnmarshalSSZ(buf[0:131256]); err != nil {
		return err
	}

	// Field (1) 'Signature'
	copy(s.Signature[:], buf[131256:131352])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedBlobSidecar object
func (s *SignedBlobSidecar) SizeSSZ() (size int) {
	size = 131352
	return
}

// HashTreeRoot ssz hashes the SignedBlobSidecar object
func (s *SignedBlobSidecar) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedBlobSidecar object with a hasher
func (s *SignedBlobSidecar) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Message'
	if s.Message == nil {
		s.Message = new(BlobSidecar)
	}
	if err = s.Message.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Signature'
	hh.PutBytes(s.Signature[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedBlobSidecar object
func (s *SignedBlobSidecar) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
