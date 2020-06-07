package svc

import "io"

// PipeRW represents a reader-writer pipe pair.
type PipeRW interface {
	Read(data []byte) (n int, err error)
	Write(data []byte) (n int, err error)
	Close() error
	CloseReader() error
	CloseWriter() error
	CloseWithError(err error) error
	CloseReaderWithError(err error) error
	CloseWriterWithError(err error) error
}

type pipeRW struct {
	r *io.PipeReader
	w *io.PipeWriter
}

func (p *pipeRW) Read(data []byte) (n int, err error) {
	return p.r.Read(data)
}

func (p *pipeRW) Write(data []byte) (n int, err error) {
	return p.w.Write(data)
}

func (p *pipeRW) Close() error {
	if err := p.r.Close(); err != nil {
		return err
	}
	if err := p.w.Close(); err != nil {
		return err
	}
	return nil
}

func (p *pipeRW) CloseReader() error {
	return p.r.Close()
}

func (p *pipeRW) CloseWriter() error {
	return p.w.Close()
}

func (p *pipeRW) CloseWithError(err error) error {
	if err := p.r.CloseWithError(err); err != nil {
		return err
	}
	if err := p.w.CloseWithError(err); err != nil {
		return err
	}
	return nil
}

func (p *pipeRW) CloseReaderWithError(err error) error {
	return p.r.CloseWithError(err)
}

func (p *pipeRW) CloseWriterWithError(err error) error {
	return p.w.CloseWithError(err)
}

func NewPipeRW() PipeRW {
	r, w := io.Pipe()
	return &pipeRW{r: r, w: w}
}
