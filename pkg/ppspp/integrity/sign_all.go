package integrity

import (
	"bufio"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/MemeLabs/go-ppspp/pkg/binmap"
	"github.com/MemeLabs/go-ppspp/pkg/ppspp/codec"
)

var errMissingSignature = errors.New("missing signature")

// SignAllOptions ...
type SignAllOptions struct {
	LiveDiscardWindow int
	ChunkSize         int
	Verifier          SignatureVerifier
	Hash              hashFunc
}

// NewSignAllSwarmVerifier ...
func NewSignAllSwarmVerifier(o *SignAllOptions) *SignAllSwarmVerifier {
	return &SignAllSwarmVerifier{
		mask:              uint64(o.LiveDiscardWindow - 1),
		size:              binmap.Bin(o.LiveDiscardWindow * 2),
		timestamps:        make([]time.Time, o.LiveDiscardWindow),
		signatures:        make([]byte, o.LiveDiscardWindow*o.Verifier.Size()),
		chunkSize:         o.ChunkSize,
		signatureVerifier: o.Verifier,
	}
}

// SignAllSwarmVerifier ...
type SignAllSwarmVerifier struct {
	lock              sync.Mutex
	mask              uint64
	size              binmap.Bin
	head, tail        binmap.Bin
	timestamps        []time.Time
	signatures        []byte
	chunkSize         int
	signatureVerifier SignatureVerifier
}

// WriteIntegrity ...
func (v *SignAllSwarmVerifier) WriteIntegrity(b binmap.Bin, m *binmap.Map, w Writer) (int, error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	var n int

	for l, r := b.BaseLeft(), b.BaseRight(); l <= r; l += 2 {
		if l < v.tail || l >= v.head {
			return n, errMissingSignature
		}

		i := uint64(l>>1) & v.mask
		vs := uint64(v.signatureVerifier.Size())

		nn, err := w.WriteSignedIntegrity(codec.SignedIntegrity{
			Address:   codec.Address(l),
			Timestamp: codec.Timestamp{Time: v.timestamps[i]},
			Signature: v.signatures[i*vs : (i+1)*vs],
		})
		n += nn
		if err != nil {
			return n, err
		}
	}

	return n, nil
}

func (v *SignAllSwarmVerifier) storeSignature(b binmap.Bin, ts time.Time, sig []byte) {
	v.lock.Lock()
	defer v.lock.Unlock()

	if b > v.head {
		v.head = b + 2
		if v.head > v.size {
			v.tail = v.head - v.size
		}
	}

	i := uint64(b>>1) & v.mask
	vs := uint64(v.signatureVerifier.Size())

	v.timestamps[i] = ts
	copy(v.signatures[i*vs:(i+1)*vs], sig)
}

// ChannelVerifier ...
func (v *SignAllSwarmVerifier) ChannelVerifier() ChannelVerifier {
	return newSignAllChannelVerifier(v)
}

func newSignAllChannelVerifier(v *SignAllSwarmVerifier) *SignAllChannelVerifier {
	return &SignAllChannelVerifier{
		chunkVerifier: SignAllChunkVerifier{
			bin:           binmap.None,
			swarmVerifier: v,
		},
	}
}

// SignAllChannelVerifier ...
type SignAllChannelVerifier struct {
	chunkVerifier SignAllChunkVerifier
}

// ChunkVerifier ...
func (v *SignAllChannelVerifier) ChunkVerifier(b binmap.Bin) ChunkVerifier {
	return &v.chunkVerifier
}

// SignAllChunkVerifier ...
type SignAllChunkVerifier struct {
	swarmVerifier *SignAllSwarmVerifier
	bin           binmap.Bin
	timestamps    []time.Time
	signatures    []byte
}

// SetSignedIntegrity ...
func (v *SignAllChunkVerifier) SetSignedIntegrity(b binmap.Bin, ts time.Time, sig []byte) {
	if v.bin == binmap.None {
		v.bin = b
	}
	v.timestamps = append(v.timestamps, ts)
	v.signatures = append(v.signatures, sig...)
}

// SetIntegrity ...
func (v *SignAllChunkVerifier) SetIntegrity(b binmap.Bin, d []byte) {}

func (v *SignAllChunkVerifier) verify(b binmap.Bin, d []byte) bool {
	if b.BaseLeft() != v.bin {
		return false
	}

	l := int(b.BaseLength())
	if l > len(v.timestamps) {
		return false
	}
	for i := 0; i < l; i++ {
		ts := v.timestamps[i]
		chunk := d[i*v.swarmVerifier.chunkSize : (i+1)*v.swarmVerifier.chunkSize]
		sig := v.signatures[i*v.swarmVerifier.signatureVerifier.Size() : (i+1)*v.swarmVerifier.signatureVerifier.Size()]

		if !v.swarmVerifier.signatureVerifier.Verify(ts, chunk, sig) {
			return false
		}
		v.swarmVerifier.storeSignature(v.bin+binmap.Bin(i*2), ts, sig)
	}

	return true
}

// Verify ...
func (v *SignAllChunkVerifier) Verify(b binmap.Bin, d []byte) bool {
	verified := v.verify(b, d)

	v.bin = binmap.None
	v.timestamps = v.timestamps[:0]
	v.signatures = v.signatures[:0]

	return verified
}

// SignAllWriterOptions ...
type SignAllWriterOptions struct {
	Verifier  *SignAllSwarmVerifier
	Writer    WriteFlusher
	ChunkSize int
	Signer    SignatureSigner
}

// NewSignAllWriter ...
func NewSignAllWriter(o *SignAllWriterOptions) *SignAllWriter {
	sw := &signAllWriter{
		chunkSize:       o.ChunkSize,
		swarmVerifier:   o.Verifier,
		signatureSigner: o.Signer,
		w:               o.Writer,
	}
	return &SignAllWriter{
		bw: bufio.NewWriterSize(sw, o.ChunkSize),
	}
}

// SignAllWriter ...
type SignAllWriter struct {
	bw *bufio.Writer
}

// Write ...
func (w *SignAllWriter) Write(p []byte) (int, error) {
	return w.bw.Write(p)
}

// Flush ...
func (w *SignAllWriter) Flush() error {
	return w.bw.Flush()
}

type signAllWriter struct {
	b               binmap.Bin
	chunkSize       int
	swarmVerifier   *SignAllSwarmVerifier
	signatureSigner SignatureSigner
	w               io.Writer
}

func (w *signAllWriter) Write(p []byte) (int, error) {
	if len(p) > w.chunkSize {
		p = p[:w.chunkSize]
	}

	ts := time.Now()
	sig := w.signatureSigner.Sign(ts, p)
	w.swarmVerifier.storeSignature(w.b, ts, sig)

	w.b += 2

	return w.w.Write(p)
}