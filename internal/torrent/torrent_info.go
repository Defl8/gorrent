package torrent

import (
	"github.com/Defl8/gorrent/internal/bencode"
	"time"
)

type MetaInfo struct {
	Announce     string
	CreationDate *time.Time
	Comment      *string
	CreatedBy    *string
	Info         Info
	InfoHash     [20]byte
	RawInfo      []byte
}

type Info struct {
	Name        string
	PieceLength int
	Pieces      []byte
	Private     *int
	Length      int
	Files       []FileInfo
}

type FileInfo struct {
	Length int
	Path   []string
}

type TorrentInfo struct {
	Meta MetaInfo
	Info Info
}
